package indexer

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"strings"

	"github.com/pilosa/pilosa"
)

// Frame represents a string field of an index. Refers to pilosa.Frame and pilosa.View.
type Frame struct {
	mu    sync.Mutex
	path  string
	index string
	name  string

	fragments map[uint64]*pilosa.Fragment //map slice to Fragment
	td        *TermDict
}

// NewFrame returns a new instance of frame, and initializes it.
func NewFrame(path, index, name string) *Frame {
	f := &Frame{
		path:      path,
		index:     index,
		name:      name,
		td:        &TermDict{Dir: path},
		fragments: make(map[uint64]*pilosa.Fragment),
	}
	for slice := uint64(0); ; slice++ {
		fp := f.FragmentPath(slice)
		if _, err := os.Stat(fp); os.IsNotExist(err) {
			//path does not exist
			break
		}
		fragment := pilosa.NewFragment(fp, f.index, f.name, pilosa.ViewStandard, slice)
		f.fragments[slice] = fragment
	}
	return f
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
func (f *Frame) MaxSlice() uint64 {
	f.mu.Lock()
	defer f.mu.Unlock()

	return uint64(len(f.fragments))
}

// SetBit sets a bit on a view within the frame.
func (f *Frame) SetBit(rowID, colID uint64) (changed bool, err error) {
	slice := colID / pilosa.SliceWidth
	fragment, ok := f.fragments[slice]
	if !ok {
		err = errors.New("column out of bounds")
		return
	}
	changed, err = fragment.SetBit(rowID, colID)
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
	changed, err = fragment.SetBit(rowID, colID)
	return
}

// ParseAndIndex parses and index a field
func (f *Frame) ParseAndIndex(docID uint64, text string) (err error) {
	terms := strings.SplitN(text, " ", -1)
	ids, err := f.td.GetTermsID(terms)
	for _, termID := range ids {
		if _, err = f.SetBit(termID, docID); err != nil {
			return
		}
	}
	return
}
