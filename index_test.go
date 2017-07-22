package indexer

import (
	"fmt"
	"testing"
	"time"

	"github.com/deepfabric/pilosa"
	"github.com/juju/testing/checkers"

	"github.com/deepfabric/indexer/cql"
)

const (
	BkdCapTest      = 10000
	T0mCapTest      = 1000
	LeafCapTest     = 100
	IntraCapTest    = 4
	CptIntervalTest = 30 * time.Minute
)

func newDocProt() *cql.DocumentWithIdx {
	return &cql.DocumentWithIdx{
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
}

//TESTCASE: normal operation sequence: create, insert, del, destroy
func TestIndexNormal(t *testing.T) {
	var err error
	var ind *Index
	var isEqual bool
	var found bool

	docProt := newDocProt()
	ind, err = NewIndex(docProt, "/tmp", BkdCapTest, T0mCapTest, LeafCapTest, IntraCapTest, CptIntervalTest)
	if isEqual, err = checkers.DeepEqual(ind.DocProt, *docProt); !isEqual {
		t.Fatalf("incorrect result of (*Index).Create, %+v", err)
	}

	doc := *docProt
	for i := 0; i < BkdCapTest; i++ {
		doc.DocID = uint64(i)
		for j := 0; j < len(doc.UintProps); j++ {
			doc.UintProps[j].Val = uint64(i * (j + 1))
		}
		if err = ind.Insert(&doc); err != nil {
			t.Fatalf("%+v", err)
		}
	}

	var rb *pilosa.Bitmap
	low := 30
	high := 600
	cs := &cql.CqlSelect{
		Index: doc.Index,
		UintPreds: map[string]cql.UintPred{
			"price": cql.UintPred{
				Name: "price",
				Low:  uint64(low),
				High: uint64(high),
			},
		},
	}
	if rb, err = ind.Select(cs); err != nil {
		t.Fatalf("%+v", err)
	}
	fmt.Println(rb.Bits())
	// low <= 2*i <= high, (low+1)/2 <= i <= high/2
	want := high/2 - (low+1)/2 + 1
	if rb.Count() != uint64(want) {
		t.Fatalf("incorrect number of matches, have %d, want %d", rb.Count(), want)
	}

	cs.OrderBy = "price"
	cs.Limit = 20
	if rb, err = ind.Select(cs); err != nil {
		t.Fatalf("%+v", err)
	}
	fmt.Println(rb.Bits())
	want = cs.Limit
	if rb.Count() != uint64(want) {
		t.Fatalf("incorrect number of matches, have %d, want %d", rb.Count(), want)
	}

	for i := 0; i < BkdCapTest; i++ {
		doc.DocID = uint64(i)
		for j := 0; j < len(doc.UintProps); j++ {
			doc.UintProps[j].Val = uint64(i * (j + 1))
		}
		if found, err = ind.Del(&doc); err != nil {
			t.Fatalf("%+v", err)
		} else if !found {
			t.Fatalf("document %v not found", doc)
		}
	}

	if err = ind.Destroy(); err != nil {
		t.Fatalf("%+v", err)
	}
}

func TestIndexOpenClose(t *testing.T) {
	var err error
	var ind, ind2 *Index
	var isEqual bool

	//create index
	docProt := newDocProt()
	ind, err = NewIndex(docProt, "/tmp", BkdCapTest, T0mCapTest, LeafCapTest, IntraCapTest, CptIntervalTest)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	//insert documents
	doc := *docProt
	for i := 0; i < BkdCapTest; i++ {
		doc.DocID = uint64(i)
		for j := 0; j < len(doc.UintProps); j++ {
			doc.UintProps[j].Val = uint64(i * (j + 1))
		}
		if err = ind.Insert(&doc); err != nil {
			t.Fatalf("%+v", err)
		}
	}

	//close index
	if err = ind.Close(); err != nil {
		t.Fatalf("%+v", err)
	}

	//open index
	if err = ind.Open(BkdCapTest, CptIntervalTest); err != nil {
		t.Fatalf("%+v", err)
	}

	//close index
	if err = ind.Close(); err != nil {
		t.Fatalf("%+v", err)
	}

	//open index with another Index object. This occurs when program restart.
	ind2, err = NewIndexExt("/tmp", "orders", BkdCapTest, CptIntervalTest)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	//verify DocProtkeeps unchanged
	docProt = newDocProt() //docProt changes after above Insert step. have to restore it
	if isEqual, err = checkers.DeepEqual(ind2.DocProt, ind.DocProt); !isEqual {
		fmt.Printf("have %v\n", ind2.DocProt)
		fmt.Printf("want %v\n", ind.DocProt)
		t.Fatalf("index DocProt %+v", err)
	}

	//close index
	if err = ind2.Close(); err != nil {
		t.Fatalf("%+v", err)
	}
}
