package indexer

import (
	"container/heap"
	"fmt"
	"path/filepath"
	"time"

	"os"

	"github.com/deepfabric/bkdtree"
	"github.com/deepfabric/indexer/cql"
	"github.com/deepfabric/pilosa"
	"github.com/pkg/errors"
)

//Index is created by CqlCreate
type Index struct {
	MainDir string
	DocProt *cql.DocumentWithIdx //document prototype. persisted to an index-specific file
	bkds    map[string]*bkdtree.BkdTree
}

//NewIndex creates index according to given conf, overwrites existing files.
func NewIndex(docProt *cql.DocumentWithIdx, mainDir string, cap, t0mCap, leafCap, intraCap int, cptInterval time.Duration) (ind *Index, err error) {
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
	}
	var bkd *bkdtree.BkdTree
	for _, uintProp := range docProt.UintProps {
		t0mCap := t0mCap
		bkdCap := cap
		numDims := 1
		bytesPerDim := uintProp.ValLen
		LeafCap := leafCap
		IntraCap := intraCap
		dir := filepath.Join(indDir, uintProp.Name)
		prefix := uintProp.Name
		bkd, err = bkdtree.NewBkdTree(t0mCap, bkdCap, numDims, bytesPerDim, LeafCap, IntraCap, dir, prefix, cptInterval)
		if err != nil {
			return
		}
		ind.bkds[uintProp.Name] = bkd
	}
	return
}

//indexWriteConf persists conf to given path.
func indexWriteConf(mainDir string, docProt *cql.DocumentWithIdx) (err error) {
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
	for prop, bkd := range ind.bkds {
		if err = bkd.Destroy(); err != nil {
			return
		}
		delete(ind.bkds, prop)
	}
	fp := filepath.Join(ind.MainDir, fmt.Sprintf("index_%s.json", ind.DocProt.Index))
	err = os.Remove(fp)
	return
}

//NewIndexExt create index according to existing files. Assumes ind.DocProt is already populated.
func NewIndexExt(mainDir, name string, cap int, cptInterval time.Duration) (ind *Index, err error) {
	docProt := &cql.DocumentWithIdx{}
	if err = indexReadConf(mainDir, name, docProt); err != nil {
		return
	}
	ind = &Index{
		MainDir: mainDir,
		DocProt: docProt,
	}
	err = ind.Open(cap, cptInterval)
	return
}

//Open opens existing index. Assumes MainDir and DocProt is already populated.
func (ind *Index) Open(cap int, cptInterval time.Duration) (err error) {
	if ind.bkds != nil {
		panic("index is already open")
	}
	indDir := filepath.Join(ind.MainDir, ind.DocProt.Index)
	ind.bkds = make(map[string]*bkdtree.BkdTree)
	var bkd *bkdtree.BkdTree
	for _, uintProp := range ind.DocProt.UintProps {
		dir := filepath.Join(indDir, uintProp.Name)
		if bkd, err = bkdtree.NewBkdTreeExt(dir, uintProp.Name, cap, cptInterval); err != nil {
			return
		}
		ind.bkds[uintProp.Name] = bkd
	}
	return
}

//Close closes index
func (ind *Index) Close() (err error) {
	for _, bkd := range ind.bkds {
		if err = bkd.Close(); err != nil {
			return
		}
	}
	ind.bkds = nil
	return
}

//Insert executes CqlInsert
func (ind *Index) Insert(doc *cql.DocumentWithIdx) (err error) {
	var bkd *bkdtree.BkdTree
	var ok bool
	if len(ind.bkds) != len(doc.UintProps) {
		err = errors.Errorf("incorrect length of CqlInsert.UintProps, have %d, want %d", len(doc.UintProps), len(ind.bkds))
		return
	}
	for _, uintProp := range doc.UintProps {
		bkd, ok = ind.bkds[uintProp.Name]
		if !ok {
			err = errors.Errorf("property %s not found in index spec", uintProp.Name)
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
	return
}

//Del executes CqlDel.
func (ind *Index) Del(doc *cql.DocumentWithIdx) (found bool, err error) {
	var bkd *bkdtree.BkdTree
	var ok bool
	if len(ind.bkds) != len(doc.UintProps) {
		err = errors.Errorf("incorrect length of CqlDel.UintProps, have %d, want %d", len(doc.UintProps), len(ind.bkds))
		return
	}
	for _, uintProp := range doc.UintProps {
		bkd, ok = ind.bkds[uintProp.Name]
		if !ok {
			err = errors.Errorf("property %s not found in index spec", uintProp.Name)
			return
		}
		p := bkdtree.Point{
			Vals:     []uint64{uintProp.Val},
			UserData: doc.DocID,
		}
		if found, err = bkd.Erase(p); err != nil {
			return
		}
	}
	return
}

//Select executes CqlSelect.
func (ind *Index) Select(q *cql.CqlSelect) (rb *pilosa.Bitmap, err error) {
	var bkd *bkdtree.BkdTree
	var ok bool
	visitor := &bkdVisitor{
		docs: pilosa.NewBitmap(),
	}

	for _, uintPred := range q.UintPreds {
		if q.OrderBy == uintPred.Name {
			continue
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
		err = bkd.Intersect(visitor)
		if err != nil {
			return
		}
		visitor.prevDocs = visitor.docs
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
		err = bkd.Intersect(visitor)
		if err != nil {
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
	//TODO: add uint64 support for roaring.Bitmap?
	if v.prevDocs == nil || v.prevDocs.Contains(docID) {
		if v.limit == 0 {
			v.docs.SetBit(docID)
		} else {
			if len(*v.h) < v.limit {
				v.h.Push(point)
			} else if point.LessThan((*v.h)[0]) {
				(*v.h)[0] = point
				heap.Fix(v.h, 0)
			}
		}
	}
}
