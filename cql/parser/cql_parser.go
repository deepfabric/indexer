// Generated from /home/zhichyu/src/github.com/deepfabric/indexer/cql/parser/CQL.g4 by ANTLR 4.7.

package parser // CQL

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 140,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 46,
	10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 52, 10, 3, 12, 3, 14, 3, 55, 11, 3,
	3, 3, 7, 3, 58, 10, 3, 12, 3, 14, 3, 61, 11, 3, 3, 3, 7, 3, 64, 10, 3,
	12, 3, 14, 3, 67, 11, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5,
	76, 10, 5, 13, 5, 14, 5, 77, 3, 6, 3, 6, 3, 6, 3, 6, 6, 6, 84, 10, 6, 13,
	6, 14, 6, 85, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 96,
	10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 116, 10,
	13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 6, 16, 123, 10, 16, 13, 16, 14,
	16, 124, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 5, 18, 134, 10,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 2, 2, 21, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 2, 4, 3, 2, 25, 26, 3,
	2, 20, 24, 2, 137, 2, 45, 3, 2, 2, 2, 4, 47, 3, 2, 2, 2, 6, 68, 3, 2, 2,
	2, 8, 71, 3, 2, 2, 2, 10, 79, 3, 2, 2, 2, 12, 87, 3, 2, 2, 2, 14, 97, 3,
	2, 2, 2, 16, 99, 3, 2, 2, 2, 18, 102, 3, 2, 2, 2, 20, 105, 3, 2, 2, 2,
	22, 108, 3, 2, 2, 2, 24, 115, 3, 2, 2, 2, 26, 117, 3, 2, 2, 2, 28, 119,
	3, 2, 2, 2, 30, 122, 3, 2, 2, 2, 32, 126, 3, 2, 2, 2, 34, 133, 3, 2, 2,
	2, 36, 135, 3, 2, 2, 2, 38, 137, 3, 2, 2, 2, 40, 46, 5, 4, 3, 2, 41, 46,
	5, 6, 4, 2, 42, 46, 5, 8, 5, 2, 43, 46, 5, 10, 6, 2, 44, 46, 5, 12, 7,
	2, 45, 40, 3, 2, 2, 2, 45, 41, 3, 2, 2, 2, 45, 42, 3, 2, 2, 2, 45, 43,
	3, 2, 2, 2, 45, 44, 3, 2, 2, 2, 46, 3, 3, 2, 2, 2, 47, 48, 7, 3, 2, 2,
	48, 49, 5, 14, 8, 2, 49, 53, 7, 4, 2, 2, 50, 52, 5, 16, 9, 2, 51, 50, 3,
	2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54,
	59, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 56, 58, 5, 18, 10, 2, 57, 56, 3, 2,
	2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 65,
	3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 64, 5, 20, 11, 2, 63, 62, 3, 2, 2,
	2, 64, 67, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 5, 3,
	2, 2, 2, 67, 65, 3, 2, 2, 2, 68, 69, 7, 5, 2, 2, 69, 70, 5, 14, 8, 2, 70,
	7, 3, 2, 2, 2, 71, 72, 7, 6, 2, 2, 72, 73, 5, 14, 8, 2, 73, 75, 5, 26,
	14, 2, 74, 76, 5, 28, 15, 2, 75, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77,
	75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 9, 3, 2, 2, 2, 79, 80, 7, 7, 2,
	2, 80, 81, 5, 14, 8, 2, 81, 83, 5, 26, 14, 2, 82, 84, 5, 28, 15, 2, 83,
	82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2,
	2, 86, 11, 3, 2, 2, 2, 87, 88, 7, 8, 2, 2, 88, 89, 5, 14, 8, 2, 89, 90,
	7, 9, 2, 2, 90, 91, 5, 30, 16, 2, 91, 92, 7, 10, 2, 2, 92, 95, 5, 22, 12,
	2, 93, 94, 7, 11, 2, 2, 94, 96, 5, 38, 20, 2, 95, 93, 3, 2, 2, 2, 95, 96,
	3, 2, 2, 2, 96, 13, 3, 2, 2, 2, 97, 98, 7, 28, 2, 2, 98, 15, 3, 2, 2, 2,
	99, 100, 5, 22, 12, 2, 100, 101, 5, 24, 13, 2, 101, 17, 3, 2, 2, 2, 102,
	103, 5, 22, 12, 2, 103, 104, 7, 16, 2, 2, 104, 19, 3, 2, 2, 2, 105, 106,
	5, 22, 12, 2, 106, 107, 7, 17, 2, 2, 107, 21, 3, 2, 2, 2, 108, 109, 7,
	28, 2, 2, 109, 23, 3, 2, 2, 2, 110, 116, 3, 2, 2, 2, 111, 116, 7, 12, 2,
	2, 112, 116, 7, 13, 2, 2, 113, 116, 7, 14, 2, 2, 114, 116, 7, 15, 2, 2,
	115, 110, 3, 2, 2, 2, 115, 111, 3, 2, 2, 2, 115, 112, 3, 2, 2, 2, 115,
	113, 3, 2, 2, 2, 115, 114, 3, 2, 2, 2, 116, 25, 3, 2, 2, 2, 117, 118, 7,
	27, 2, 2, 118, 27, 3, 2, 2, 2, 119, 120, 9, 2, 2, 2, 120, 29, 3, 2, 2,
	2, 121, 123, 5, 32, 17, 2, 122, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2,
	124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 31, 3, 2, 2, 2, 126, 127,
	5, 22, 12, 2, 127, 128, 5, 34, 18, 2, 128, 129, 5, 28, 15, 2, 129, 33,
	3, 2, 2, 2, 130, 134, 5, 36, 19, 2, 131, 134, 7, 18, 2, 2, 132, 134, 7,
	19, 2, 2, 133, 130, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 132, 3, 2, 2,
	2, 134, 35, 3, 2, 2, 2, 135, 136, 9, 3, 2, 2, 136, 37, 3, 2, 2, 2, 137,
	138, 7, 27, 2, 2, 138, 39, 3, 2, 2, 2, 12, 45, 53, 59, 65, 77, 85, 95,
	115, 124, 133,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'IDX.CREATE'", "'SCHEMA'", "'IDX.DESTROY'", "'IDX.INSERT'", "'IDX.DEL'",
	"'IDX.SELECT'", "'WHERE'", "'ORDERBY'", "'LIMIT'", "'UINT8'", "'UINT16'",
	"'UINT32'", "'UINT64'", "'ENUM'", "'STRING'", "'IN'", "'CONTAINS'", "'<'",
	"'>'", "'='", "'<='", "'>='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "K_UINT8", "K_UINT16", "K_UINT32",
	"K_UINT64", "K_ENUM", "K_STRING", "K_IN", "K_CONTAINS", "K_LT", "K_BT",
	"K_EQ", "K_LE", "K_BE", "STRING", "NUMBER", "INT", "Identifier", "WS",
}

var ruleNames = []string{
	"cql", "create", "destroy", "insert", "del", "query", "indexName", "uintPropDef",
	"enumPropDef", "strPropDef", "property", "uintType", "docId", "value",
	"predicates", "predicate", "relate", "compare", "limit",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CQLParser struct {
	*antlr.BaseParser
}

func NewCQLParser(input antlr.TokenStream) *CQLParser {
	this := new(CQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CQL.g4"

	return this
}

// CQLParser tokens.
const (
	CQLParserEOF        = antlr.TokenEOF
	CQLParserT__0       = 1
	CQLParserT__1       = 2
	CQLParserT__2       = 3
	CQLParserT__3       = 4
	CQLParserT__4       = 5
	CQLParserT__5       = 6
	CQLParserT__6       = 7
	CQLParserT__7       = 8
	CQLParserT__8       = 9
	CQLParserK_UINT8    = 10
	CQLParserK_UINT16   = 11
	CQLParserK_UINT32   = 12
	CQLParserK_UINT64   = 13
	CQLParserK_ENUM     = 14
	CQLParserK_STRING   = 15
	CQLParserK_IN       = 16
	CQLParserK_CONTAINS = 17
	CQLParserK_LT       = 18
	CQLParserK_BT       = 19
	CQLParserK_EQ       = 20
	CQLParserK_LE       = 21
	CQLParserK_BE       = 22
	CQLParserSTRING     = 23
	CQLParserNUMBER     = 24
	CQLParserINT        = 25
	CQLParserIdentifier = 26
	CQLParserWS         = 27
)

// CQLParser rules.
const (
	CQLParserRULE_cql         = 0
	CQLParserRULE_create      = 1
	CQLParserRULE_destroy     = 2
	CQLParserRULE_insert      = 3
	CQLParserRULE_del         = 4
	CQLParserRULE_query       = 5
	CQLParserRULE_indexName   = 6
	CQLParserRULE_uintPropDef = 7
	CQLParserRULE_enumPropDef = 8
	CQLParserRULE_strPropDef  = 9
	CQLParserRULE_property    = 10
	CQLParserRULE_uintType    = 11
	CQLParserRULE_docId       = 12
	CQLParserRULE_value       = 13
	CQLParserRULE_predicates  = 14
	CQLParserRULE_predicate   = 15
	CQLParserRULE_relate      = 16
	CQLParserRULE_compare     = 17
	CQLParserRULE_limit       = 18
)

// ICqlContext is an interface to support dynamic dispatch.
type ICqlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCqlContext differentiates from other interfaces.
	IsCqlContext()
}

type CqlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCqlContext() *CqlContext {
	var p = new(CqlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_cql
	return p
}

func (*CqlContext) IsCqlContext() {}

func NewCqlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlContext {
	var p = new(CqlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_cql

	return p
}

func (s *CqlContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlContext) Create() ICreateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreateContext)
}

func (s *CqlContext) Destroy() IDestroyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDestroyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDestroyContext)
}

func (s *CqlContext) Insert() IInsertContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsertContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsertContext)
}

func (s *CqlContext) Del() IDelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelContext)
}

func (s *CqlContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *CqlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterCql(s)
	}
}

func (s *CqlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitCql(s)
	}
}

func (s *CqlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitCql(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) Cql() (localctx ICqlContext) {
	localctx = NewCqlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CQLParserRULE_cql)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(43)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Create()
		}

	case CQLParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Destroy()
		}

	case CQLParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
			p.Insert()
		}

	case CQLParserT__4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(41)
			p.Del()
		}

	case CQLParserT__5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(42)
			p.Query()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICreateContext is an interface to support dynamic dispatch.
type ICreateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateContext differentiates from other interfaces.
	IsCreateContext()
}

type CreateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateContext() *CreateContext {
	var p = new(CreateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_create
	return p
}

func (*CreateContext) IsCreateContext() {}

func NewCreateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateContext {
	var p = new(CreateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_create

	return p
}

func (s *CreateContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateContext) IndexName() IIndexNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *CreateContext) AllUintPropDef() []IUintPropDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUintPropDefContext)(nil)).Elem())
	var tst = make([]IUintPropDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUintPropDefContext)
		}
	}

	return tst
}

func (s *CreateContext) UintPropDef(i int) IUintPropDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUintPropDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUintPropDefContext)
}

func (s *CreateContext) AllEnumPropDef() []IEnumPropDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumPropDefContext)(nil)).Elem())
	var tst = make([]IEnumPropDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumPropDefContext)
		}
	}

	return tst
}

func (s *CreateContext) EnumPropDef(i int) IEnumPropDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumPropDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumPropDefContext)
}

func (s *CreateContext) AllStrPropDef() []IStrPropDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStrPropDefContext)(nil)).Elem())
	var tst = make([]IStrPropDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStrPropDefContext)
		}
	}

	return tst
}

func (s *CreateContext) StrPropDef(i int) IStrPropDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrPropDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStrPropDefContext)
}

func (s *CreateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterCreate(s)
	}
}

func (s *CreateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitCreate(s)
	}
}

func (s *CreateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitCreate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) Create() (localctx ICreateContext) {
	localctx = NewCreateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CQLParserRULE_create)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(CQLParserT__0)
	}
	{
		p.SetState(46)
		p.IndexName()
	}
	{
		p.SetState(47)
		p.Match(CQLParserT__1)
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(48)
				p.UintPropDef()
			}

		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(54)
				p.EnumPropDef()
			}

		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserIdentifier {
		{
			p.SetState(60)
			p.StrPropDef()
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDestroyContext is an interface to support dynamic dispatch.
type IDestroyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDestroyContext differentiates from other interfaces.
	IsDestroyContext()
}

type DestroyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDestroyContext() *DestroyContext {
	var p = new(DestroyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_destroy
	return p
}

func (*DestroyContext) IsDestroyContext() {}

func NewDestroyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestroyContext {
	var p = new(DestroyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_destroy

	return p
}

func (s *DestroyContext) GetParser() antlr.Parser { return s.parser }

func (s *DestroyContext) IndexName() IIndexNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *DestroyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestroyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DestroyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterDestroy(s)
	}
}

func (s *DestroyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitDestroy(s)
	}
}

func (s *DestroyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitDestroy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) Destroy() (localctx IDestroyContext) {
	localctx = NewDestroyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CQLParserRULE_destroy)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(CQLParserT__2)
	}
	{
		p.SetState(67)
		p.IndexName()
	}

	return localctx
}

// IInsertContext is an interface to support dynamic dispatch.
type IInsertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsertContext differentiates from other interfaces.
	IsInsertContext()
}

type InsertContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertContext() *InsertContext {
	var p = new(InsertContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_insert
	return p
}

func (*InsertContext) IsInsertContext() {}

func NewInsertContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertContext {
	var p = new(InsertContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_insert

	return p
}

func (s *InsertContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertContext) IndexName() IIndexNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *InsertContext) DocId() IDocIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDocIdContext)
}

func (s *InsertContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *InsertContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *InsertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterInsert(s)
	}
}

func (s *InsertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitInsert(s)
	}
}

func (s *InsertContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitInsert(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) Insert() (localctx IInsertContext) {
	localctx = NewInsertContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CQLParserRULE_insert)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(CQLParserT__3)
	}
	{
		p.SetState(70)
		p.IndexName()
	}
	{
		p.SetState(71)
		p.DocId()
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CQLParserSTRING || _la == CQLParserNUMBER {
		{
			p.SetState(72)
			p.Value()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDelContext is an interface to support dynamic dispatch.
type IDelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelContext differentiates from other interfaces.
	IsDelContext()
}

type DelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelContext() *DelContext {
	var p = new(DelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_del
	return p
}

func (*DelContext) IsDelContext() {}

func NewDelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelContext {
	var p = new(DelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_del

	return p
}

func (s *DelContext) GetParser() antlr.Parser { return s.parser }

func (s *DelContext) IndexName() IIndexNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *DelContext) DocId() IDocIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDocIdContext)
}

func (s *DelContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *DelContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterDel(s)
	}
}

func (s *DelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitDel(s)
	}
}

func (s *DelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitDel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) Del() (localctx IDelContext) {
	localctx = NewDelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CQLParserRULE_del)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(CQLParserT__4)
	}
	{
		p.SetState(78)
		p.IndexName()
	}
	{
		p.SetState(79)
		p.DocId()
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CQLParserSTRING || _la == CQLParserNUMBER {
		{
			p.SetState(80)
			p.Value()
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) IndexName() IIndexNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *QueryContext) Predicates() IPredicatesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicatesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicatesContext)
}

func (s *QueryContext) Property() IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *QueryContext) Limit() ILimitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CQLParserRULE_query)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(CQLParserT__5)
	}
	{
		p.SetState(86)
		p.IndexName()
	}
	{
		p.SetState(87)
		p.Match(CQLParserT__6)
	}
	{
		p.SetState(88)
		p.Predicates()
	}
	{
		p.SetState(89)
		p.Match(CQLParserT__7)
	}
	{
		p.SetState(90)
		p.Property()
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserT__8 {
		{
			p.SetState(91)
			p.Match(CQLParserT__8)
		}
		{
			p.SetState(92)
			p.Limit()
		}

	}

	return localctx
}

// IIndexNameContext is an interface to support dynamic dispatch.
type IIndexNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexNameContext differentiates from other interfaces.
	IsIndexNameContext()
}

type IndexNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexNameContext() *IndexNameContext {
	var p = new(IndexNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_indexName
	return p
}

func (*IndexNameContext) IsIndexNameContext() {}

func NewIndexNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexNameContext {
	var p = new(IndexNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_indexName

	return p
}

func (s *IndexNameContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CQLParserIdentifier, 0)
}

func (s *IndexNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterIndexName(s)
	}
}

func (s *IndexNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitIndexName(s)
	}
}

func (s *IndexNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitIndexName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) IndexName() (localctx IIndexNameContext) {
	localctx = NewIndexNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CQLParserRULE_indexName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(CQLParserIdentifier)
	}

	return localctx
}

// IUintPropDefContext is an interface to support dynamic dispatch.
type IUintPropDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUintPropDefContext differentiates from other interfaces.
	IsUintPropDefContext()
}

type UintPropDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUintPropDefContext() *UintPropDefContext {
	var p = new(UintPropDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_uintPropDef
	return p
}

func (*UintPropDefContext) IsUintPropDefContext() {}

func NewUintPropDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UintPropDefContext {
	var p = new(UintPropDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_uintPropDef

	return p
}

func (s *UintPropDefContext) GetParser() antlr.Parser { return s.parser }

func (s *UintPropDefContext) Property() IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *UintPropDefContext) UintType() IUintTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUintTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUintTypeContext)
}

func (s *UintPropDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UintPropDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UintPropDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterUintPropDef(s)
	}
}

func (s *UintPropDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitUintPropDef(s)
	}
}

func (s *UintPropDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitUintPropDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) UintPropDef() (localctx IUintPropDefContext) {
	localctx = NewUintPropDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CQLParserRULE_uintPropDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Property()
	}
	{
		p.SetState(98)
		p.UintType()
	}

	return localctx
}

// IEnumPropDefContext is an interface to support dynamic dispatch.
type IEnumPropDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumPropDefContext differentiates from other interfaces.
	IsEnumPropDefContext()
}

type EnumPropDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumPropDefContext() *EnumPropDefContext {
	var p = new(EnumPropDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_enumPropDef
	return p
}

func (*EnumPropDefContext) IsEnumPropDefContext() {}

func NewEnumPropDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumPropDefContext {
	var p = new(EnumPropDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_enumPropDef

	return p
}

func (s *EnumPropDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumPropDefContext) Property() IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *EnumPropDefContext) K_ENUM() antlr.TerminalNode {
	return s.GetToken(CQLParserK_ENUM, 0)
}

func (s *EnumPropDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumPropDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumPropDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterEnumPropDef(s)
	}
}

func (s *EnumPropDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitEnumPropDef(s)
	}
}

func (s *EnumPropDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitEnumPropDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) EnumPropDef() (localctx IEnumPropDefContext) {
	localctx = NewEnumPropDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CQLParserRULE_enumPropDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Property()
	}
	{
		p.SetState(101)
		p.Match(CQLParserK_ENUM)
	}

	return localctx
}

// IStrPropDefContext is an interface to support dynamic dispatch.
type IStrPropDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrPropDefContext differentiates from other interfaces.
	IsStrPropDefContext()
}

type StrPropDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrPropDefContext() *StrPropDefContext {
	var p = new(StrPropDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_strPropDef
	return p
}

func (*StrPropDefContext) IsStrPropDefContext() {}

func NewStrPropDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrPropDefContext {
	var p = new(StrPropDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_strPropDef

	return p
}

func (s *StrPropDefContext) GetParser() antlr.Parser { return s.parser }

func (s *StrPropDefContext) Property() IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *StrPropDefContext) K_STRING() antlr.TerminalNode {
	return s.GetToken(CQLParserK_STRING, 0)
}

func (s *StrPropDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrPropDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrPropDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterStrPropDef(s)
	}
}

func (s *StrPropDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitStrPropDef(s)
	}
}

func (s *StrPropDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitStrPropDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) StrPropDef() (localctx IStrPropDefContext) {
	localctx = NewStrPropDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CQLParserRULE_strPropDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Property()
	}
	{
		p.SetState(104)
		p.Match(CQLParserK_STRING)
	}

	return localctx
}

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_property
	return p
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CQLParserIdentifier, 0)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (s *PropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CQLParserRULE_property)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(CQLParserIdentifier)
	}

	return localctx
}

// IUintTypeContext is an interface to support dynamic dispatch.
type IUintTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUintTypeContext differentiates from other interfaces.
	IsUintTypeContext()
}

type UintTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUintTypeContext() *UintTypeContext {
	var p = new(UintTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_uintType
	return p
}

func (*UintTypeContext) IsUintTypeContext() {}

func NewUintTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UintTypeContext {
	var p = new(UintTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_uintType

	return p
}

func (s *UintTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *UintTypeContext) K_UINT8() antlr.TerminalNode {
	return s.GetToken(CQLParserK_UINT8, 0)
}

func (s *UintTypeContext) K_UINT16() antlr.TerminalNode {
	return s.GetToken(CQLParserK_UINT16, 0)
}

func (s *UintTypeContext) K_UINT32() antlr.TerminalNode {
	return s.GetToken(CQLParserK_UINT32, 0)
}

func (s *UintTypeContext) K_UINT64() antlr.TerminalNode {
	return s.GetToken(CQLParserK_UINT64, 0)
}

func (s *UintTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UintTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UintTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterUintType(s)
	}
}

func (s *UintTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitUintType(s)
	}
}

func (s *UintTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitUintType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) UintType() (localctx IUintTypeContext) {
	localctx = NewUintTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CQLParserRULE_uintType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserEOF, CQLParserIdentifier:
		p.EnterOuterAlt(localctx, 1)

	case CQLParserK_UINT8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.Match(CQLParserK_UINT8)
		}

	case CQLParserK_UINT16:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.Match(CQLParserK_UINT16)
		}

	case CQLParserK_UINT32:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(111)
			p.Match(CQLParserK_UINT32)
		}

	case CQLParserK_UINT64:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(112)
			p.Match(CQLParserK_UINT64)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDocIdContext is an interface to support dynamic dispatch.
type IDocIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocIdContext differentiates from other interfaces.
	IsDocIdContext()
}

type DocIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocIdContext() *DocIdContext {
	var p = new(DocIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_docId
	return p
}

func (*DocIdContext) IsDocIdContext() {}

func NewDocIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocIdContext {
	var p = new(DocIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_docId

	return p
}

func (s *DocIdContext) GetParser() antlr.Parser { return s.parser }

func (s *DocIdContext) INT() antlr.TerminalNode {
	return s.GetToken(CQLParserINT, 0)
}

func (s *DocIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterDocId(s)
	}
}

func (s *DocIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitDocId(s)
	}
}

func (s *DocIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitDocId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) DocId() (localctx IDocIdContext) {
	localctx = NewDocIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CQLParserRULE_docId)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(CQLParserINT)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CQLParserNUMBER, 0)
}

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(CQLParserSTRING, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CQLParserRULE_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(117)
	_la = p.GetTokenStream().LA(1)

	if !(_la == CQLParserSTRING || _la == CQLParserNUMBER) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IPredicatesContext is an interface to support dynamic dispatch.
type IPredicatesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicatesContext differentiates from other interfaces.
	IsPredicatesContext()
}

type PredicatesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicatesContext() *PredicatesContext {
	var p = new(PredicatesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_predicates
	return p
}

func (*PredicatesContext) IsPredicatesContext() {}

func NewPredicatesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicatesContext {
	var p = new(PredicatesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_predicates

	return p
}

func (s *PredicatesContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicatesContext) AllPredicate() []IPredicateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPredicateContext)(nil)).Elem())
	var tst = make([]IPredicateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPredicateContext)
		}
	}

	return tst
}

func (s *PredicatesContext) Predicate(i int) IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *PredicatesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicatesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicatesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterPredicates(s)
	}
}

func (s *PredicatesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitPredicates(s)
	}
}

func (s *PredicatesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitPredicates(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) Predicates() (localctx IPredicatesContext) {
	localctx = NewPredicatesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CQLParserRULE_predicates)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CQLParserIdentifier {
		{
			p.SetState(119)
			p.Predicate()
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) Property() IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *PredicateContext) Relate() IRelateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelateContext)
}

func (s *PredicateContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (s *PredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CQLParserRULE_predicate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Property()
	}
	{
		p.SetState(125)
		p.Relate()
	}
	{
		p.SetState(126)
		p.Value()
	}

	return localctx
}

// IRelateContext is an interface to support dynamic dispatch.
type IRelateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelateContext differentiates from other interfaces.
	IsRelateContext()
}

type RelateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelateContext() *RelateContext {
	var p = new(RelateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_relate
	return p
}

func (*RelateContext) IsRelateContext() {}

func NewRelateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelateContext {
	var p = new(RelateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_relate

	return p
}

func (s *RelateContext) GetParser() antlr.Parser { return s.parser }

func (s *RelateContext) Compare() ICompareContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareContext)
}

func (s *RelateContext) K_IN() antlr.TerminalNode {
	return s.GetToken(CQLParserK_IN, 0)
}

func (s *RelateContext) K_CONTAINS() antlr.TerminalNode {
	return s.GetToken(CQLParserK_CONTAINS, 0)
}

func (s *RelateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterRelate(s)
	}
}

func (s *RelateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitRelate(s)
	}
}

func (s *RelateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitRelate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) Relate() (localctx IRelateContext) {
	localctx = NewRelateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CQLParserRULE_relate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserK_LT, CQLParserK_BT, CQLParserK_EQ, CQLParserK_LE, CQLParserK_BE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Compare()
		}

	case CQLParserK_IN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Match(CQLParserK_IN)
		}

	case CQLParserK_CONTAINS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Match(CQLParserK_CONTAINS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICompareContext is an interface to support dynamic dispatch.
type ICompareContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompareContext differentiates from other interfaces.
	IsCompareContext()
}

type CompareContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareContext() *CompareContext {
	var p = new(CompareContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_compare
	return p
}

func (*CompareContext) IsCompareContext() {}

func NewCompareContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareContext {
	var p = new(CompareContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_compare

	return p
}

func (s *CompareContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareContext) K_LT() antlr.TerminalNode {
	return s.GetToken(CQLParserK_LT, 0)
}

func (s *CompareContext) K_BT() antlr.TerminalNode {
	return s.GetToken(CQLParserK_BT, 0)
}

func (s *CompareContext) K_EQ() antlr.TerminalNode {
	return s.GetToken(CQLParserK_EQ, 0)
}

func (s *CompareContext) K_LE() antlr.TerminalNode {
	return s.GetToken(CQLParserK_LE, 0)
}

func (s *CompareContext) K_BE() antlr.TerminalNode {
	return s.GetToken(CQLParserK_BE, 0)
}

func (s *CompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterCompare(s)
	}
}

func (s *CompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitCompare(s)
	}
}

func (s *CompareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitCompare(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) Compare() (localctx ICompareContext) {
	localctx = NewCompareContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CQLParserRULE_compare)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(133)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CQLParserK_LT)|(1<<CQLParserK_BT)|(1<<CQLParserK_EQ)|(1<<CQLParserK_LE)|(1<<CQLParserK_BE))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ILimitContext is an interface to support dynamic dispatch.
type ILimitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitContext differentiates from other interfaces.
	IsLimitContext()
}

type LimitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitContext() *LimitContext {
	var p = new(LimitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_limit
	return p
}

func (*LimitContext) IsLimitContext() {}

func NewLimitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitContext {
	var p = new(LimitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_limit

	return p
}

func (s *LimitContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitContext) INT() antlr.TerminalNode {
	return s.GetToken(CQLParserINT, 0)
}

func (s *LimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterLimit(s)
	}
}

func (s *LimitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitLimit(s)
	}
}

func (s *LimitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitLimit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) Limit() (localctx ILimitContext) {
	localctx = NewLimitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CQLParserRULE_limit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(CQLParserINT)
	}

	return localctx
}
