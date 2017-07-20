package indexer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"

	"strings"

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
func NewFrame(path, index, name string) (f *Frame, err error) {
	var td *TermDict
	if td, err = NewTermDict(path); err != nil {
		return
	}
	f = &Frame{
		path:      path,
		index:     index,
		name:      name,
		td:        td,
		fragments: make(map[uint64]*pilosa.Fragment),
	}
	err = f.Open()
	return
}

//Open opens an existing frame
func (f *Frame) Open() (err error) {
	var sliceList []uint64
	if sliceList, err = getSliceList(f.path); err != nil {
		return
	}
	fmt.Printf("sliceList: %v\n", sliceList)
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
	var d *os.File
	var fns []string
	var num uint64
	fragDir := filepath.Join(dir, "fragments")
	if err = os.MkdirAll(fragDir, 0700); err != nil {
		err = errors.Wrap(err, "")
		return
	}
	d, err = os.Open(fragDir)
	if err != nil {
		err = errors.Wrap(err, "")
		return
	}
	fns, err = d.Readdirnames(0)
	if err != nil {
		err = errors.Wrap(err, "")
		return
	}
	re := regexp.MustCompile("^(?P<num>[0-9]+)$")
	for _, fn := range fns {
		subs := re.FindStringSubmatch(fn)
		if subs == nil {
			continue
		}
		num, err = strconv.ParseUint(subs[1], 10, 64)
		if err != nil {
			err = errors.Wrap(err, "")
			return
		}
		numList = append(numList, num)
	}
	return
}

// Close closes all fragments without removing files on disk.
func (f *Frame) Close() (err error) {
	for _, fragment := range f.fragments {
		if err = fragment.Close(); err != nil {
			err = errors.Wrap(err, "")
			return
		}
	}
	err = f.td.Close()
	return
}

// Destroy closes all fragments, removes all files on disk.
func (f *Frame) Destroy() (err error) {
	for _, fragment := range f.fragments {
		if err = fragment.Close(); err != nil {
			err = errors.Wrap(err, "")
			return
		}
	}
	if err = os.RemoveAll(filepath.Join(f.path, "fragments")); err != nil {
		err = errors.Wrap(err, "")
		return
	}
	err = f.td.Destroy()
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

// SetBit sets a bit within the frame, and expands fragments if necessary.
func (f *Frame) SetBit(rowID, colID uint64) (changed bool, err error) {
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

// ClearBit clears a bit within the frame.
func (f *Frame) ClearBit(rowID, colID uint64) (changed bool, err error) {
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
		if _, err = f.SetBit(termID, docID); err != nil {
			return
		}
	}
	return
}

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
