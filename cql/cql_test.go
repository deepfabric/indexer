package cql

import (
	"fmt"
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
	}
	indexDefs := make(map[string]IndexDef)
	for _, tc := range tcs {
		res, err := ParseCql(tc, indexDefs)
		if err != nil {
			t.Fatalf("%+v", err)
		}
		switch res.(type) {
		case *CqlCreate:
			create := res.(*CqlCreate)
			fmt.Printf("Create index %s, schema %v\n", create.Index, create.IndexDef)
			indexDefs[create.Index] = create.IndexDef
		case *CqlDestroy:
			destroy := res.(*CqlDestroy)
			fmt.Printf("Destroy index %s\n", destroy.Index)
			delete(indexDefs, destroy.Index)
		case *CqlInsert:
			cqlInsert := res.(*CqlInsert)
			fmt.Printf("Insert %v\n", cqlInsert)
		case *CqlDel:
			cqlDel := res.(*CqlDel)
			fmt.Printf("Del %v\n", cqlDel)
		default:
			//There shouldn't be any parsing error for above test cases.
			t.Fatalf("res %+v\n", res)
		}
	}
}
