package indexer

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/pilosa/pilosa"
)

// TextFrame represents a string field of an index. Refers to pilosa.Frame and pilosa.View.
type TextTokFrame struct {
	TextFrame
	td *TermDict
}

// NewTextTokFrame returns a new instance of frame, and initializes it.
func NewTextTokFrame(path, index, name string, overwrite bool) (f *TextTokFrame, err error) {
	f = &TextTokFrame{}
	if err = f.init(path, index, name, overwrite); err != nil {
		return
	}
	if f.td, err = NewTermDict(path, overwrite); err != nil {
		return
	}
	return
}

//Open opens an existing frame
func (f *TextTokFrame) Open() (err error) {
	if err = f.TextFrame.Open(); err != nil {
		return
	}
	err = f.td.Open()
	return
}

// Close closes all fragments without removing files on disk.
// It's allowed to invoke Close multiple times.
func (f *TextTokFrame) Close() (err error) {
	if err = f.TextFrame.Close(); err != nil {
		return
	}
	err = f.td.Close()
	return
}

// Destroy closes all fragments, removes all files on disk.
// It's allowed to invoke Close before or after Destroy.
func (f *TextTokFrame) Destroy() (err error) {
	if err = f.TextFrame.Destroy(); err != nil {
		return
	}
	err = f.td.Destroy()
	return
}

// DoIndex parses and index a field.
func (f *TextTokFrame) DoIndex(docID uint64, text string) (err error) {
	//https://stackoverflow.com/questions/13737745/split-a-string-on-whitespace-in-go
	/*terms := strings.Fields(text)
	for i, term := range terms {
		terms[i] = strings.ToLower(term)
	}*/
	terms := ParseWords(text)
	ids, err := f.td.CreateTermsIfNotExist(terms)
	if err != nil {
		return
	}
	for _, termID := range ids {
		if _, err = f.setBit(termID, docID); err != nil {
			return
		}
	}
	return
}

//Query query which documents contain the given term.
func (f *TextTokFrame) Query(text string) (bm *pilosa.Bitmap) {
	words := ParseWords(text)
	var bm2 *pilosa.Bitmap
	for _, word := range words {
		termID, found := f.td.GetTermID(word)
		if !found {
			bm = pilosa.NewBitmap()
			return
		}
		bm2 = f.row(termID)
		if bm != nil {
			bm = bm.Intersect(bm2)
		} else {
			bm = bm2
		}
	}
	return
}

var asciiSpace = [128]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

//ParseWords parses text(encoded in UTF-8) for words.
//A word is a non-ascii-space lowered ASCII character sequence, or a non-ASCII non-unicode-space non-chinese-punctuate character.
//Note: words are not de-duplicated.
func ParseWords(text string) (words []string) {
	lenText := len(text)
	words = make([]string, 0, lenText/3)
	i := 0
	for i < lenText {
		j := i
		var c byte
		for j < lenText {
			if c = text[j]; c < 0x80 && asciiSpace[c] == 0 && unicode.IsPrint(rune(c)) && !unicode.IsPunct(rune(c)) {
				j++
			} else {
				break
			}
		}
		if i < j {
			// text[i:j] is a printable non-space ASCII character sequence.
			words = append(words, strings.ToLower(text[i:j]))
			i = j
		} else if c < 0x80 {
			// i==j, text[i] is an ascii space, non-printable or punctuation character.
			i++
		} else {
			// i==j, text[i] is the begin of an non-ascii character
			r, w := utf8.DecodeRuneInString(text[i:])
			if unicode.IsPrint(rune(c)) && !unicode.IsSpace(r) && !unicode.IsPunct(r) {
				words = append(words, text[i:i+w])
			}
			i += w
		}
	}
	return
}
