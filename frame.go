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

// Frame represents a string field of an index. Refers to pilosa.Frame and pilosa.View.
type Frame struct {
	mu       sync.Mutex
	path     string
	index    string
	name     string
	maxSlice uint64

	fragments map[uint64]*pilosa.Fragment //map slice to Fragment
	td        *TermDict
}

// NewFrame returns a new instance of frame, and initializes it.
func NewFrame(path, index, name string, overwrite bool) (f *Frame, err error) {
	var td *TermDict
	if td, err = NewTermDict(path, overwrite); err != nil {
		return
	}
	if overwrite {
		if err = os.RemoveAll(filepath.Join(path, "fragments")); err != nil {
			err = errors.Wrap(err, "")
			return
		}
	}
	f = &Frame{
		path:      path,
		index:     index,
		name:      name,
		td:        td,
		fragments: make(map[uint64]*pilosa.Fragment),
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
	if sliceList, err = getSliceList(f.path); err != nil {
		return
	}
	for _, slice := range sliceList {
		fp := f.FragmentPath(slice)
		fragment := pilosa.NewFragment(fp, f.index, f.name, pilosa.ViewStandard, slice)
		if err = fragment.Open(); err != nil {
			err = errors.Wrap(err, "")
			return
		}
		f.fragments[slice] = fragment
		if f.maxSlice < slice {
			f.maxSlice = slice
		}
	}
	return
}

func getSliceList(dir string) (numList []uint64, err error) {
	var num uint64
	var matches [][]string
	fragDir := filepath.Join(dir, "fragments")
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
	if err = os.RemoveAll(filepath.Join(f.path, "fragments")); err != nil {
		err = errors.Wrap(err, "")
		return
	}
	err = f.td.Destroy()
	return
}

func (f *Frame) closeFragments() (err error) {
	for slice, fragment := range f.fragments {
		if err = fragment.Close(); err != nil {
			err = errors.Wrap(err, "")
			return
		}
		delete(f.fragments, slice)
	}
	return
}

// FragmentPath returns the path to a fragment
func (f *Frame) FragmentPath(slice uint64) string {
	return filepath.Join(f.path, "fragments", strconv.FormatUint(slice, 10))
}

// Fragment returns a fragment in the view by slice.
func (f *Frame) Fragment(slice uint64) *pilosa.Fragment {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.fragments[slice]
}

// Name returns the name the frame was initialized with.
func (f *Frame) Name() string { return f.name }

// Index returns the index name the frame was initialized with.
func (f *Frame) Index() string { return f.index }

// Path returns the path the frame was initialized with.
func (f *Frame) Path() string { return f.path }

// MaxSlice returns the max slice in the frame.
func (f *Frame) MaxSlice() uint64 { return f.maxSlice }

// setBit sets a bit within the frame, and expands fragments if necessary.
func (f *Frame) setBit(rowID, colID uint64) (changed bool, err error) {
	slice := colID / pilosa.SliceWidth
	fragment, ok := f.fragments[slice]
	if !ok {
		fp := f.FragmentPath(slice)
		fragment = pilosa.NewFragment(fp, f.index, f.name, pilosa.ViewStandard, slice)
		if err = fragment.Open(); err != nil {
			err = errors.Wrap(err, "")
			return
		}
		f.fragments[slice] = fragment
	}
	changed, err = fragment.SetBit(rowID, colID)
	if f.maxSlice < slice {
		f.maxSlice = slice
	}
	return
}

// clearBit clears a bit within the frame.
func (f *Frame) clearBit(rowID, colID uint64) (changed bool, err error) {
	slice := colID / pilosa.SliceWidth
	fragment, ok := f.fragments[slice]
	if !ok {
		err = errors.New("column out of bounds")
		return
	}
	changed, err = fragment.ClearBit(rowID, colID)
	return
}

// ParseAndIndex parses and index a field.
func (f *Frame) ParseAndIndex(docID uint64, text string) (err error) {
	terms := strings.SplitN(text, " ", -1)
	ids, err := f.td.CreateTermsIfNotExist(terms)
	if err != nil {
		return
	}
	for _, termID := range ids {
		if _, err = f.setBit(termID, docID); err != nil {
			return
		}
	}
	return
}

// RemoveDoc clear bits of given document.
func (f *Frame) RemoveDoc(docID uint64) (err error) {
	//TODO: using mark-deletion to speed up
	numTerms := f.td.Count()
	for termID := uint64(0); termID < numTerms; termID++ {
		if _, err = f.clearBit(termID, docID); err != nil {
			return
		}
	}
	return
}

//Query query which documents contain the given term.
func (f *Frame) Query(term string) (bm *pilosa.Bitmap, err error) {
	bm = pilosa.NewBitmap()
	termID, found := f.td.GetTermID(term)
	if !found {
		return
	}
	for _, fragment := range f.fragments {
		bm2 := fragment.Row(termID)
		bm.Merge(bm2)
	}
	return
}
