package indexer

import (
	"io"
	"os"
	"sync"

	"github.com/pilosa/pilosa/roaring"
)

const (
	// SliceWidth is the number of column IDs in a slice.
	SliceWidth = 1048576

	// SnapshotExt is the file extension used for an in-process snapshot.
	SnapshotExt = ".snapshotting"

	// CopyExt is the file extension used for the temp file used while copying.
	CopyExt = ".copying"

	// CacheExt is the file extension for persisted cache ids.
	CacheExt = ".cache"

	// HashBlockSize is the number of rows in a merkle hash block.
	HashBlockSize = 100
)

// Fragment represents the intersection of a frame and slice in an index.
type Fragment struct {
	mu sync.Mutex

	// Composite identifiers
	index string
	frame string
	view  string
	slice uint64

	// File-backed storage
	path        string
	file        *os.File
	storage     *roaring.Bitmap
	storageData []byte
	opN         int // number of ops since snapshot

	// Stats reporting.
	maxRowID uint64

	// Number of operations performed before performing a snapshot.
	// This limits the size of fragments on the heap and flushes them to disk
	// so that they can be mmapped and heap utilization can be kept low.
	MaxOpN int

	// Writer used for out-of-band log entries.
	LogOutput io.Writer
}
