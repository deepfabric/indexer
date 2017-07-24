package indexer

import (
	"fmt"
	"testing"
	"time"

	"github.com/deepfabric/indexer/cql"
	"github.com/deepfabric/pilosa"
	"github.com/juju/testing/checkers"
)

const (
	BkdCapTest      = 10000
	T0mCapTest      = 100
	LeafCapTest     = 10
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
			StrProps: []cql.StrProp{
				cql.StrProp{
					Name: "description",
					Val:  "",
				},
				cql.StrProp{
					Name: "note",
					Val:  "",
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
	var bits map[uint64][]uint64

	docProt := newDocProt()
	if ind, err = NewIndex(docProt, "/tmp/index_test", BkdCapTest, T0mCapTest, LeafCapTest, IntraCapTest, CptIntervalTest); err != nil {
		t.Fatalf("incorrect result of NewIndex, %+v", err)
	}
	if isEqual, err = checkers.DeepEqual(ind.DocProt, docProt); !isEqual {
		t.Fatalf("incorrect result of NewIndex, %+v", err)
	}

	for i := 0; i < BkdCapTest; i++ {
		doc := newDocProt()
		doc.DocID = uint64(i)
		for j := 0; j < len(doc.UintProps); j++ {
			doc.UintProps[j].Val = uint64(i * (j + 1))
		}
		for j := 0; j < len(doc.StrProps); j++ {
			doc.StrProps[j].Val = fmt.Sprintf("%d_%d and some random text", i, j)
		}
		if err = ind.Insert(doc); err != nil {
			t.Fatalf("%+v", err)
		}
	}

	// query numerical range
	var rb *pilosa.Bitmap
	low := 30
	high := 600
	cs := &cql.CqlSelect{
		Index: docProt.Index,
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

	// query numerical range + order by + text
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

	// dump bits
	for name, frame := range ind.frames {
		var termID uint64
		if termID, found = frame.td.GetTermID("17_1"); !found {
			continue
		}
		if bits, err = frame.Bits(pilosa.ViewInverse); err != nil {
			t.Fatalf("%+v", err)
		}
		//fmt.Printf("frmae %v bits: %v\n", name, bits)
		fmt.Printf("frmae %v bits[%v]: %v\n", name, termID, bits[termID])
	}

	// query numerical range + text
	cs.StrPreds = map[string]cql.StrPred{
		"note": cql.StrPred{
			Name:     "note",
			ContWord: "17_1",
			//			ContWord: "random",
		},
	}
	cs.Limit = 20
	if rb, err = ind.Select(cs); err != nil {
		t.Fatalf("%+v", err)
	}
	fmt.Println(rb.Bits())
	want = 1
	if rb.Count() != uint64(want) {
		t.Fatalf("incorrect number of matches, have %d, want %d", rb.Count(), want)
	}

	for i := 0; i < BkdCapTest; i++ {
		doc := newDocProt()
		doc.DocID = uint64(i)
		for j := 0; j < len(doc.UintProps); j++ {
			doc.UintProps[j].Val = uint64(i * (j + 1))
		}
		if found, err = ind.Del(doc); err != nil {
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
	ind, err = NewIndex(docProt, "/tmp/index_test", BkdCapTest, T0mCapTest, LeafCapTest, IntraCapTest, CptIntervalTest)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	//insert documents
	for i := 0; i < BkdCapTest; i++ {
		doc := newDocProt()
		doc.DocID = uint64(i)
		for j := 0; j < len(doc.UintProps); j++ {
			doc.UintProps[j].Val = uint64(i * (j + 1))
		}
		if err = ind.Insert(doc); err != nil {
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
	ind2, err = NewIndexExt("/tmp/index_test", "orders", BkdCapTest, CptIntervalTest)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	//verify DocProt keeps unchanged
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
