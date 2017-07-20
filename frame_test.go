package indexer

import (
	"fmt"
	"testing"
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
