package indexer

import (
	"log"

	"github.com/google/codesearch/index"
	"github.com/google/codesearch/regexp"
	"github.com/google/codesearch/sparse"
	"github.com/pilosa/pilosa"
)

// TextTriFrame represents a text field of an index which can be searched with trigrams. Refers to pilosa.Frame and pilosa.View.
type TextTriFrame struct {
	TextFrame
}

// NewTextTriFrame returns a new instance of frame, and initializes it.
func NewTextTriFrame(path, index, name string, overwrite bool) (f *TextTriFrame, err error) {
	f = &TextTriFrame{}
	err = f.TextFrame.init(path, index, name, overwrite)
	return
}

// DoIndex parses and index a field.
func (f *TextTriFrame) DoIndex(docID uint64, text string) (err error) {
	grams := ParseTrigrams(text)
	for _, tri := range grams {
		if _, err = f.setBit(uint64(tri), docID); err != nil {
			return
		}
	}
	return
}

//Query query which documents contain the given trigrams.
func (f *TextTriFrame) QueryPat(pat string) (bm *pilosa.Bitmap) {
	q := RegexpQuery(pat)
	bm = f.PostingQuery(q)
	return
}

//Query query which documents contain the given trigrams.
func (f *TextTriFrame) Query(q *index.Query) (bm *pilosa.Bitmap) {
	bm = f.PostingQuery(q)
	return
}

func (f *TextTriFrame) PostingQuery(q *index.Query) *pilosa.Bitmap {
	return f.postingQuery(q, nil)
}

func RegexpQuery(pat string) (q *index.Query) {
	pat = "(?m)" + pat
	re, err := regexp.Compile(pat)
	if err != nil {
		log.Fatal(err)
	}
	q = index.RegexpQuery(re.Syntax)
	//log.Printf("query: %s\n", q)
	return
}

//ParseTrigrams parses text(encoded in UTF-8) for trigrams.
func ParseTrigrams(text string) (trigrams []uint32) {
	grams := sparse.NewSet(1 << 24)
	var (
		c  = byte(0)
		n  = 0
		tv = uint32(0)
	)
	for i := 0; i < len(text); i++ {
		c = text[i]
		tv = (tv << 8) & (1<<24 - 1)
		tv |= uint32(c)
		if n++; n >= 3 {
			if validUTF8((tv>>8)&0xFF, tv&0xFF) {
				grams.Add(tv)
			}
		}
	}
	trigrams = grams.Dense()
	return
}

// validUTF8 reports whether the byte pair can appear in a
// valid sequence of UTF-8-encoded code points.
func validUTF8(c1, c2 uint32) bool {
	switch {
	case c1 < 0x80:
		// 1-byte, must be followed by 1-byte or first of multi-byte
		return c2 < 0x80 || 0xc0 <= c2 && c2 < 0xf8
	case c1 < 0xc0:
		// continuation byte, can be followed by nearly anything
		return c2 < 0xf8
	case c1 < 0xf8:
		// first of multi-byte, must be followed by continuation byte
		return 0x80 <= c2 && c2 < 0xc0
	}
	return false
}

//restrict==nil means all
//ret==nil means all.
func (f *TextTriFrame) postingQuery(q *index.Query, restrict *pilosa.Bitmap) *pilosa.Bitmap {
	var list *pilosa.Bitmap
	switch q.Op {
	case index.QNone:
		return pilosa.NewBitmap()
	case index.QAll:
		return restrict
	case index.QAnd:
		for _, t := range q.Trigram {
			tri := uint32(t[0])<<16 | uint32(t[1])<<8 | uint32(t[2])
			if list == nil {
				list = f.postingList(tri, restrict)
			} else {
				list = list.Intersect(f.postingList(tri, restrict))
			}
			if list.Count() == 0 {
				return list
			}
		}
		for _, sub := range q.Sub {
			if list == nil {
				list = restrict
			}
			list = f.postingQuery(sub, list)
			if list.Count() == 0 {
				return list
			}
		}
	case index.QOr:
		for _, t := range q.Trigram {
			tri := uint32(t[0])<<16 | uint32(t[1])<<8 | uint32(t[2])
			if list == nil {
				list = f.postingList(tri, restrict)
			} else {
				list = list.Union(f.postingList(tri, restrict))
			}
		}
		for _, sub := range q.Sub {
			if list == nil {
				list = pilosa.NewBitmap()
			}
			list = list.Union(f.postingQuery(sub, restrict))
		}
	}
	return list
}

//read tri's bitmap
func (f *TextTriFrame) postingList(tri uint32, restrict *pilosa.Bitmap) (ret *pilosa.Bitmap) {
	ret = f.row(uint64(tri))
	if restrict != nil {
		ret = ret.Intersect(restrict)
	}
	return
}
