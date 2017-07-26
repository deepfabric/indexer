package indexer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/deepfabric/bkdtree"
	"github.com/deepfabric/indexer/cql"
	"github.com/deepfabric/pilosa"
	"github.com/pkg/errors"
)

//Conf originate from config file. It's comman to all indices.
type Conf struct {
	//BKD specific items
	T0mCap   int
	LeafCap  int
	IntraCap int
}

//Indexer shall be singleton
type Indexer struct {
	MainDir  string                          //the main directory where stores all indices
	Conf     Conf                            //indexer conf
	docProts map[string]*cql.DocumentWithIdx //index meta, need to persist
	indices  map[string]*Index               //index data, need to persist
}

//NewIndexer creates an Indexer.
func NewIndexer(mainDir string, conf *Conf, overwirte bool) (ir *Indexer, err error) {
	ir = &Indexer{
		MainDir: mainDir,
		Conf:    *conf,
	}
	if err = os.MkdirAll(mainDir, 0700); err != nil {
		err = errors.Wrap(err, "")
		return
	}
	if overwirte {
		err = ir.removeIndices()
	} else {
		err = ir.Open()
	}
	return
}

//Destroy close and remove index files
func (ir *Indexer) Destroy() (err error) {
	if err = ir.Close(); err != nil {
		return
	}
	if err = ir.removeIndices(); err != nil {
		return
	}
	return
}

//Open opens all indices. Assumes ir.MainDir is already populated.
func (ir *Indexer) Open() (err error) {
	if ir.indices != nil || ir.docProts != nil {
		panic("indexer already open")
	}
	ir.docProts = make(map[string]*cql.DocumentWithIdx)
	ir.indices = make(map[string]*Index)
	if err = ir.readMeta(); err != nil {
		return
	}
	var ind *Index
	for name, docProt := range ir.docProts {
		if ind, err = NewIndexExt(ir.MainDir, docProt.Index); err != nil {
			return
		}
		ir.indices[name] = ind
	}
	return
}

// Close close indexer
func (ir *Indexer) Close() (err error) {
	for _, ind := range ir.indices {
		if err = ind.Close(); err != nil {
			return
		}
	}
	ir.indices = nil
	ir.docProts = nil
	return
}

// CreateIndex creates index
func (ir *Indexer) CreateIndex(docProt *cql.DocumentWithIdx) (err error) {
	if ir.docProts == nil {
		ir.docProts = make(map[string]*cql.DocumentWithIdx)
		ir.indices = make(map[string]*Index)
	}
	if _, found := ir.docProts[docProt.Index]; found {
		panic("CreateIndex conflict with existing index")
	}
	if err = indexWriteConf(ir.MainDir, docProt); err != nil {
		return
	}
	var ind *Index
	if ind, err = NewIndex(docProt, ir.MainDir, ir.Conf.T0mCap, ir.Conf.LeafCap, ir.Conf.IntraCap); err != nil {
		return
	}
	ir.indices[docProt.Index] = ind
	ir.docProts[docProt.Index] = docProt
	return
}

//DestroyIndex destroy given index
func (ir *Indexer) DestroyIndex(name string) (err error) {
	delete(ir.indices, name)
	delete(ir.docProts, name)
	err = ir.removeIndex(name)
	return
}

//Insert executes CqlInsert
func (ir *Indexer) Insert(doc *cql.DocumentWithIdx) (err error) {
	var ind *Index
	var found bool
	if ind, found = ir.indices[doc.Index]; !found {
		err = errors.Errorf("failed to insert %v to non-existing index %v", doc, doc.Index)
		return
	}
	err = ind.Insert(doc)
	return
}

//Del executes CqlDel.
func (ir *Indexer) Del(doc *cql.DocumentWithIdx) (found bool, err error) {
	var ind *Index
	var fnd bool
	if ind, fnd = ir.indices[doc.Index]; !fnd {
		err = errors.Errorf("failed to delete %v from non-existing index %v", doc, doc.Index)
		return
	}
	found, err = ind.Del(doc)
	return
}

//Select executes CqlSelect.
func (ir *Indexer) Select(q *cql.CqlSelect) (rb *pilosa.Bitmap, err error) {
	var ind *Index
	var found bool
	if ind, found = ir.indices[q.Index]; !found {
		err = errors.Errorf("failed to select %v from non-existing index %v", q, q.Index)
		return
	}
	rb, err = ind.Select(q)
	return
}

//writeMeta persists Conf and DocProts to files.
func (ir *Indexer) writeMeta() (err error) {
	for _, docProt := range ir.docProts {
		if err = indexWriteConf(ir.MainDir, docProt); err != nil {
			return
		}
	}
	return
}

//readMeta parses Conf and DocProts from files.
func (ir *Indexer) readMeta() (err error) {
	var matches [][]string
	patt := `^index_(?P<name>[^.]+)\.json$`
	if matches, err = bkdtree.FilepathGlob(ir.MainDir, patt); err != nil {
		return
	}
	for _, match := range matches {
		var doc cql.DocumentWithIdx
		if err = indexReadConf(ir.MainDir, match[1], &doc); err != nil {
			return
		}
		ir.docProts[match[1]] = &doc
	}
	return
}

func (ir *Indexer) removeIndices() (err error) {
	var matches [][]string
	patt := `^index_(?P<name>[^.]+)\.json$`
	if matches, err = bkdtree.FilepathGlob(ir.MainDir, patt); err != nil {
		return
	}
	for _, match := range matches {
		if err = ir.removeIndex(match[1]); err != nil {
			return
		}
	}
	return
}

func (ir *Indexer) removeIndex(name string) (err error) {
	var fp string
	fp = filepath.Join(ir.MainDir, fmt.Sprintf("index_%s.json", name))
	if err = os.Remove(fp); err != nil {
		err = errors.Wrap(err, "")
		return
	}
	fp = filepath.Join(ir.MainDir, name)
	if err = os.RemoveAll(fp); err != nil {
		err = errors.Wrap(err, "")
	}
	return
}
