package indexer

import (
	"testing"

	"os"

	"github.com/juju/testing/checkers"
)

func TestGetTermsID(t *testing.T) {
	var err error
	var isEqual bool

	//TESTCASE: query and insert term to an empty dict
	err = os.Remove("/tmp/terms")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	td := &TermDict{
		Dir: "/tmp",
	}
	terms := []string{
		"sunday",
		"mon",
		"tue",
		"wen",
		"thurs",
		"friday",
		"satur",
	}
	expIds := []int64{0, 1, 2, 3, 4, 5, 6}

	ids, err := td.GetTermsID(terms)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if isEqual, err = checkers.DeepEqual(ids, expIds); !isEqual {
		t.Fatalf("incorrect result of (*TermDict).GetTermsID, %+v", err)
	}

	//TESTCASE: query and insert term to an existing dict
	td2 := &TermDict{
		Dir: "/tmp",
	}
	terms = []string{
		"friday",
		"wikepedia",
		"thurs",
	}
	expIds = []int64{5, 7, 4}

	ids, err = td2.GetTermsID(terms)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if isEqual, err = checkers.DeepEqual(ids, expIds); !isEqual {
		t.Fatalf("incorrect result of (*TermDict).GetTermsID, %+v", err)
	}
}
