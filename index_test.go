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
	BkdCapTest = 10000
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

func newTestIndex() (ind *Index) {
	ind = &Index{
		Def: IndexDef{
			Conf: IndexerConf{
				Dir:         "/tmp",
				Cap:         BkdCapTest,
				T0mCap:      1000,
				LeafCap:     100,
				IntraCap:    4,
				CptInterval: 30 * time.Minute,
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

	docProt := newDocProt()
	ind := newTestIndex()
	q := &cql.CqlCreate{
		DocumentWithIdx: *docProt,
	}
	if err = ind.Create(q); err != nil {
		t.Fatalf("%+v", err)
	}
	if isEqual, err = checkers.DeepEqual(ind.Def.DocProt, *docProt); !isEqual {
		t.Fatalf("incorrect result of (*Index).Create, %+v", err)
	}

	doc := *docProt
	for i := 0; i < BkdCapTest; i++ {
		doc.DocID = uint64(i)
		for j := 0; j < len(doc.UintProps); j++ {
			doc.UintProps[j].Val = uint64(i * (j + 1))
		}
		ins := &cql.CqlInsert{
			DocumentWithIdx: doc,
		}
		if err = ind.Insert(ins); err != nil {
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
		del := &cql.CqlDel{
			DocumentWithIdx: doc,
		}
		if found, err = ind.Del(del); err != nil {
			t.Fatalf("%+v", err)
		} else if !found {
			t.Fatalf("document %v not found", doc)
		}
	}

	q2 := &cql.CqlDestroy{
		Index: "orders",
	}
	if err = ind.Destroy(q2); err != nil {
		t.Fatalf("%+v", err)
	}
}

func TestIndexOpenClose(t *testing.T) {
	var err error
	var isEqual bool

	//create index
	docProt := newDocProt()
	ind := newTestIndex()
	q := &cql.CqlCreate{
		DocumentWithIdx: *docProt,
	}
	if err = ind.Create(q); err != nil {
		t.Fatalf("%+v", err)
	}

	//insert documents
	doc := *docProt
	for i := 0; i < BkdCapTest; i++ {
		doc.DocID = uint64(i)
		for j := 0; j < len(doc.UintProps); j++ {
			doc.UintProps[j].Val = uint64(i * (j + 1))
		}
		ins := &cql.CqlInsert{
			DocumentWithIdx: doc,
		}
		if err = ind.Insert(ins); err != nil {
			t.Fatalf("%+v", err)
		}
	}

	//close index
	if err = ind.Close(); err != nil {
		t.Fatalf("%+v", err)
	}

	//open index
	dir := ind.Def.Conf.Dir
	name := doc.Index
	if err = ind.Open(dir, name); err != nil {
		t.Fatalf("%+v", err)
	}

	//close index
	if err = ind.Close(); err != nil {
		t.Fatalf("%+v", err)
	}

	//open index with another Index object. This occurs when program restart.
	ind2 := newTestIndex()
	if err = ind2.Open(dir, name); err != nil {
		t.Fatalf("%+v", err)
	}

	//verify index definion keeps unchanged
	if isEqual, err = checkers.DeepEqual(ind2.Def.Conf, ind.Def.Conf); !isEqual {
		fmt.Printf("have %v\n", ind2.Def)
		fmt.Printf("want %v\n", ind.Def)
		t.Fatalf("index Def.Conf %+v", err)
	}
	docProt = newDocProt() //docProt changes after above Insert step. have to restore it
	if isEqual, err = checkers.DeepEqual(ind2.Def.DocProt, *docProt); !isEqual {
		fmt.Printf("have %v\n", ind2.Def.DocProt)
		fmt.Printf("want %v\n", *docProt)
		t.Fatalf("index Def.DocProt %+v", err)
	}

	//close index
	if err = ind2.Close(); err != nil {
		t.Fatalf("%+v", err)
	}
}
