package cql

const (
	TypeUint8 = iota //0
	TypeUint16
	TypeUint32
	TypeUint64
	TypeEnum
	TypeStr
)

type UintProp struct {
	Name   string
	ValLen int //one of 1, 2, 4, 8
	Val    uint64
}

type EnumProp struct {
	Name string
	Val  int
}

type StrProp struct {
	Name string
	Val  string
}

type Document struct {
	UintProps []UintProp
	EnumProps []EnumProp
	StrProps  []StrProp
}

type IndexDef struct {
	DocProt   Document
	PropTypes map[string]int //property name -> type
}

type UintPred struct {
	Name      string
	Low, High uint64
}

type EnumPred struct {
	Name   string
	InVals []int
}

type StrPred struct {
	Name      string
	ContWords []string
}

type CqlCreate struct {
	IndexDef
}

type CqlDestroy struct {
	Index string
}

type CqlInsert struct {
	Index string
	Doc   Document
}

type CqlDel struct {
	CqlInsert
}

type CqlQuery struct {
	Index     string
	UintPreds []UintPred
	EnumPreds []UintPred
	StrPreds  []StrPred
	OrderBy   string
	Limit     int
}

//ParseCql parse CQL. res type is one of CqlCreate/CqlDestroy/CqlInsert/CqlDel/CqlQuery.
func ParseCql(cql string, indexDefs map[string]IndexDef) (res interface{}, err error) {
	err = nil
	return
}
