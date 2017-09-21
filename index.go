package indexer

import (
	"container/heap"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/deepfabric/bkdtree"
	"github.com/deepfabric/indexer/cql"
	"github.com/deepfabric/pilosa"
	"github.com/pkg/errors"
)

const (
	LiveDocs string = "__liveDocs" // the directory where stores Index.liveDocs
)

//Index is created by CqlCreate
type Index struct {
	MainDir string
	DocProt *cql.DocumentWithIdx //document prototype. persisted to an index-specific file

	rwlock   sync.RWMutex //concurrent access of bkds, frames, liveDocs
	bkds     map[string]*bkdtree.BkdTree
	frames   map[string]*Frame
	liveDocs *Frame //row 0 of this frame stores a bitmap of live docIDs. other rows are not used.
}

//NewIndex creates index according to given conf, overwrites existing files.
func NewIndex(docProt *cql.DocumentWithIdx, mainDir string, t0mCap, leafCap, intraCap int) (ind *Index, err error) {
	if err = indexWriteConf(mainDir, docProt); err != nil {
		return
	}
	// ensure per-index sub-directory exists
	indDir := filepath.Join(mainDir, docProt.Index)
	if err = os.MkdirAll(indDir, 0700); err != nil {
		return
	}
	ind = &Index{
		MainDir: mainDir,
		DocProt: docProt,
		bkds:    make(map[string]*bkdtree.BkdTree),
		frames:  make(map[string]*Frame),
	}
	var bkd *bkdtree.BkdTree
	for _, uintProp := range docProt.UintProps {
		t0mCap := t0mCap
		LeafCap := leafCap
		IntraCap := intraCap
		numDims := 1
		bytesPerDim := uintProp.ValLen
		dir := filepath.Join(indDir, uintProp.Name)
		prefix := uintProp.Name
		bkd, err = bkdtree.NewBkdTree(t0mCap, LeafCap, IntraCap, numDims, bytesPerDim, dir, prefix)
		if err != nil {
			return
		}
		ind.bkds[uintProp.Name] = bkd
	}
	var fm *Frame
	for _, strProp := range docProt.StrProps {
		dir := filepath.Join(indDir, strProp.Name)
		if fm, err = NewFrame(dir, docProt.Index, strProp.Name, true); err != nil {
			return
		}
		ind.frames[strProp.Name] = fm
	}
	dir := filepath.Join(indDir, LiveDocs)
	if fm, err = NewFrame(dir, docProt.Index, LiveDocs, true); err != nil {
		return
	}
	ind.liveDocs = fm
	return
}

//indexWriteConf persists conf to given path.
func indexWriteConf(mainDir string, docProt *cql.DocumentWithIdx) (err error) {
	if err = os.MkdirAll(mainDir, 0700); err != nil {
		err = errors.Wrap(err, "")
		return
	}
	fp := filepath.Join(mainDir, fmt.Sprintf("index_%s.json", docProt.Index))
	err = bkdtree.FileMarshal(fp, docProt)
	return
}

//indexReadConf parses conf
func indexReadConf(mainDir string, name string, docProt *cql.DocumentWithIdx) (err error) {
	fp := filepath.Join(mainDir, fmt.Sprintf("index_%s.json", name))
	err = bkdtree.FileUnmarshal(fp, docProt)
	return
}

//Destroy removes data and conf files on disk.
func (ind *Index) Destroy() (err error) {
	ind.rwlock.Lock()
	defer ind.rwlock.Unlock()
	if ind.liveDocs != nil {
		for _, bkd := range ind.bkds {
			if err = bkd.Destroy(); err != nil {
				return
			}
		}
		for _, fm := range ind.frames {
			if err = fm.Destroy(); err != nil {
				return
			}
		}
		if err = ind.liveDocs.Destroy(); err != nil {
			return
		}
		ind.bkds = nil
		ind.frames = nil
		ind.liveDocs = nil
	}

	paths := make([]string, 0)
	for _, uintProp := range ind.DocProt.UintProps {
		paths = append(paths, filepath.Join(ind.MainDir, uintProp.Name))
	}
	for _, strProp := range ind.DocProt.StrProps {
		paths = append(paths, filepath.Join(ind.MainDir, strProp.Name))
	}
	paths = append(paths, filepath.Join(ind.MainDir, LiveDocs))
	paths = append(paths, filepath.Join(ind.MainDir, fmt.Sprintf("index_%s.json", ind.DocProt.Index)))
	for _, fp := range paths {
		if err = os.RemoveAll(fp); err != nil {
			err = errors.Wrap(err, "")
		}
	}
	return
}

//NewIndexExt create index according to existing files.
func NewIndexExt(mainDir, name string) (ind *Index, err error) {
	docProt := &cql.DocumentWithIdx{}
	if err = indexReadConf(mainDir, name, docProt); err != nil {
		return
	}
	ind = &Index{
		MainDir: mainDir,
		DocProt: docProt,
	}
	err = ind.Open()
	return
}

//Open opens existing index. Assumes MainDir and DocProt is already populated.
func (ind *Index) Open() (err error) {
	ind.rwlock.Lock()
	defer ind.rwlock.Unlock()
	if ind.liveDocs != nil {
		//index is already open
		return
	}
	indDir := filepath.Join(ind.MainDir, ind.DocProt.Index)
	ind.bkds = make(map[string]*bkdtree.BkdTree)
	var bkd *bkdtree.BkdTree
	for _, uintProp := range ind.DocProt.UintProps {
		dir := filepath.Join(indDir, uintProp.Name)
		if bkd, err = bkdtree.NewBkdTreeExt(dir, uintProp.Name); err != nil {
			return
		}
		ind.bkds[uintProp.Name] = bkd
	}
	ind.frames = make(map[string]*Frame)
	var fm *Frame
	for _, strProp := range ind.DocProt.StrProps {
		dir := filepath.Join(indDir, strProp.Name)
		if fm, err = NewFrame(dir, ind.DocProt.Index, strProp.Name, false); err != nil {
			return
		}
		ind.frames[strProp.Name] = fm
	}
	dir := filepath.Join(indDir, LiveDocs)
	if fm, err = NewFrame(dir, ind.DocProt.Index, LiveDocs, false); err != nil {
		return
	}
	ind.liveDocs = fm
	return
}

//Close closes index
func (ind *Index) Close() (err error) {
	ind.rwlock.Lock()
	defer ind.rwlock.Unlock()
	if ind.liveDocs == nil {
		//index is already closed
		return
	}
	for _, bkd := range ind.bkds {
		if err = bkd.Close(); err != nil {
			return
		}
	}
	for _, fm := range ind.frames {
		if err = fm.Close(); err != nil {
			return
		}
	}
	if err = ind.liveDocs.Close(); err != nil {
		return
	}
	ind.bkds = nil
	ind.frames = nil
	ind.liveDocs = nil
	return
}

//Insert executes CqlInsert
func (ind *Index) Insert(doc *cql.DocumentWithIdx) (err error) {
	var bkd *bkdtree.BkdTree
	var fm *Frame
	var changed, ok bool
	ind.rwlock.RLock()
	defer ind.rwlock.RUnlock()
	//check if doc.DocID is already there before insertion.
	if changed, err = ind.liveDocs.setBit(0, doc.DocID); err != nil {
		return
	} else if !changed {
		err = errors.Errorf("document %v is alaredy there before insertion", doc.DocID)
		return
	}
	for _, uintProp := range doc.UintProps {
		if bkd, ok = ind.bkds[uintProp.Name]; !ok {
			err = errors.Errorf("property %v is missing at index spec, document %v, index spec %v", uintProp.Name, doc, ind.DocProt)
			return
		}
		p := bkdtree.Point{
			Vals:     []uint64{uintProp.Val},
			UserData: doc.DocID,
		}
		if err = bkd.Insert(p); err != nil {
			return
		}
	}
	for _, strProp := range doc.StrProps {
		if fm, ok = ind.frames[strProp.Name]; !ok {
			err = errors.Errorf("property %v is missing at index spec, document %v, index spec %v", strProp.Name, doc, ind.DocProt)
			return
		}
		if err = fm.ParseAndIndex(doc.DocID, strProp.Val); err != nil {
			return
		}
	}
	return
}

//Del executes CqlDel. Do mark-deletion only. The caller shall rebuild index in order to recycle disk space.
func (ind *Index) Del(docID uint64) (found bool, err error) {
	var changed bool
	ind.rwlock.RLock()
	defer ind.rwlock.RUnlock()
	if changed, err = ind.liveDocs.clearBit(0, docID); err != nil {
		return
	} else if !changed {
		err = errors.Errorf("document %v does not exist before deletion", docID)
		return
	}
	found = true
	return
}

//Select executes CqlSelect.
func (ind *Index) Select(q *cql.CqlSelect) (rb *pilosa.Bitmap, err error) {
	var fm *Frame
	var bkd *bkdtree.BkdTree
	var ok bool
	var prevDocs, docs *pilosa.Bitmap

	ind.rwlock.RLock()
	defer ind.rwlock.RUnlock()
	prevDocs = ind.liveDocs.row(0)
	if prevDocs.Count() == 0 {
		rb = prevDocs
		return
	}
	if len(q.StrPreds) != 0 {
		for _, strPred := range q.StrPreds {
			if fm, ok = ind.frames[strPred.Name]; !ok {
				err = errors.Errorf("property %s not found in index spec", strPred.Name)
				return
			}
			docs = fm.Query(strPred.ContWord)
			prevDocs = prevDocs.Intersect(docs)
			if prevDocs.Count() == 0 {
				rb = prevDocs
				return
			}
		}
	}

	if len(q.UintPreds) == 0 {
		rb = prevDocs
		return
	}

	visitor := &bkdVisitor{
		prevDocs: prevDocs,
		docs:     pilosa.NewBitmap(),
	}

	for _, uintPred := range q.UintPreds {
		if q.OrderBy == uintPred.Name {
			continue
		}
		if bkd, ok = ind.bkds[uintPred.Name]; !ok {
			err = errors.Errorf("property %s not found in index spec", uintPred.Name)
			return
		}
		visitor.lowPoint = bkdtree.Point{
			Vals: []uint64{uintPred.Low},
		}
		visitor.highPoint = bkdtree.Point{
			Vals: []uint64{uintPred.High},
		}
		if err = bkd.Intersect(visitor); err != nil {
			return
		}
		visitor.prevDocs = visitor.docs
		if visitor.prevDocs.Count() == 0 {
			rb = visitor.prevDocs
			return
		}
	}

	if q.OrderBy != "" {
		var uintPred cql.UintPred
		if uintPred, ok = q.UintPreds[q.OrderBy]; !ok {
			err = errors.Errorf("invalid ORDERBY %s", q.OrderBy)
			return
		}
		bkd, ok = ind.bkds[uintPred.Name]
		if !ok {
			err = errors.Errorf("property %s not found in index spec", uintPred.Name)
			return
		}
		visitor.lowPoint = bkdtree.Point{
			Vals: []uint64{uintPred.Low},
		}
		visitor.highPoint = bkdtree.Point{
			Vals: []uint64{uintPred.High},
		}
		visitor.limit = q.Limit
		visitor.h = &bkdtree.PointMaxHeap{}
		if err = bkd.Intersect(visitor); err != nil {
			return
		}
		rb = pilosa.NewBitmap()
		for _, point := range *visitor.h {
			rb.SetBit(point.UserData)
		}
		visitor.prevDocs = rb
	}

	rb = visitor.prevDocs
	return
}

type bkdVisitor struct {
	lowPoint  bkdtree.Point
	highPoint bkdtree.Point
	prevDocs  *pilosa.Bitmap
	docs      *pilosa.Bitmap
	limit     int
	h         *bkdtree.PointMaxHeap
}

func (v *bkdVisitor) GetLowPoint() bkdtree.Point { return v.lowPoint }

func (v *bkdVisitor) GetHighPoint() bkdtree.Point { return v.highPoint }

func (v *bkdVisitor) VisitPoint(point bkdtree.Point) {
	docID := point.UserData
	if v.prevDocs == nil || v.prevDocs.Contains(docID) {
		if v.limit == 0 {
			v.docs.SetBit(docID)
		} else {
			if len(*v.h) < v.limit {
				heap.Push(v.h, point)
			} else if point.LessThan((*v.h)[0]) {
				(*v.h)[0] = point
				heap.Fix(v.h, 0)
			}
		}
	}
}
