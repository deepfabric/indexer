package indexer

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

//TermDict stores terms in a map. Note that the term dict is insertion-only.
type TermDict struct {
	Dir    string
	f      *os.File
	terms  map[string]int64
	rwlock sync.RWMutex //concurrent access of TermDict
}

//GetTermID get id of the given term, will insert the term implicitly if it is not in the dict.
func (td *TermDict) GetTermID(term string) (id int64, err error) {
	if td.terms == nil {
		//According to https://blog.golang.org/defer-panic-and-recover,
		//"A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns.""
		td.rwlock.Lock()
		fp := filepath.Join(td.Dir, "terms")
		if td.f, err = os.OpenFile(fp, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600); err != nil {
			td.rwlock.Unlock()
			err = errors.Wrap(err, "")
			return
		}
		td.terms = make(map[string]int64)
		reader := bufio.NewReader(td.f)
		var line string
		var num int64
		for {
			line, err = reader.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				err = errors.Wrap(err, "")
				td.rwlock.Unlock()
				return
			}
			tmpTerm := strings.TrimSpace(line)
			td.terms[tmpTerm] = num
			num++
		}
		td.rwlock.Unlock()
	}
	var found bool
	td.rwlock.RLock()
	if id, found = td.terms[term]; found {
		td.rwlock.RUnlock()
		return id, nil
	}
	td.rwlock.RUnlock()
	td.rwlock.Lock()
	defer td.rwlock.Unlock()
	if id, found = td.terms[term]; found {
		return id, nil
	}
	id = int64(len(td.terms))
	td.terms[term] = id
	line := term + "\n"
	if _, err = td.f.WriteString(line); err != nil {
		err = errors.Wrap(err, "")
		return
	}
	return
}

//GetTermsID get id for an array of terms
func (td *TermDict) GetTermsID(terms []string) (ids []int64, err error) {
	ids = make([]int64, len(terms))
	for i, term := range terms {
		if ids[i], err = td.GetTermID(term); err != nil {
			return
		}
	}
	return
}
