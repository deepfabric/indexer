package indexer

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/deepfabric/pilosa"
	"github.com/juju/testing/checkers"
)

func TestFrameParseAndIndex(t *testing.T) {
	var err error
	var found bool
	var f *Frame
	var terms []string

	//TESTCASE: query and insert term to an empty dict
	if f, err = NewFrame("/tmp/frame_test", "i", "f", true); err != nil {
		t.Fatalf("%+v", err)
	}
	defer f.Close()

	text := "Go's standard library does not have a function solely intended to check if a file exists or not (like Python's os.path.exists). What is the idiomatic way to do it? 小明硕士毕业于中国科学院计算所，后在日本京都大学深造。"
	if err = f.ParseAndIndex(3, text); err != nil {
		t.Fatalf("%+v", err)
	}
	fmt.Printf("termdict size: %d\n", f.td.Count())

	terms = []string{"go", "it", "科学院", "中国科学院"}
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

func TestFrameQuery(t *testing.T) {
	var err error
	var f *Frame
	var terms []string
	var bm *pilosa.Bitmap
	var isEqual bool
	var bits map[uint64][]uint64

	//TESTCASE: query and insert term to an empty dict
	if f, err = NewFrame("/tmp/frame_test", "i", "f", true); err != nil {
		t.Fatalf("%+v", err)
	}
	defer f.Close()

	docIDs := []uint64{1, 10}
	texts := []string{
		"Go's standard library does not have a function solely intended to check if a file exists or not (like Python's os.path.exists). What is the idiomatic way to do it? 你好",
		"This is a listing of successful results of all the various data storage and processing system benchmarks I've conducted using the dataset produced in the Billion Taxi Rides in Redshift blog post. The dataset itself has 1.1 billion records, 51 columns and takes up about 500 GB of 中文世界 disk space uncompressed.",
	}
	for i := 0; i < len(docIDs); i++ {
		if err = f.ParseAndIndex(docIDs[i], texts[i]); err != nil {
			t.Fatalf("%+v", err)
		}
	}
	fmt.Printf("termdict size after indexing: %d\n", f.td.Count())
	if bits, err = f.Bits(); err != nil {
		t.Fatalf("%+v", err)
	}
	fmt.Printf("frame bits: %v\n", bits)

	terms = []string{"The", "disk", "你好", "中文", "世界", "你"}
	expDocIDs := [][]uint64{[]uint64{1, 10}, []uint64{10}, []uint64{1}, []uint64{10}, []uint64{10}, []uint64{}}
	for i, term := range terms {
		bm = f.Query(term)
		docIDs = bm.Bits()
		fmt.Printf("found term %s in documents: %v\n", term, docIDs)
		if isEqual, err = checkers.DeepEqual(docIDs, expDocIDs[i]); !isEqual {
			t.Fatalf("incorrect result of (*Frame).Query, %+v", err)
		}
		for _, docID := range expDocIDs[i] {
			if !bm.Contains(docID) {
				t.Fatalf("incorrect result of (*Frame).Query")
			}
		}
	}
}

func TestFrameDestroy(t *testing.T) {
	var err error
	var f *Frame

	if f, err = NewFrame("/tmp/frame_test", "i", "f", true); err != nil {
		t.Fatalf("%+v", err)
	}
	defer f.Close()

	text := "Go's standard library does not have a function solely intended to check if a file exists or not (like Python's os.path.exists). What is the idiomatic way to do it?"
	if err = f.ParseAndIndex(3, text); err != nil {
		t.Fatalf("%+v", err)
	}
	fmt.Printf("termdict size: %d\n", f.td.Count())

	if err = f.Destroy(); err != nil {
		t.Fatalf("%+v", err)
	}

	fps := []string{filepath.Join(f.path, "terms"), filepath.Join(f.path, "fragments")}
	for _, fp := range fps {
		if _, err := os.Stat(fp); err == nil || !os.IsNotExist(err) {
			t.Fatalf("path %s exists, want removed", fp)
		}
	}
	if 0 != f.td.Count() {
		t.Fatalf("f.td.Count() is %d, want 0", f.td.Count())
	}
}

func BenchmarkFrameParseAndIndex(b *testing.B) {
	var err error
	var f *Frame
	if f, err = NewFrame("/tmp/frame_test", "i", "f", true); err != nil {
		b.Fatalf("%+v", err)
	}
	defer f.Close()

	b.ResetTimer()
	text := "Go's standard library does not have a function solely intended to check if a file exists or not (like Python's os.path.exists). What is the idiomatic way to do it?"
	for i := 0; i < b.N; i++ {
		if err = f.ParseAndIndex(uint64(i), text); err != nil {
			b.Fatalf("%+v", err)
		}
	}
}
