package indexer

import (
	"fmt"
	"testing"

	"github.com/juju/testing/checkers"
	"github.com/pilosa/pilosa"
)

func TestParseAndIndex(t *testing.T) {
	var err error
	var found bool
	var f *Frame
	var terms []string

	//TESTCASE: query and insert term to an empty dict
	if f, err = NewFrame("/tmp", "i", "f"); err != nil {
		t.Fatalf("%+v", err)
	}
	text := "Go's standard library does not have a function solely intended to check if a file exists or not (like Python's os.path.exists). What is the idiomatic way to do it?"
	if err = f.ParseAndIndex(3, text); err != nil {
		t.Fatalf("%+v", err)
	}
	fmt.Printf("termdict size: %d\n", f.td.Count())

	terms = []string{"Go's", "it?"}
	for _, term := range terms {
		if _, found = f.td.GetTermID(term); !found {
			t.Fatalf("Term %s not found, want found", term)
		}
	}

	terms = []string{"java", "php"}
	for _, term := range terms {
		if _, found = f.td.GetTermID(term); found {
			t.Fatalf("Term %s found, want not-found", term)
		}
	}
}
func TestQuery(t *testing.T) {
	var err error
	var f *Frame
	var terms []string
	var bm *pilosa.Bitmap
	var isEqual bool

	//TESTCASE: query and insert term to an empty dict
	if f, err = NewFrame("/tmp", "i", "f"); err != nil {
		t.Fatalf("%+v", err)
	}
	docIDs := []uint64{1, 10}
	texts := []string{
		"Go's standard library does not have a function solely intended to check if a file exists or not (like Python's os.path.exists). What is the idiomatic way to do it?",
		"This is a listing of successful results of all the various data storage and processing system benchmarks I've conducted using the dataset produced in the Billion Taxi Rides in Redshift blog post. The dataset itself has 1.1 billion records, 51 columns and takes up about 500 GB of disk space uncompressed.",
	}
	for i := 0; i < len(docIDs); i++ {
		if err = f.ParseAndIndex(docIDs[i], texts[i]); err != nil {
			t.Fatalf("%+v", err)
		}
	}
	fmt.Printf("termdict size: %d\n", f.td.Count())

	terms = []string{"the", "disk"}
	expDocIDs := [][]uint64{[]uint64{1, 10}, []uint64{10}}
	for i, term := range terms {
		if bm, err = f.Query(term); err != nil {
			t.Fatalf("%+v", err)
		}
		docIDs = bm.Bits()
		fmt.Printf("found term %s in documents: %v\n", term, docIDs)
		if isEqual, err = checkers.DeepEqual(docIDs, expDocIDs[i]); !isEqual {
			t.Fatalf("incorrect result of (*TermDict).GetTermsID, %+v", err)
		}
	}
}
