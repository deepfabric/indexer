package indexer

import (
	"time"

	"github.com/pkg/errors"

	"github.com/RoaringBitmap/roaring"
	"github.com/deepfabric/bkdtree"
	"github.com/deepfabric/indexer/cql"
)

type BkdConf struct {
	t0mCap, leafCap, intraCap int
	dir                       string
	cptInterval               time.Duration
}

type IndexConf struct {
	bkd BkdConf
	cap int //upper limit of number of documents
}

type Index struct {
	Conf IndexConf
	def  cql.IndexDef
	bkds map[string]*bkdtree.BkdTree
}

func (ind *Index) Create(q *cql.CqlCreate) (err error) {
	ind.def = q.IndexDef
	ind.bkds = make(map[string]*bkdtree.BkdTree)
	var bkd *bkdtree.BkdTree
	for _, uintProp := range q.UintProps {
		t0mCap := ind.Conf.bkd.t0mCap
		bkdCap := ind.Conf.cap
		numDims := 1
		bytesPerDim := uintProp.ValLen
		leafCap := ind.Conf.bkd.leafCap
		intraCap := ind.Conf.bkd.intraCap
		dir := ind.Conf.bkd.dir
		prefix := uintProp.Name
		cptInterval := ind.Conf.bkd.cptInterval
		bkd, err = bkdtree.NewBkdTree(t0mCap, bkdCap, numDims, bytesPerDim, leafCap, intraCap, dir, prefix, cptInterval)
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
	rb = visitor.prevDocs
	return
}

type bkdVisitor struct {
	lowPoint  bkdtree.Point
	highPoint bkdtree.Point
	prevDocs  *roaring.Bitmap
	docs      *roaring.Bitmap
}

func (v *bkdVisitor) GetLowPoint() bkdtree.Point { return v.lowPoint }

func (v *bkdVisitor) GetHighPoint() bkdtree.Point { return v.highPoint }

func (v *bkdVisitor) VisitPoint(point bkdtree.Point) {
	docID := uint32(point.UserData)
	//TODO: add uint64 support for roaring.Bitmap?
	if v.prevDocs == nil || v.prevDocs.Contains(docID) {
		v.docs.Add(docID)
	}
}
