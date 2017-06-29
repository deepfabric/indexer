package indexer

import (
	"container/heap"
	"time"

	"github.com/pkg/errors"

	"github.com/RoaringBitmap/roaring"
	"github.com/deepfabric/bkdtree"
	"github.com/deepfabric/indexer/cql"
)

//IndexConf originate from config file. It's comman to all indices.
type IndexConf struct {
	//generic items
	Dir string
	Cap int //upper limit of number of documents

	//BKD specific items
	T0mCap      int
	LeafCap     int
	IntraCap    int
	CptInterval time.Duration
}

//IndexDef will be persisted to an index-specific file in order to support reopening an index
type IndexDef struct {
	Conf    IndexConf
	DocProt cql.DocumentWithIdx //document prototype
}

type Index struct {
	Def  IndexDef
	bkds map[string]*bkdtree.BkdTree
}

// Create execute CqlCreate. Assume ind.Def.Conf has been populated.
func (ind *Index) Create(q *cql.CqlCreate) (err error) {
	ind.Def.DocProt = q.DocumentWithIdx
	ind.bkds = make(map[string]*bkdtree.BkdTree)
	var bkd *bkdtree.BkdTree
	for _, uintProp := range q.UintProps {
		t0mCap := ind.Def.Conf.T0mCap
		bkdCap := ind.Def.Conf.Cap
		numDims := 1
		bytesPerDim := uintProp.ValLen
		LeafCap := ind.Def.Conf.LeafCap
		IntraCap := ind.Def.Conf.IntraCap
		dir := ind.Def.Conf.Dir
		prefix := uintProp.Name
		cptInterval := ind.Def.Conf.CptInterval
		bkd, err = bkdtree.NewBkdTree(t0mCap, bkdCap, numDims, bytesPerDim, LeafCap, IntraCap, dir, prefix, cptInterval)
		if err != nil {
			return
		}
		ind.bkds[uintProp.Name] = bkd
	}
	return
}

func (ind *Index) Destroy(q *cql.CqlDestroy) (err error) {
	for prop, bkd := range ind.bkds {
		if err = bkd.Destroy(); err != nil {
			return
		}
		delete(ind.bkds, prop)
	}
	return
}

func (ind *Index) Insert(q *cql.CqlInsert) (err error) {
	var bkd *bkdtree.BkdTree
	var ok bool
	if len(ind.bkds) != len(q.UintProps) {
		err = errors.Errorf("incorrect length of CqlInsert.UintProps, have %d, want %d", len(q.UintProps), len(ind.bkds))
		return
	}
	for _, uintProp := range q.UintProps {
		bkd, ok = ind.bkds[uintProp.Name]
		if !ok {
			err = errors.Errorf("property %s not found in index spec", uintProp.Name)
			return
		}
		p := bkdtree.Point{
			Vals:     []uint64{uintProp.Val},
			UserData: q.DocID,
		}
		if err = bkd.Insert(p); err != nil {
			return
		}
	}
	return
}

func (ind *Index) Del(q *cql.CqlDel) (found bool, err error) {
	var bkd *bkdtree.BkdTree
	var ok bool
	if len(ind.bkds) != len(q.UintProps) {
		err = errors.Errorf("incorrect length of CqlDel.UintProps, have %d, want %d", len(q.UintProps), len(ind.bkds))
		return
	}
	for _, uintProp := range q.UintProps {
		bkd, ok = ind.bkds[uintProp.Name]
		if !ok {
			err = errors.Errorf("property %s not found in index spec", uintProp.Name)
			return
		}
		p := bkdtree.Point{
			Vals:     []uint64{uintProp.Val},
			UserData: q.DocID,
		}
		if found, err = bkd.Erase(p); err != nil {
			return
		}
	}
	return
}

func (ind *Index) Select(q *cql.CqlSelect) (rb *roaring.Bitmap, err error) {
	var bkd *bkdtree.BkdTree
	var ok bool
	visitor := &bkdVisitor{
		docs: roaring.NewBitmap(),
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
		rb = roaring.NewBitmap()
		for _, point := range *visitor.h {
			rb.Add(uint32(point.UserData))
		}
		visitor.prevDocs = rb
	}

	rb = visitor.prevDocs
	return
}

type bkdVisitor struct {
	lowPoint  bkdtree.Point
	highPoint bkdtree.Point
	prevDocs  *roaring.Bitmap
	docs      *roaring.Bitmap
	limit     int
	h         *bkdtree.PointMaxHeap
}

func (v *bkdVisitor) GetLowPoint() bkdtree.Point { return v.lowPoint }

func (v *bkdVisitor) GetHighPoint() bkdtree.Point { return v.highPoint }

func (v *bkdVisitor) VisitPoint(point bkdtree.Point) {
	docID := uint32(point.UserData)
	//TODO: add uint64 support for roaring.Bitmap?
	if v.prevDocs == nil || v.prevDocs.Contains(docID) {
		if v.limit == 0 {
			v.docs.Add(docID)
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
