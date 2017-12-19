package indexer

import (
	"fmt"
	"testing"

	"github.com/juju/testing/checkers"
	"github.com/pilosa/pilosa"
)

func TestTextTriFrameParseTrigrams(t *testing.T) {
	text := "cindex为若干路径创建索引。索引是trigram倒排表。"
	expect := []uint32{6515054, 6909540, 7234661, 6579576, 6650084, 7922872, 14989498, 12106472, 12249227, 15240101, 9151973, 10872249, 15055282, 12169960, 11724983, 15251375, 12038117, 11527614, 15056516, 12485861, 8709512, 15042715, 8952805, 10216891, 15055802, 12303079, 12249012, 15185058, 11838181, 10675644, 15056021, 12359139, 9823104, 14909570, 8422119, 8578996, 12359142, 9823896, 15112367, 10006388, 11498610, 7631465, 7498087, 6907762, 6779489, 7496045, 6385125, 7202176, 15040658, 8426214, 9627278, 15109778, 9343720, 9627809, 15245736, 10594531, 11068288}
	trigrams := ParseTrigrams(text)
	fmt.Printf("text: %v\n", text)
	fmt.Printf("trigrams: %v\n", trigrams)
	var err error
	var isEqual bool
	if isEqual, err = checkers.DeepEqual(trigrams, expect); !isEqual {
		t.Fatalf("incorrect result of ParseTrigrams, %+v", err)
	}
}

func TestTextTriFrameQuery(t *testing.T) {
	var err error
	var f *TextTriFrame
	var pats []string
	var bm *pilosa.Bitmap
	var isEqual bool
	var bits map[uint64][]uint64

	//TESTCASE: query and insert term to an empty dict
	if f, err = NewTextTriFrame("/tmp/text_tri_frame_test", "i", "f", true); err != nil {
		t.Fatalf("%+v", err)
	}
	defer f.Close()

	docIDs := []uint64{1, 10}
	texts := []string{
		"Go's standard library does not have a function solely intended to check if a file exists or not (like Python's os.path.exists). What is the idiomatic way to do it? 你好，世界",
		"This is a listing of successful results of all the various data storage and processing system benchmarks I've conducted using the dataset produced in the Billion Taxi Rides in Redshift blog post. The dataset itself has 1.1 billion records, 51 columns and takes up about 500 GB of disk space uncompressed.",
	}
	for i := 0; i < len(docIDs); i++ {
		if err = f.DoIndex(docIDs[i], texts[i]); err != nil {
			t.Fatalf("%+v", err)
		}
	}
	if bits, err = f.Bits(); err != nil {
		t.Fatalf("%+v", err)
	}
	fmt.Printf("frame bits: %v\n", bits)

	pats = []string{"The", "(?i)The", "disk", "fu(nc|l)", "你好", "你坏"}
	expDocIDs := [][]uint64{[]uint64{10}, []uint64{1, 10}, []uint64{10}, []uint64{1, 10}, []uint64{1}, []uint64{}}
	for i, pat := range pats {
		bm = f.QueryPat(pat)
		docIDs = bm.Bits()
		if isEqual, err = checkers.DeepEqual(docIDs, expDocIDs[i]); !isEqual {
			t.Fatalf("incorrect result of (*TextTriFrame).Query, %+v", err)
		}
	}
}

func BenchmarkTextTriFrameDoIndex(b *testing.B) {
	var err error
	var f *TextTriFrame
	if f, err = NewTextTriFrame("/tmp/text_tri_frame_test", "i", "f", true); err != nil {
		b.Fatalf("%+v", err)
	}
	defer f.Close()

	b.ResetTimer()
	text := "Go's standard library does not have a function solely intended to check if a file exists or not (like Python's os.path.exists). What is the idiomatic way to do it? cindex为若干路径创建索引。索引是trigram倒排表。trigram是UTF-8文档中的连续3字节(可以是中英文混合)。posting list就是文档ID列表，将它们的delta以变长编码方式存放。整个索引存储在一个文件，在read时mmap到内存。所以索引尺寸受限于RAM。"
	for i := 0; i < b.N; i++ {
		if err = f.DoIndex(uint64(i), text); err != nil {
			b.Fatalf("%+v", err)
		}
	}
}
