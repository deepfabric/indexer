package cql

import (
	"fmt"

	"strconv"

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
	DocID     uint64
	UintProps []UintProp
	EnumProps []EnumProp
	StrProps  []StrProp
}

type DocumentWithIdx struct {
	Document
	Index string
}

type IndexDef struct {
	DocumentWithIdx
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
	Name     string
	ContWord string
}

type CqlCreate struct {
	IndexDef
}

type CqlDestroy struct {
	Index string
}

type CqlInsert struct {
	DocumentWithIdx
}

type CqlDel struct {
	DocumentWithIdx
}

type CqlQuery struct {
	Index     string
	UintPreds []UintPred
	EnumPreds []EnumPred
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
	indexDefs map[string]IndexDef
	res       interface{} //record the result of visitor
}

func (v *myCqlVisitor) VisitCql(ctx *parser.CqlContext) interface{} {
	//If there are multiple subrules, then check one by one.
	if create := ctx.Create(); create != nil {
		v.res = v.VisitCreate(create.(*parser.CreateContext))
	} else if destroy := ctx.Destroy(); destroy != nil {
		v.res = v.VisitDestroy(destroy.(*parser.DestroyContext))
	} else if ins := ctx.Insert(); ins != nil {
		v.res = v.VisitInsert(ins.(*parser.InsertContext))
	} else if del := ctx.Del(); del != nil {
		v.res = v.VisitDel(del.(*parser.DelContext))
	} else if query := ctx.Query(); query != nil {
		v.res = v.VisitQuery(query.(*parser.QueryContext))
	}
	return nil //TODO: better to log something and return error?
}

func (v *myCqlVisitor) VisitCreate(ctx *parser.CreateContext) interface{} {
	q := &CqlCreate{}
	q.Index = ctx.IndexName().GetText()
	q.PropTypes = make(map[string]int)
	for _, popDef := range ctx.AllUintPropDef() {
		pop := v.VisitUintPropDef(popDef.(*parser.UintPropDefContext))
		if pop.(UintProp).ValLen == 0 {
			continue
		}
		q.UintProps = append(q.UintProps, pop.(UintProp))
	}
	for _, popDef := range ctx.AllEnumPropDef() {
		pop := v.VisitEnumPropDef(popDef.(*parser.EnumPropDefContext))
		q.EnumProps = append(q.EnumProps, pop.(EnumProp))
	}
	for _, popDef := range ctx.AllStrPropDef() {
		pop := v.VisitStrPropDef(popDef.(*parser.StrPropDefContext))
		q.StrProps = append(q.StrProps, pop.(StrProp))
	}
	for _, pop := range q.UintProps {
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
	for _, pop := range q.EnumProps {
		q.PropTypes[pop.Name] = TypeEnum
	}
	for _, pop := range q.StrProps {
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
	q := &CqlDestroy{}
	q.Index = ctx.IndexName().GetText()
	return q
}

func (v *myCqlVisitor) VisitInsert(ctx *parser.InsertContext) interface{} {
	itf := v.VisitDocument(ctx.Document().(*parser.DocumentContext))
	if itf == nil {
		return nil
	}
	doc := itf.(*DocumentWithIdx)
	q := &CqlInsert{} //TODO: better way to copy doc?
	q.DocumentWithIdx = *doc
	return q
}

func (v *myCqlVisitor) VisitDel(ctx *parser.DelContext) interface{} {
	itf := v.VisitDocument(ctx.Document().(*parser.DocumentContext))
	if itf == nil {
		return nil
	}
	doc := itf.(*DocumentWithIdx)
	q := &CqlDel{} //TODO: better way to copy doc?
	q.DocumentWithIdx = *doc
	return q
}

func (v *myCqlVisitor) VisitDocument(ctx *parser.DocumentContext) interface{} {
	index := ctx.IndexName().GetText()
	indexDef, ok := v.indexDefs[index]
	if !ok {
		fmt.Printf("failed to find the definion of index %s\n", index)
		return nil //TODO: it's better to log something and return error?
	} else if len(indexDef.PropTypes) != len(ctx.AllValue()) {
		fmt.Printf("incorrect number of values, is %d, want %d\n", len(ctx.AllValue()), len(indexDef.PropTypes))
		return nil
	}
	doc := &DocumentWithIdx{}
	*doc = indexDef.DocumentWithIdx
	doc.Index = ctx.IndexName().GetText()
	u64, err := strconv.ParseUint(ctx.DocId().GetText(), 10, 64)
	if err != nil {
		return nil
	}
	doc.DocID = u64

	vals := ctx.AllValue()
	for i := 0; i < len(doc.UintProps); i++ {
		u64, err := strconv.ParseUint(vals[i].GetText(), 10, 64)
		if err != nil {
			return nil
		}
		doc.UintProps[i].Val = u64
	}
	for i := 0; i < len(doc.EnumProps); i++ {
		tmp, err := strconv.Atoi(vals[i+len(doc.UintProps)].GetText())
		if err != nil {
			return nil
		}
		doc.EnumProps[i].Val = tmp
	}
	for i := 0; i < len(doc.StrProps); i++ {
		doc.StrProps[i].Val = vals[i+len(doc.UintProps)+len(doc.EnumProps)].GetText()
	}
	return doc
}

func (v *myCqlVisitor) VisitQuery(ctx *parser.QueryContext) interface{} {
	q := &CqlQuery{}
	var err error
	q.Index = ctx.IndexName().GetText()
	if ordCtx := ctx.Order(); ordCtx != nil {
		q.OrderBy = ordCtx.GetText()
	}
	if lmtCtx := ctx.Limit(); lmtCtx != nil {
		q.Limit, err = strconv.Atoi(lmtCtx.GetText())
		if err != nil {
			return nil
		}
	}
	for _, predCtx := range ctx.AllUintPred() {
		if pred, err := v.VisitUintPred(predCtx.(*parser.UintPredContext)); err == nil {
			q.UintPreds = append(q.UintPreds, *pred)
		}
	}
	for _, predCtx := range ctx.AllEnumPred() {
		if pred, err := v.VisitEnumPred(predCtx.(*parser.EnumPredContext)); err == nil {
			q.EnumPreds = append(q.EnumPreds, *pred)
		}
	}
	for _, predCtx := range ctx.AllStrPred() {
		if pred, err := v.VisitStrPred(predCtx.(*parser.StrPredContext)); err == nil {
			q.StrPreds = append(q.StrPreds, *pred)
		}
	}
	return q
}

func (v *myCqlVisitor) VisitUintPred(ctx *parser.UintPredContext) (pred *UintPred, err error) {
	var val uint64
	pred = &UintPred{Low: 0, High: ^uint64(0)}
	pred.Name = ctx.Property().GetText()
	val, err = strconv.ParseUint(ctx.INT().GetText(), 10, 64)
	if err != nil {
		return
	}
	cmp := ctx.Compare().(*parser.CompareContext)
	if lt := cmp.K_LT(); lt != nil {
		pred.High = val - 1
	} else if bt := cmp.K_BT(); bt != nil {
		pred.Low = val + 1
	} else if eq := cmp.K_EQ(); eq != nil {
		pred.Low = val
		pred.High = val
	} else if le := cmp.K_LE(); le != nil {
		pred.High = val
	} else if be := cmp.K_BE(); be != nil {
		pred.Low = val
	} else {
		//TODO: how to disable parser recovery?
		//TODO: logging
		fmt.Printf("incorrect compare: %v\n", cmp.GetText())
	}
	return
}

func (v *myCqlVisitor) VisitEnumPred(ctx *parser.EnumPredContext) (pred *EnumPred, err error) {
	pred = &EnumPred{}
	pred.Name = ctx.Property().GetText()
	pred.InVals, err = v.VisitIntList(ctx.IntList().(*parser.IntListContext))
	return
}

func (v *myCqlVisitor) VisitIntList(ctx *parser.IntListContext) (intList []int, err error) {
	intList = make([]int, 0, len(ctx.AllINT()))
	var val int
	for _, it := range ctx.AllINT() {
		val, err = strconv.Atoi(it.GetText())
		if err != nil {
			err = errors.Errorf("failed to convert following text to integer: %s\n", it.GetText())
			return
		}
		intList = append(intList, val)
	}
	return
}

func (v *myCqlVisitor) VisitStrPred(ctx *parser.StrPredContext) (pred *StrPred, err error) {
	pred = &StrPred{}
	pred.Name = ctx.Property().GetText()
	pred.ContWord = ctx.STRING().GetText()
	return
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

	visitor := &myCqlVisitor{indexDefs: indexDefs}
	tree.Accept(visitor)
	res = visitor.res

	return
}
