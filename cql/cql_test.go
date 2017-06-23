package cql

import (
	"fmt"
	"testing"
)

func TestParseCql(t *testing.T) {
	tcs := []string{
		"IDX.CREATE orders SCHEMA object UINT64 price UINT32 number UINT32 date UINT64 desc STRING",
		"IDX.DESTROY orders",
	}
	for _, tc := range tcs {
		res, err := ParseCql(tc, nil)
		if err != nil {
			t.Fatalf("%+v", err)
		}
		switch res.(type) {
		case CqlCreate:
			create := res.(CqlCreate)
			fmt.Printf("Create index %s schema %v\n", create.Index, create.DocProt)
		case CqlDestroy:
			destroy := res.(CqlDestroy)
			fmt.Printf("Destroy index %s\n", destroy.Index)
		}
	}
}
