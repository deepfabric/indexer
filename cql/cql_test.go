package cql

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseCql(t *testing.T) {
	tcs := []string{
		"IDX.CREATE orders SCHEMA type ENUM",
		"IDX.CREATE orders SCHEMA object UINT64 price UINT32 number UINT32 date UINT64 type ENUM desc STRING",
		"IDX.CREATE orders SCHEMA object UINT64 price UINT32 number UINT32 date UINT64",
		"IDX.CREATE orders SCHEMA object UINT64 price UINT32 number UINT32 date UINT64 type ENUM",
		"IDX.CREATE orders SCHEMA object UINT64 price UINT32 number UINT32 date UINT64 desc STRING",
		"IDX.INSERT orders 615 11 22 33 44 \"description\"",
		"IDX.DEL orders 615 11 22 33 44 \"description\"",
		"IDX.DESTROY orders",
		"IDX.SELECT orders WHERE price>=30 price<40 date<2017 type IN [1,3] desc CONTAINS \"pen\" ORDERBY date",
		"IDX.SELECT orders WHERE price>=30 price<=40 date<2017 type IN [1,3] ORDERBY date LIMIT 30",
		"IDX.SELECT orders WHERE price>=30 price<=40 type IN [1,3]",
	}
	indexDefs := make(map[string]IndexDef)
	for i, tc := range tcs {
		res, err := ParseCql(tc, indexDefs)
		if err != nil {
			t.Fatalf("case %d, error %+v", i, err)
		}
		switch r := res.(type) {
		case *CqlCreate:
			fmt.Printf("Create index %s, schema %v\n", r.Index, r.IndexDef)
			indexDefs[r.Index] = r.IndexDef
		case *CqlDestroy:
			fmt.Printf("Destroy index %s\n", r.Index)
			delete(indexDefs, r.Index)
		case *CqlInsert:
			fmt.Printf("Insert %v\n", r)
		case *CqlDel:
			fmt.Printf("Del %v\n", r)
		case *CqlSelect:
			fmt.Printf("Select %v\n", r)
		default:
			//There shouldn't be any parsing error for above test cases.
			t.Fatalf("case %d, res %+v\n", i, res)
		}
	}
}

func TestParseCqlSelect(t *testing.T) {
	var res interface{}
	var err error
	var q *CqlSelect
	var uintPred UintPred
	var enumPred EnumPred
	var strPred StrPred
	var ok bool
	//TESTCASE: multiple UintPred of the same property into one
	res, err = ParseCql("IDX.SELECT orders WHERE price>=30 price<=40 price<35 price>20", nil)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	q = res.(*CqlSelect)
	uintPred, ok = q.UintPreds["price"]
	if !ok {
		t.Fatalf("UintPred price is gone")
	} else if uintPred.Low != 30 || uintPred.High != 34 {
		t.Fatalf("incorrect folded UintPred price, have (%v, %v), want (%d, %d)", uintPred.Low, uintPred.High, 30, 34)
	}

	//TESTCASE: normal EnumPred
	res, err = ParseCql("IDX.SELECT orders WHERE type IN [1,3]", nil)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	q = res.(*CqlSelect)
	enumPred, ok = q.EnumPreds["type"]
	if !ok {
		t.Fatalf("EnumPred type is gone")
	} else if len(enumPred.InVals) != 2 || enumPred.InVals[0] != 1 || enumPred.InVals[1] != 3 {
		t.Fatalf("incorrect EnumPred type, have %v, want %v", enumPred.InVals, []int{1, 3})
	}

	//TESTCASE: invalid query due to multiple EnumPred of a property
	res, err = ParseCql("IDX.SELECT orders WHERE type IN [1,3] type IN [3,9]", nil)
	if err == nil {
		t.Fatalf("incorrect EnumPred type, have %v, want error", res)
	}

	//TESTCASE: normal StrPred
	res, err = ParseCql("IDX.SELECT orders WHERE desc CONTAINS \"pen\"", nil)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	q = res.(*CqlSelect)
	strPred, ok = q.StrPreds["desc"]
	if !ok {
		t.Fatalf("StrPred desc is gone")
	} else if !strings.EqualFold(strPred.ContWord, "pen") {
		t.Fatalf("incorrect StrPred desc, have %v, want %v", strPred.ContWord, "pen")
	}

	//TESTCASE: invalid query due to multiple StrPred of a property
	res, err = ParseCql("IDX.SELECT orders WHERE desc CONTAINS \"pen\" desc CONTAINS \"pencil\"", nil)
	if err == nil {
		t.Fatalf("incorrect StrPred type, have %v, want error", res)
	}

	//TESTCASE: invalid query due to OBDERBY property doesn't occur in WHERE
	res, err = ParseCql("IDX.SELECT orders WHERE price>=30 price<=40 ORDERBY date", nil)
	if err == nil {
		t.Fatalf("invalid OBDERBY, have %v, want error", res)
	}

	//TESTCASE: invalid query due to OBDERBY property doesn't occur as a UintPred
	res, err = ParseCql("IDX.SELECT orders WHERE price>=30 price<=40 type IN [1,3] ORDERBY type", nil)
	if err == nil {
		t.Fatalf("invalid OBDERBY, have %v, want error", res)
	}
}
