package indexer

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	"os"
	"path/filepath"

	"github.com/deepfabric/bkdtree"
	"github.com/deepfabric/indexer/cql"
)

//Conf originate from config file. It's comman to all indices.
type Conf struct {
	//generic items
	Cap int //upper limit of number of documents

	//BKD specific items
	T0mCap      int
	LeafCap     int
	IntraCap    int
	CptInterval time.Duration
}

//Indexer shall be singleton
type Indexer struct {
	MainDir  string                          //the main directory where stores all indices
	Conf     Conf                            //need to persist
	DocProts map[string]*cql.DocumentWithIdx //need to persist
	indices  map[string]*Index
}

//NewIndexer creates and initializes Indexer singleton.
func NewIndexer(dir string) (ir *Indexer, err error) {
	if err = os.MkdirAll(dir, 0700); err != nil {
		err = errors.Wrap(err, "")
		return
	}
	ir = &Indexer{
		MainDir: dir,
	}
	err = ir.Open()
	return
}

//Open opens all indices. Assumes ir.MainDir is already populated.
func (ir *Indexer) Open() (err error) {
	if ir.indices != nil {
		panic("indexer already open")
	}
	if err = ir.readMeta(); err != nil {
		return
	}
	ir.indices = make(map[string]*Index)
	var ind *Index
	for name, docProt := range ir.DocProts {
		if ind, err = ir.openIndex(name, docProt); err != nil {
			return
		}
		ir.indices[name] = ind
	}
	return
}

//writeMeta persists Conf and DocProts to files.
func (ir *Indexer) writeMeta() (err error) {
	var fp string
	fp = filepath.Join(ir.MainDir, "indexer.json")
	if err = bkdtree.FileMarshal(fp, ir.Conf); err != nil {
		return
	}
	for name, docProt := range ir.DocProts {
		if err = indexWriteConf(ir.MainDir, docProt); err != nil {
			return
		}
	}
	return
}

//readMeta parse Conf and DocProts from files.
func (ir *Indexer) readMeta() (err error) {
	fp := filepath.Join(ir.MainDir, "indexer.json")
	if err = bkdtree.FileUnmarshal(fp, &ir.Conf); err != nil {
		return
	}

	var matches [][]string
	patt := `^index_(?P<name>[^.]+)\.json$`
	if matches, err = bkdtree.FilepathGlob(ir.MainDir, patt); err != nil {
		return
	}
	for _, match := range matches {
		var doc cql.Document
		fp = filepath.Join(ir.MainDir, match[0])
		if err = bkdtree.FileUnmarshal(fp, &doc); err != nil {
			return
		}
		ir.DocProts[match[1]] = &doc
	}
	return
}

//openIndex opens the given index from existing index
func (ir *Indexer) openIndex(name string, docProt *cql.Document) (ind *Index, err error) {
	ind = &Index{
		DocProt: docProt,
	}
	ind.bkds = make(map[string]*bkdtree.BkdTree)
	var bkd *bkdtree.BkdTree
	for _, uintProp := range docProt.UintProps {
		bkd = &bkdtree.BkdTree{}
		bkdCap := ir.Conf.Cap
		dir := ir.MainDir
		prefix := uintProp.Name
		cptInterval := ir.Conf.CptInterval
		if err = bkd.Open(dir, prefix, bkdCap, cptInterval); err != nil {
			return
		}
		//TODO: verify T0mCap, LeafCap, IntraCap, NumDims, BytesPerDim?
		ind.bkds[uintProp.Name] = bkd
	}
	return
}

//closeIndices close all indices
func (ir *Indexer) closeIndices() (err error) {
	for name, ind := range ir.indices {
		if err = ind.Close(); err != nil {
			return
		}
		delete(ir.indices, name)
	}
	return
}

// Close close index
func (ind *Index) Close() (err error) {
	for prop, bkd := range ind.bkds {
		if err = bkd.Close(); err != nil {
			return
		}
		delete(ind.bkds, prop)
	}
	return
}

func removeIndexDef(dir, name string) (err error) {
	fp := filepath.Join(dir, fmt.Sprintf("%s.json", name))
	err = os.Remove(fp)
	return
}
