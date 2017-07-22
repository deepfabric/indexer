package indexer

import (
	"fmt"
	"testing"
	"time"

	"github.com/deepfabric/indexer/cql"
	"github.com/deepfabric/pilosa"
)

const ()

func newDocProt1() *cql.DocumentWithIdx {
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

func newDocProt2() *cql.DocumentWithIdx {
	return &cql.DocumentWithIdx{
		Document: cql.Document{
			DocID: 0,
			UintProps: []cql.UintProp{
				cql.UintProp{
					Name:   "ip",
					ValLen: 4,
					Val:    0,
				},
				cql.UintProp{
					Name:   "created",
					ValLen: 8,
					Val:    0,
				},
				cql.UintProp{
					Name:   "updated",
					ValLen: 8,
					Val:    0,
				},
			},
		},
		Index: "addrs",
	}
}

//TESTCASE: normal operation sequence: new_empty, create, insert, close, new_existing, query, del, destroy
func TestIndexerNormal(t *testing.T) {
	var err error
	var docProt *cql.DocumentWithIdx
	var ir, ir2 *Indexer
	var found bool

	conf := Conf{
		Cap:         10000,
		T0mCap:      1000,
		LeafCap:     100,
		IntraCap:    4,
		CptInterval: 30 * time.Second,
	}

	//create empty indexer
	if ir, err = NewIndexer("/tmp", &conf, true); err != nil {
		t.Fatalf("%+v", err)
	}

	//create index 1
	docProt = newDocProt1()
	if err = ir.CreateIndex(docProt); err != nil {
		return
	}

	//create index 2
	docProt = newDocProt2()
	if err = ir.CreateIndex(docProt); err != nil {
		return
	}

	//insert documents
	for i := 0; i < conf.Cap; i++ {
		doc := newDocProt1()
		doc.DocID = uint64(i)
		for j := 0; j < len(doc.UintProps); j++ {
			doc.UintProps[j].Val = uint64(i * (j + 1))
		}
		if err = ir.Insert(doc); err != nil {
			t.Fatalf("%+v", err)
		}
	}

	//close indexer
	if err = ir.Close(); err != nil {
		t.Fatalf("%+v", err)
	}

	//create indexer with existing data
	if ir2, err = NewIndexer("/tmp", &conf, false); err != nil {
		t.Fatalf("%+v", err)
	}

	//query
	var rb *pilosa.Bitmap
	low := 30
	high := 600
	cs := &cql.CqlSelect{
		Index: "orders",
		UintPreds: map[string]cql.UintPred{
			"price": cql.UintPred{
				Name: "price",
				Low:  uint64(low),
				High: uint64(high),
			},
		},
	}
	if rb, err = ir2.Select(cs); err != nil {
		t.Fatalf("%+v", err)
	}
	fmt.Println(rb.Bits())

	//delete documents
	for i := 0; i < BkdCapTest; i++ {
		doc := newDocProt1()
		doc.DocID = uint64(i)
		for j := 0; j < len(doc.UintProps); j++ {
			doc.UintProps[j].Val = uint64(i * (j + 1))
		}
		if found, err = ir2.Del(doc); err != nil {
			t.Fatalf("%+v", err)
		} else if !found {
			t.Fatalf("document %v not found", doc)
		}
	}

	//destroy index 1
	if err = ir2.DestroyIndex("orders"); err != nil {
		t.Fatalf("%+v", err)
	}

	//destroy index 2
	if err = ir2.DestroyIndex("addrs"); err != nil {
		t.Fatalf("%+v", err)
	}
}

func TestIndexerOpenClose(t *testing.T) {
	var err error
	var ir *Indexer

	conf := Conf{
		Cap:         10000,
		T0mCap:      1000,
		LeafCap:     100,
		IntraCap:    4,
		CptInterval: 30 * time.Second,
	}

	//create indexer
	if ir, err = NewIndexer("/tmp", &conf, true); err != nil {
		t.Fatalf("%+v", err)
	}

	//create index
	docProt := newDocProt1()
	if err = ir.CreateIndex(docProt); err != nil {
		return
	}

	//close indexer
	if err = ir.Close(); err != nil {
		t.Fatalf("%+v", err)
	}

	//open indexer
	if err = ir.Open(); err != nil {
		t.Fatalf("%+v", err)
	}

	//close indexer
	if err = ir.Close(); err != nil {
		t.Fatalf("%+v", err)
	}
}
