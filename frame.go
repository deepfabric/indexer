package indexer

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/deepfabric/bkdtree"
	"github.com/deepfabric/pilosa"
	"github.com/pkg/errors"
)

//View maps slice to Fragment.
type View struct {
	frags    map[uint64]*pilosa.Fragment //maps slice to Fragment
	maxSlice uint64
}

// Frame represents a string field of an index. Refers to pilosa.Frame and pilosa.View.
type Frame struct {
	mu    sync.Mutex
	path  string
	index string
	name  string

	views map[string]*View //map view mode to view.
	td    *TermDict
}

// NewFrame returns a new instance of frame, and initializes it.
func NewFrame(path, index, name string, overwrite bool) (f *Frame, err error) {
	var td *TermDict
	if td, err = NewTermDict(path, overwrite); err != nil {
		return
	}
	if overwrite {
		viewModes := []string{pilosa.ViewStandard, pilosa.ViewInverse}
		for _, mode := range viewModes {
			if err = os.RemoveAll(filepath.Join(path, mode)); err != nil {
				err = errors.Wrap(err, "")
				return
			}
		}
	}
	f = &Frame{
		path:  path,
		index: index,
		name:  name,
		td:    td,
	}
	err = f.openFragments()
	return
}

//Open opens an existing frame
func (f *Frame) Open() (err error) {
	if err = f.openFragments(); err != nil {
		return
	}
	err = f.td.Open()
	return
}

func (f *Frame) openFragments() (err error) {
	var sliceList []uint64
	var view *View
	if f.views != nil {
		panic("frame is already open")
	}
	f.views = make(map[string]*View)
	viewModes := []string{pilosa.ViewStandard, pilosa.ViewInverse}
	for _, mode := range viewModes {
		if sliceList, err = getSliceList(f.path, mode); err != nil {
			return
		}
		view = &View{
			frags:    make(map[uint64]*pilosa.Fragment),
			maxSlice: 0,
		}
		f.views[mode] = view
		for _, slice := range sliceList {
			fp := f.FragmentPath(mode, slice)
			fragment := pilosa.NewFragment(fp, f.index, f.name, mode, slice)
			if err = fragment.Open(); err != nil {
				err = errors.Wrap(err, "")
				return
			}
			view.frags[slice] = fragment
			if view.maxSlice < slice {
				view.maxSlice = slice
			}
		}
	}
	return
}

func getSliceList(dir, viewMode string) (numList []uint64, err error) {
	var num uint64
	var matches [][]string
	fragDir := filepath.Join(dir, viewMode)
	if err = os.MkdirAll(fragDir, 0700); err != nil {
		err = errors.Wrap(err, "")
		return
	}
	if matches, err = bkdtree.FilepathGlob(fragDir, "^(?P<num>[0-9]+)$"); err != nil {
		return
	}
	for _, match := range matches {
		num, err = strconv.ParseUint(match[1], 10, 64)
		if err != nil {
			err = errors.Wrap(err, "")
			return
		}
		numList = append(numList, num)
	}
	return
}

// Close closes all fragments without removing files on disk.
// It's allowed to invoke Close multiple times.
func (f *Frame) Close() (err error) {
	if err = f.closeFragments(); err != nil {
		return
	}
	err = f.td.Close()
	return
}

// Destroy closes all fragments, removes all files on disk.
// It's allowed to invoke Close before or after Destroy.
func (f *Frame) Destroy() (err error) {
	if err = f.closeFragments(); err != nil {
		return
	}
	viewModes := []string{pilosa.ViewStandard, pilosa.ViewInverse}
	for _, mode := range viewModes {
		if err = os.RemoveAll(filepath.Join(f.path, mode)); err != nil {
			err = errors.Wrap(err, "")
			return
		}
	}
	err = f.td.Destroy()
	return
}

func (f *Frame) closeFragments() (err error) {
	for _, view := range f.views {
		for _, fragment := range view.frags {
			if err = fragment.Close(); err != nil {
				err = errors.Wrap(err, "")
				return
			}
		}
	}
	f.views = nil
	return
}

// FragmentPath returns the path to a fragment
func (f *Frame) FragmentPath(viewMode string, slice uint64) string {
	return filepath.Join(f.path, viewMode, strconv.FormatUint(slice, 10))
}

// Name returns the name the frame was initialized with.
func (f *Frame) Name() string { return f.name }

// Index returns the index name the frame was initialized with.
func (f *Frame) Index() string { return f.index }

// Path returns the path the frame was initialized with.
func (f *Frame) Path() string { return f.path }

// MaxStandardSlice returns the max slice in the frame.
func (f *Frame) MaxStandardSlice() uint64 { return f.views[pilosa.ViewStandard].maxSlice }

// MaxInverseSlice returns the max slice in the frame.
func (f *Frame) MaxInverseSlice() uint64 { return f.views[pilosa.ViewInverse].maxSlice }

// setBit sets a bit within the frame, and expands fragments if necessary.
func (f *Frame) setBit(docID, termID uint64) (changed bool, err error) {
	var slice uint64
	var fragment *pilosa.Fragment
	var ok bool
	var rowID, colID uint64
	//ViewStandard uses docID as row id, termID as column id.
	//ViewInverse uses termID as row id, docID as column id.
	for mode, view := range f.views {
		switch mode {
		case pilosa.ViewStandard:
			rowID, colID = docID, termID
		case pilosa.ViewInverse:
			rowID, colID = termID, docID
		default:
			panic("unknown view mode")
		}
		slice = colID / pilosa.SliceWidth
		fragment, ok = view.frags[slice]
		if !ok {
			fp := f.FragmentPath(mode, slice)
			fragment = pilosa.NewFragment(fp, f.index, f.name, mode, slice)
			if err = fragment.Open(); err != nil {
				err = errors.Wrap(err, "")
				return
			}
			view.frags[slice] = fragment
		}
		changed, err = fragment.SetBit(rowID, colID)
		if view.maxSlice < slice {
			view.maxSlice = slice
		}
	}
	return
}

// clearBit clears a bit within the frame.
func (f *Frame) clearBit(docID, termID uint64) (changed bool, err error) {
	var rowID, colID uint64
	//ViewStandard uses docID as row id, termID as column id.
	//ViewInverse uses termID as row id, docID as column id.
	for mode, view := range f.views {
		switch mode {
		case pilosa.ViewStandard:
			rowID, colID = docID, termID
		case pilosa.ViewInverse:
			rowID, colID = termID, docID
		default:
			panic("unknown view mode")
		}
		slice := colID / pilosa.SliceWidth
		fragment, ok := view.frags[slice]
		if !ok {
			err = errors.New("column out of bounds")
			return
		}
		if changed, err = fragment.ClearBit(rowID, colID); err != nil {
			return
		}
	}
	return
}

// Bits returns bits set in frame.
func (f *Frame) Bits(viewMode string) (bits map[uint64][]uint64, err error) {
	bits = make(map[uint64][]uint64)
	var columns []uint64
	var view *View
	var ok bool
	if view, ok = f.views[viewMode]; !ok {
		err = errors.Errorf("unknown view mode %s", viewMode)
		return
	}
	for _, fragment := range view.frags {
		err = fragment.ForEachBit(
			func(rowID, columnID uint64) error {
				columns, ok = bits[rowID]
				if ok {
					columns = append(columns, columnID)
				} else {
					columns = []uint64{columnID}
				}
				bits[rowID] = columns
				return nil
			},
		)
		if err != nil {
			return
		}
	}
	return
}

// ParseAndIndex parses and index a field.
func (f *Frame) ParseAndIndex(docID uint64, text string) (err error) {
	//https://stackoverflow.com/questions/13737745/split-a-string-on-whitespace-in-go
	terms := strings.Fields(text)
	ids, err := f.td.CreateTermsIfNotExist(terms)
	if err != nil {
		return
	}
	for _, termID := range ids {
		if _, err = f.setBit(docID, termID); err != nil {
			return
		}
	}
	return
}

// RemoveDoc clear bits of given document.
func (f *Frame) RemoveDoc(docID uint64) (err error) {
	//TODO: using mark-deletion to speed up?
	var bm *pilosa.Bitmap
	if bm, err = f.getColumns(pilosa.ViewStandard, docID); err != nil {
		return
	}
	termIDs := bm.Bits()
	for _, termID := range termIDs {
		if _, err = f.clearBit(docID, termID); err != nil {
			return
		}
	}
	return
}

//Query query which documents contain the given term.
func (f *Frame) Query(term string) (bm *pilosa.Bitmap, err error) {
	termID, found := f.td.GetTermID(term)
	if !found {
		bm = pilosa.NewBitmap()
		return
	}
	bm, err = f.getColumns(pilosa.ViewInverse, termID)
	return
}

//Query query which documents contain the given term.
func (f *Frame) getColumns(viewMode string, rowID uint64) (bm *pilosa.Bitmap, err error) {
	var view *View
	var ok bool
	if view, ok = f.views[viewMode]; !ok {
		err = errors.Errorf("unknown view mode %s", viewMode)
		return
	}
	bm = pilosa.NewBitmap()
	for _, fragment := range view.frags {
		bm2 := fragment.Row(rowID)
		bm.Merge(bm2)
	}
	return
}
