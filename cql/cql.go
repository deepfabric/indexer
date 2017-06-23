package cql

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/deepfabric/indexer/cql/parser"
	"github.com/pkg/errors"
)

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
	Index     string
	PropTypes map[string]int //property name -> type
	DocProt   Document
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

type myErrListener struct {
	antlr.DefaultErrorListener
	exp antlr.RecognitionException
	msg string
}

func (el *myErrListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	parser := recognizer.(antlr.Parser)
	stack := parser.GetRuleInvocationStack(parser.GetParserRuleContext())
	el.msg = fmt.Sprintf("rule stack: %v, line %d:%d at %v: %s\n", stack, line, column, offendingSymbol, msg)
	el.exp = e
}

type myCqlVisitor struct {
	parser.BaseCQLVisitor
	res interface{} //record the result of visitor
}

func (v *myCqlVisitor) VisitCql(ctx *parser.CqlContext) interface{} {
	//If there are multiple subrules, then check one by one.
	if create := ctx.Create(); create != nil {
		v.res = v.VisitCreate(create.(*parser.CreateContext))
	} else if destroy := ctx.Destroy(); destroy != nil {
		v.res = v.VisitDestroy(destroy.(*parser.DestroyContext))
	}
	return nil
}

func (v *myCqlVisitor) VisitCreate(ctx *parser.CreateContext) interface{} {
	q := new(CqlCreate)
	q.PropTypes = make(map[string]int)
	q.Index = ctx.IndexName().GetText()
	for _, popDef := range ctx.AllUintPropDef() {
		pop := v.VisitUintPropDef(popDef.(*parser.UintPropDefContext))
		q.DocProt.UintProps = append(q.DocProt.UintProps, pop.(UintProp))
	}
	for _, popDef := range ctx.AllEnumPropDef() {
		pop := v.VisitEnumPropDef(popDef.(*parser.EnumPropDefContext))
		q.DocProt.EnumProps = append(q.DocProt.EnumProps, pop.(EnumProp))
	}
	for _, popDef := range ctx.AllStrPropDef() {
		pop := v.VisitStrPropDef(popDef.(*parser.StrPropDefContext))
		q.DocProt.StrProps = append(q.DocProt.StrProps, pop.(StrProp))
	}
	for _, pop := range q.DocProt.UintProps {
		switch pop.ValLen {
		case 1:
			q.PropTypes[pop.Name] = TypeUint8
		case 2:
			q.PropTypes[pop.Name] = TypeUint16
		case 4:
			q.PropTypes[pop.Name] = TypeUint32
		case 8:
			q.PropTypes[pop.Name] = TypeUint64
		default:
			panic(fmt.Sprintf("incorrect pop.ValLen %d", pop.ValLen))
		}
	}
	for _, pop := range q.DocProt.EnumProps {
		q.PropTypes[pop.Name] = TypeEnum
	}
	for _, pop := range q.DocProt.StrProps {
		q.PropTypes[pop.Name] = TypeStr
	}
	return q
}

func (v *myCqlVisitor) VisitUintPropDef(ctx *parser.UintPropDefContext) interface{} {
	pop := UintProp{}
	pop.Name = ctx.Property().GetText()
	uintType := ctx.UintType().(*parser.UintTypeContext)
	if u8 := uintType.K_UINT8(); u8 != nil {
		pop.ValLen = 1
	} else if u16 := uintType.K_UINT16(); u16 != nil {
		pop.ValLen = 2
	} else if u32 := uintType.K_UINT32(); u32 != nil {
		pop.ValLen = 4
	} else if u64 := uintType.K_UINT64(); u64 != nil {
		pop.ValLen = 8
	} else {
		//TODO: how to disable parser recovery?
		//TODO: logging
		fmt.Printf("incorrect uintType: %v %v\n", pop.Name, ctx.UintType().GetText())
	}
	return pop
}

func (v *myCqlVisitor) VisitEnumPropDef(ctx *parser.EnumPropDefContext) interface{} {
	pop := EnumProp{}
	pop.Name = ctx.Property().GetText()
	return pop
}

func (v *myCqlVisitor) VisitStrPropDef(ctx *parser.StrPropDefContext) interface{} {
	pop := StrProp{}
	pop.Name = ctx.Property().GetText()
	return pop
}

func (v *myCqlVisitor) VisitDestroy(ctx *parser.DestroyContext) interface{} {
	var q CqlDestroy
	q.Index = ctx.IndexName().GetText()
	return q
}

//ParseCql parse CQL. res type is one of CqlCreate/CqlDestroy/CqlInsert/CqlDel/CqlQuery.
func ParseCql(cql string, indexDefs map[string]IndexDef) (res interface{}, err error) {
	input := antlr.NewInputStream(cql)
	lexer := parser.NewCQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewCQLParser(stream)
	el := new(myErrListener)
	parser.AddErrorListener(el)

	tree := parser.Cql()
	if el.exp != nil {
		err = errors.Errorf("ParseCql error: %s", el.msg)
		return
	}

	visitor := new(myCqlVisitor)
	tree.Accept(visitor)
	res = visitor.res

	return
}
