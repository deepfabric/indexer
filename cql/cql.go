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

type CqlSelect struct {
	Index     string
	UintPreds map[string]UintPred
	EnumPreds map[string]EnumPred
	StrPreds  map[string]StrPred
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
	res       interface{} //record the intermediate and final result of visitor
}

func (v *myCqlVisitor) VisitCql(ctx *parser.CqlContext) (err interface{}) {
	//If there are multiple subrules, then check one by one.
	if create := ctx.Create(); create != nil {
		err = v.VisitCreate(create.(*parser.CreateContext))
	} else if destroy := ctx.Destroy(); destroy != nil {
		err = v.VisitDestroy(destroy.(*parser.DestroyContext))
	} else if ins := ctx.Insert(); ins != nil {
		err = v.VisitInsert(ins.(*parser.InsertContext))
	} else if del := ctx.Del(); del != nil {
		err = v.VisitDel(del.(*parser.DelContext))
	} else if query := ctx.Query(); query != nil {
		err = v.VisitQuery(query.(*parser.QueryContext))
	} else {
		err = errors.Errorf("unsupported subrule of cql")
	}
	//v.res has already be populated
	return
}

func (v *myCqlVisitor) VisitCreate(ctx *parser.CreateContext) (err interface{}) {
	q := &CqlCreate{}
	q.Index = ctx.IndexName().GetText()
	q.PropTypes = make(map[string]int)
	for _, popDef := range ctx.AllUintPropDef() {
		if err = v.VisitUintPropDef(popDef.(*parser.UintPropDefContext)); err != nil {
			return
		}
		q.UintProps = append(q.UintProps, v.res.(UintProp))
	}
	for _, popDef := range ctx.AllEnumPropDef() {
		if err = v.VisitEnumPropDef(popDef.(*parser.EnumPropDefContext)); err != nil {
			return
		}
		q.EnumProps = append(q.EnumProps, v.res.(EnumProp))
	}
	for _, popDef := range ctx.AllStrPropDef() {
		if err = v.VisitStrPropDef(popDef.(*parser.StrPropDefContext)); err != nil {
			return
		}
		q.StrProps = append(q.StrProps, v.res.(StrProp))
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
			err = errors.Errorf("incorrect pop.ValLen %d", pop.ValLen)
			return
		}
	}
	for _, pop := range q.EnumProps {
		q.PropTypes[pop.Name] = TypeEnum
	}
	for _, pop := range q.StrProps {
		q.PropTypes[pop.Name] = TypeStr
	}
	v.res = q
	return
}

func (v *myCqlVisitor) VisitUintPropDef(ctx *parser.UintPropDefContext) (err interface{}) {
	var pop UintProp
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
		err = errors.Errorf("incorrect uintType: %v %v\n", pop.Name, ctx.UintType().GetText())
		return
	}
	v.res = pop
	return
}

func (v *myCqlVisitor) VisitEnumPropDef(ctx *parser.EnumPropDefContext) (err interface{}) {
	var pop EnumProp
	pop.Name = ctx.Property().GetText()
	v.res = pop
	return
}

func (v *myCqlVisitor) VisitStrPropDef(ctx *parser.StrPropDefContext) (err interface{}) {
	var pop StrProp
	pop.Name = ctx.Property().GetText()
	v.res = pop
	return
}

func (v *myCqlVisitor) VisitDestroy(ctx *parser.DestroyContext) (err interface{}) {
	q := &CqlDestroy{}
	q.Index = ctx.IndexName().GetText()
	v.res = q
	return
}

func (v *myCqlVisitor) VisitInsert(ctx *parser.InsertContext) (err interface{}) {
	if err = v.VisitDocument(ctx.Document().(*parser.DocumentContext)); err != nil {
		return
	}
	q := &CqlInsert{} //TODO: better way to copy doc?
	q.DocumentWithIdx = *(v.res.(*DocumentWithIdx))
	v.res = q
	return
}

func (v *myCqlVisitor) VisitDel(ctx *parser.DelContext) (err interface{}) {
	if err = v.VisitDocument(ctx.Document().(*parser.DocumentContext)); err != nil {
		return
	}
	q := &CqlDel{} //TODO: better way to copy doc?
	q.DocumentWithIdx = *(v.res.(*DocumentWithIdx))
	v.res = q
	return
}

func (v *myCqlVisitor) VisitDocument(ctx *parser.DocumentContext) (err interface{}) {
	index := ctx.IndexName().GetText()
	indexDef, ok := v.indexDefs[index]
	if !ok {
		err = errors.Errorf("failed to find the definion of index %s\n", index)
		return
	} else if len(indexDef.PropTypes) != len(ctx.AllValue()) {
		err = errors.Errorf("incorrect number of values, is %d, want %d\n", len(ctx.AllValue()), len(indexDef.PropTypes))
		return
	}
	doc := &DocumentWithIdx{}
	*doc = indexDef.DocumentWithIdx
	doc.Index = ctx.IndexName().GetText()
	var tmpU64 uint64
	var tmpInt int
	tmpU64, err = strconv.ParseUint(ctx.DocId().GetText(), 10, 64)
	if err != nil {
		err = errors.Wrap(err.(error), "")
		return
	}
	doc.DocID = tmpU64

	vals := ctx.AllValue()
	for i := 0; i < len(doc.UintProps); i++ {
		tmpU64, err = strconv.ParseUint(vals[i].GetText(), 10, 64)
		if err != nil {
			err = errors.Wrap(err.(error), "")
			return
		}
		doc.UintProps[i].Val = tmpU64
	}
	for i := 0; i < len(doc.EnumProps); i++ {
		tmpInt, err = strconv.Atoi(vals[i+len(doc.UintProps)].GetText())
		if err != nil {
			err = errors.Wrap(err.(error), "")
			return
		}
		doc.EnumProps[i].Val = tmpInt
	}
	for i := 0; i < len(doc.StrProps); i++ {
		doc.StrProps[i].Val = vals[i+len(doc.UintProps)+len(doc.EnumProps)].GetText()
	}
	v.res = doc
	return
}

func (v *myCqlVisitor) VisitQuery(ctx *parser.QueryContext) (err interface{}) {
	q := &CqlSelect{
		Index:     ctx.IndexName().GetText(),
		UintPreds: make(map[string]UintPred),
		EnumPreds: make(map[string]EnumPred),
		StrPreds:  make(map[string]StrPred),
	}
	if ordCtx := ctx.Order(); ordCtx != nil {
		q.OrderBy = ordCtx.GetText()
	}
	if lmtCtx := ctx.Limit(); lmtCtx != nil {
		q.Limit, err = strconv.Atoi(lmtCtx.GetText())
		if err != nil {
			err = errors.Wrap(err.(error), "")
			return
		}
	}
	for _, predCtx := range ctx.AllUintPred() {
		if err = v.VisitUintPred(predCtx.(*parser.UintPredContext)); err != nil {
			return
		}
		uintPred := *(v.res.(*UintPred))
		uintPred2, ok := q.UintPreds[uintPred.Name]
		if ok {
			//fold multiple UintPred of the same property into one
			uintPred.Low = maxU64(uintPred.Low, uintPred2.Low)
			uintPred.High = minU64(uintPred.High, uintPred2.High)
		}
		q.UintPreds[uintPred.Name] = uintPred
	}
	for _, predCtx := range ctx.AllEnumPred() {
		if err = v.VisitEnumPred(predCtx.(*parser.EnumPredContext)); err != nil {
			return
		}
		enumPred := *(v.res.(*EnumPred))
		if _, ok := q.EnumPreds[enumPred.Name]; ok {
			err = errors.Errorf("invalid query due to multiple EnumPred of property %s", enumPred.Name)
			return
		}
		q.EnumPreds[enumPred.Name] = enumPred
	}
	for _, predCtx := range ctx.AllStrPred() {
		if err = v.VisitStrPred(predCtx.(*parser.StrPredContext)); err != nil {
			return
		}
		strPred := *(v.res.(*StrPred))
		if _, ok := q.StrPreds[strPred.Name]; ok {
			err = errors.Errorf("invalid query due to multiple StrPred of property %s", strPred.Name)
			return
		}
		q.StrPreds[strPred.Name] = strPred
	}
	v.res = q
	return
}

//https://stackoverflow.com/questions/27516387/what-is-the-correct-way-to-find-the-min-between-two-integers-in-go
func minU64(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func maxU64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

//https://stackoverflow.com/questions/44222554/how-to-remove-quotes-from-around-a-string-in-golang
func stripQuote(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

func (v *myCqlVisitor) VisitUintPred(ctx *parser.UintPredContext) (err interface{}) {
	var val uint64
	pred := &UintPred{Low: 0, High: ^uint64(0)}
	pred.Name = ctx.Property().GetText()
	val, err = strconv.ParseUint(ctx.INT().GetText(), 10, 64)
	if err != nil {
		err = errors.Wrap(err.(error), "")
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
		err = errors.Errorf("incorrect compare: %v\n", cmp.GetText())
		return
	}
	v.res = pred
	return
}

func (v *myCqlVisitor) VisitEnumPred(ctx *parser.EnumPredContext) (err interface{}) {
	pred := &EnumPred{}
	pred.Name = ctx.Property().GetText()
	if err = v.VisitIntList(ctx.IntList().(*parser.IntListContext)); err != nil {
		return
	}
	pred.InVals = v.res.([]int)
	v.res = pred
	return
}

func (v *myCqlVisitor) VisitIntList(ctx *parser.IntListContext) (err interface{}) {
	intList := make([]int, 0, len(ctx.AllINT()))
	var val int
	for _, it := range ctx.AllINT() {
		val, err = strconv.Atoi(it.GetText())
		if err != nil {
			err = errors.Wrap(err.(error), "")
			return
		}
		intList = append(intList, val)
	}
	v.res = intList
	return
}

func (v *myCqlVisitor) VisitStrPred(ctx *parser.StrPredContext) (err interface{}) {
	pred := &StrPred{}
	pred.Name = ctx.Property().GetText()
	pred.ContWord = stripQuote(ctx.STRING().GetText())
	v.res = pred
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
	if err1 := tree.Accept(visitor); err1 != nil {
		err = err1.(error) //panic: interface conversion: interface is nil, not error
		return
	}
	res = visitor.res
	return
}
