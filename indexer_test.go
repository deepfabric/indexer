package indexer

import (
	"fmt"
	"testing"
	"time"

	"github.com/RoaringBitmap/roaring"
	"github.com/juju/testing/checkers"

	"github.com/deepfabric/indexer/cql"
)

func newTestIndex1() (ind *Index) {
	ind = &Index{
		Conf: IndexConf{
			bkd: BkdConf{
				t0mCap:      1000,
				leafCap:     100,
				intraCap:    4,
				dir:         "/tmp",
				cptInterval: 30 * time.Minute,
			},
			cap: 100000,
		},
	}
	return
}

func newTestIndex2() (ind *Index) {
	ind = &Index{
		Conf: IndexConf{
			bkd: BkdConf{
				t0mCap:      1000,
				leafCap:     100,
				intraCap:    4,
				dir:         "/tmp",
				cptInterval: 30 * time.Minute,
			},
			cap: 100000,
		},
		def: cql.IndexDef{
			DocumentWithIdx: cql.DocumentWithIdx{
				Document: cql.Document{
					DocID: 0,
					UintProps: []cql.UintProp{
						cql.UintProp{
							Name:   "object",
							ValLen: 8,
							Val:    0,
						},
						cql.UintProp{
							Name:   "price",
							ValLen: 4,
							Val:    0,
						},
						cql.UintProp{
							Name:   "number",
							ValLen: 4,
							Val:    0,
						},
						cql.UintProp{
							Name:   "date",
							ValLen: 8,
							Val:    0,
						},
					},
				},
				Index: "orders",
			},
		},
	}
	return
}

//TESTCASE: normal operation sequence: create, insert, del, destroy
func TestIndexNormal(t *testing.T) {
	var err error
	var isEqual bool
	var found bool

	ind1 := newTestIndex1()
	ind2 := newTestIndex2()
	q := &cql.CqlCreate{
		IndexDef: ind2.def,
	}
	if err = ind1.Create(q); err != nil {
		t.Fatalf("%+v", err)
	}
	if isEqual, err = checkers.DeepEqual(ind1.def, ind2.def); !isEqual {
		t.Fatalf("incorrect result of (*Index).Create, %+v", err)
	}

	doc := cql.DocumentWithIdx{
		Document: cql.Document{
			DocID: 0,
			UintProps: []cql.UintProp{
				cql.UintProp{
					Name:   "object",
					ValLen: 8,
					Val:    0,
				},
				cql.UintProp{
					Name:   "price",
					ValLen: 4,
					Val:    0,
				},
				cql.UintProp{
					Name:   "number",
					ValLen: 4,
					Val:    0,
				},
				cql.UintProp{
					Name:   "date",
					ValLen: 8,
					Val:    0,
				},
			},
		},
		Index: "orders",
	}

	for i := 0; i < 2000; i++ {
		doc.DocID = uint64(i)
		for j := 0; j < len(doc.UintProps); j++ {
			doc.UintProps[j].Val = uint64(i * (j + 1))
		}
		ins := &cql.CqlInsert{
			DocumentWithIdx: doc,
		}
		if err = ind1.Insert(ins); err != nil {
			t.Fatalf("%+v", err)
		}
	}

	var rb *roaring.Bitmap
	cs := &cql.CqlSelect{
		Index: doc.Index,
		UintPreds: map[string]cql.UintPred{
			"price": cql.UintPred{
				Name: "price",
				Low:  uint64(30),
				High: uint64(100),
			},
		},
	}
	if rb, err = ind1.Select(cs); err != nil {
		t.Fatalf("%+v", err)
	}
	fmt.Println(rb.String())

	for i := 0; i < 2000; i++ {
		doc.DocID = uint64(i)
		for j := 0; j < len(doc.UintProps); j++ {
			doc.UintProps[j].Val = uint64(i * (j + 1))
		}
		del := &cql.CqlDel{
			DocumentWithIdx: doc,
		}
		if found, err = ind1.Del(del); err != nil {
			t.Fatalf("%+v", err)
		} else if !found {
			t.Fatalf("document %v not found", doc)
		}
	}

	q2 := &cql.CqlDestroy{
		Index: "orders",
	}
	if err = ind1.Destroy(q2); err != nil {
		t.Fatalf("%+v", err)
	}
}
