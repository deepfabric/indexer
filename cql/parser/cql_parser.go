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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 120,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 42, 10, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 6, 3, 48, 10, 3, 13, 3, 14, 3, 49, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 6, 5, 59, 10, 5, 13, 5, 14, 5, 60, 3, 6, 3, 6, 3, 6, 3, 6, 6,
	6, 67, 10, 6, 13, 6, 14, 6, 68, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 5, 7, 79, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 96, 10, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 6, 14, 103, 10, 14, 13, 14, 14, 14,
	104, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 5, 16, 114, 10, 16,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 2, 2, 19, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 2, 4, 4, 2, 12, 14, 29, 30, 3, 2, 24,
	28, 2, 120, 2, 41, 3, 2, 2, 2, 4, 43, 3, 2, 2, 2, 6, 51, 3, 2, 2, 2, 8,
	54, 3, 2, 2, 2, 10, 62, 3, 2, 2, 2, 12, 70, 3, 2, 2, 2, 14, 80, 3, 2, 2,
	2, 16, 82, 3, 2, 2, 2, 18, 85, 3, 2, 2, 2, 20, 95, 3, 2, 2, 2, 22, 97,
	3, 2, 2, 2, 24, 99, 3, 2, 2, 2, 26, 102, 3, 2, 2, 2, 28, 106, 3, 2, 2,
	2, 30, 113, 3, 2, 2, 2, 32, 115, 3, 2, 2, 2, 34, 117, 3, 2, 2, 2, 36, 42,
	5, 4, 3, 2, 37, 42, 5, 6, 4, 2, 38, 42, 5, 8, 5, 2, 39, 42, 5, 10, 6, 2,
	40, 42, 5, 12, 7, 2, 41, 36, 3, 2, 2, 2, 41, 37, 3, 2, 2, 2, 41, 38, 3,
	2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 40, 3, 2, 2, 2, 42, 3, 3, 2, 2, 2, 43,
	44, 7, 3, 2, 2, 44, 45, 5, 14, 8, 2, 45, 47, 7, 4, 2, 2, 46, 48, 5, 16,
	9, 2, 47, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50,
	3, 2, 2, 2, 50, 5, 3, 2, 2, 2, 51, 52, 7, 5, 2, 2, 52, 53, 5, 14, 8, 2,
	53, 7, 3, 2, 2, 2, 54, 55, 7, 6, 2, 2, 55, 56, 5, 14, 8, 2, 56, 58, 5,
	22, 12, 2, 57, 59, 5, 24, 13, 2, 58, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2,
	60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 9, 3, 2, 2, 2, 62, 63, 7, 7,
	2, 2, 63, 64, 5, 14, 8, 2, 64, 66, 5, 22, 12, 2, 65, 67, 5, 24, 13, 2,
	66, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3,
	2, 2, 2, 69, 11, 3, 2, 2, 2, 70, 71, 7, 8, 2, 2, 71, 72, 5, 14, 8, 2, 72,
	73, 7, 9, 2, 2, 73, 74, 5, 26, 14, 2, 74, 75, 7, 10, 2, 2, 75, 78, 5, 18,
	10, 2, 76, 77, 7, 11, 2, 2, 77, 79, 5, 34, 18, 2, 78, 76, 3, 2, 2, 2, 78,
	79, 3, 2, 2, 2, 79, 13, 3, 2, 2, 2, 80, 81, 7, 32, 2, 2, 81, 15, 3, 2,
	2, 2, 82, 83, 5, 18, 10, 2, 83, 84, 5, 20, 11, 2, 84, 17, 3, 2, 2, 2, 85,
	86, 7, 32, 2, 2, 86, 19, 3, 2, 2, 2, 87, 96, 3, 2, 2, 2, 88, 96, 7, 15,
	2, 2, 89, 96, 7, 16, 2, 2, 90, 96, 7, 17, 2, 2, 91, 96, 7, 18, 2, 2, 92,
	96, 7, 19, 2, 2, 93, 96, 7, 20, 2, 2, 94, 96, 7, 21, 2, 2, 95, 87, 3, 2,
	2, 2, 95, 88, 3, 2, 2, 2, 95, 89, 3, 2, 2, 2, 95, 90, 3, 2, 2, 2, 95, 91,
	3, 2, 2, 2, 95, 92, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 94, 3, 2, 2, 2,
	96, 21, 3, 2, 2, 2, 97, 98, 7, 31, 2, 2, 98, 23, 3, 2, 2, 2, 99, 100, 9,
	2, 2, 2, 100, 25, 3, 2, 2, 2, 101, 103, 5, 28, 15, 2, 102, 101, 3, 2, 2,
	2, 103, 104, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105,
	27, 3, 2, 2, 2, 106, 107, 5, 18, 10, 2, 107, 108, 5, 30, 16, 2, 108, 109,
	5, 24, 13, 2, 109, 29, 3, 2, 2, 2, 110, 114, 5, 32, 17, 2, 111, 114, 7,
	22, 2, 2, 112, 114, 7, 23, 2, 2, 113, 110, 3, 2, 2, 2, 113, 111, 3, 2,
	2, 2, 113, 112, 3, 2, 2, 2, 114, 31, 3, 2, 2, 2, 115, 116, 9, 3, 2, 2,
	116, 33, 3, 2, 2, 2, 117, 118, 7, 31, 2, 2, 118, 35, 3, 2, 2, 2, 10, 41,
	49, 60, 68, 78, 95, 104, 113,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'IDX.CREATE'", "'SCHEMA'", "'IDX.DESTROY'", "'IDX.INSERT'", "'IDX.DEL'",
	"'IDX.SELECT'", "'WHERE'", "'ORDERBY'", "'LIMIT'", "'true'", "'false'",
	"'null'", "'UINT8'", "'UINT16'", "'UINT32'", "'UINT64'", "'FLOAT'", "'ENUM'",
	"'STRING'", "'IN'", "'CONTAINS'", "'<'", "'>'", "'='", "'<='", "'>='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "K_UINT8", "K_UINT16",
	"K_UINT32", "K_UINT64", "K_FLOAT", "K_ENUM", "K_STRING", "K_IN", "K_CONTAINS",
	"K_LT", "K_BT", "K_EQ", "K_LE", "K_BE", "STRING", "NUMBER", "INT", "Identifier",
	"WS",
}

var ruleNames = []string{
	"cql", "create", "destroy", "insert", "del", "query", "indexName", "propertyDef",
	"property", "popType", "docId", "value", "predicates", "predicate", "relate",
	"compare", "limit",
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
	CQLParserT__9       = 10
	CQLParserT__10      = 11
	CQLParserT__11      = 12
	CQLParserK_UINT8    = 13
	CQLParserK_UINT16   = 14
	CQLParserK_UINT32   = 15
	CQLParserK_UINT64   = 16
	CQLParserK_FLOAT    = 17
	CQLParserK_ENUM     = 18
	CQLParserK_STRING   = 19
	CQLParserK_IN       = 20
	CQLParserK_CONTAINS = 21
	CQLParserK_LT       = 22
	CQLParserK_BT       = 23
	CQLParserK_EQ       = 24
	CQLParserK_LE       = 25
	CQLParserK_BE       = 26
	CQLParserSTRING     = 27
	CQLParserNUMBER     = 28
	CQLParserINT        = 29
	CQLParserIdentifier = 30
	CQLParserWS         = 31
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
	CQLParserRULE_propertyDef = 7
	CQLParserRULE_property    = 8
	CQLParserRULE_popType     = 9
	CQLParserRULE_docId       = 10
	CQLParserRULE_value       = 11
	CQLParserRULE_predicates  = 12
	CQLParserRULE_predicate   = 13
	CQLParserRULE_relate      = 14
	CQLParserRULE_compare     = 15
	CQLParserRULE_limit       = 16
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

	p.SetState(39)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Create()
		}

	case CQLParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Destroy()
		}

	case CQLParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.Insert()
		}

	case CQLParserT__4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(37)
			p.Del()
		}

	case CQLParserT__5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(38)
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

func (s *CreateContext) AllPropertyDef() []IPropertyDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyDefContext)(nil)).Elem())
	var tst = make([]IPropertyDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyDefContext)
		}
	}

	return tst
}

func (s *CreateContext) PropertyDef(i int) IPropertyDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyDefContext)
}

func (s *CreateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(CQLParserT__0)
	}
	{
		p.SetState(42)
		p.IndexName()
	}
	{
		p.SetState(43)
		p.Match(CQLParserT__1)
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CQLParserIdentifier {
		{
			p.SetState(44)
			p.PropertyDef()
		}

		p.SetState(47)
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
		p.SetState(49)
		p.Match(CQLParserT__2)
	}
	{
		p.SetState(50)
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
		p.SetState(52)
		p.Match(CQLParserT__3)
	}
	{
		p.SetState(53)
		p.IndexName()
	}
	{
		p.SetState(54)
		p.DocId()
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CQLParserT__9)|(1<<CQLParserT__10)|(1<<CQLParserT__11)|(1<<CQLParserSTRING)|(1<<CQLParserNUMBER))) != 0) {
		{
			p.SetState(55)
			p.Value()
		}

		p.SetState(58)
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
		p.SetState(60)
		p.Match(CQLParserT__4)
	}
	{
		p.SetState(61)
		p.IndexName()
	}
	{
		p.SetState(62)
		p.DocId()
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CQLParserT__9)|(1<<CQLParserT__10)|(1<<CQLParserT__11)|(1<<CQLParserSTRING)|(1<<CQLParserNUMBER))) != 0) {
		{
			p.SetState(63)
			p.Value()
		}

		p.SetState(66)
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
		p.SetState(68)
		p.Match(CQLParserT__5)
	}
	{
		p.SetState(69)
		p.IndexName()
	}
	{
		p.SetState(70)
		p.Match(CQLParserT__6)
	}
	{
		p.SetState(71)
		p.Predicates()
	}
	{
		p.SetState(72)
		p.Match(CQLParserT__7)
	}
	{
		p.SetState(73)
		p.Property()
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserT__8 {
		{
			p.SetState(74)
			p.Match(CQLParserT__8)
		}
		{
			p.SetState(75)
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
		p.SetState(78)
		p.Match(CQLParserIdentifier)
	}

	return localctx
}

// IPropertyDefContext is an interface to support dynamic dispatch.
type IPropertyDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyDefContext differentiates from other interfaces.
	IsPropertyDefContext()
}

type PropertyDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyDefContext() *PropertyDefContext {
	var p = new(PropertyDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_propertyDef
	return p
}

func (*PropertyDefContext) IsPropertyDefContext() {}

func NewPropertyDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyDefContext {
	var p = new(PropertyDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_propertyDef

	return p
}

func (s *PropertyDefContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyDefContext) Property() IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *PropertyDefContext) PopType() IPopTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPopTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPopTypeContext)
}

func (s *PropertyDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitPropertyDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) PropertyDef() (localctx IPropertyDefContext) {
	localctx = NewPropertyDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CQLParserRULE_propertyDef)

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
		p.SetState(80)
		p.Property()
	}
	{
		p.SetState(81)
		p.PopType()
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
	p.EnterRule(localctx, 16, CQLParserRULE_property)

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
		p.SetState(83)
		p.Match(CQLParserIdentifier)
	}

	return localctx
}

// IPopTypeContext is an interface to support dynamic dispatch.
type IPopTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPopTypeContext differentiates from other interfaces.
	IsPopTypeContext()
}

type PopTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPopTypeContext() *PopTypeContext {
	var p = new(PopTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLParserRULE_popType
	return p
}

func (*PopTypeContext) IsPopTypeContext() {}

func NewPopTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PopTypeContext {
	var p = new(PopTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_popType

	return p
}

func (s *PopTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PopTypeContext) K_UINT8() antlr.TerminalNode {
	return s.GetToken(CQLParserK_UINT8, 0)
}

func (s *PopTypeContext) K_UINT16() antlr.TerminalNode {
	return s.GetToken(CQLParserK_UINT16, 0)
}

func (s *PopTypeContext) K_UINT32() antlr.TerminalNode {
	return s.GetToken(CQLParserK_UINT32, 0)
}

func (s *PopTypeContext) K_UINT64() antlr.TerminalNode {
	return s.GetToken(CQLParserK_UINT64, 0)
}

func (s *PopTypeContext) K_FLOAT() antlr.TerminalNode {
	return s.GetToken(CQLParserK_FLOAT, 0)
}

func (s *PopTypeContext) K_ENUM() antlr.TerminalNode {
	return s.GetToken(CQLParserK_ENUM, 0)
}

func (s *PopTypeContext) K_STRING() antlr.TerminalNode {
	return s.GetToken(CQLParserK_STRING, 0)
}

func (s *PopTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PopTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PopTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CQLVisitor:
		return t.VisitPopType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CQLParser) PopType() (localctx IPopTypeContext) {
	localctx = NewPopTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CQLParserRULE_popType)

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

	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserEOF, CQLParserIdentifier:
		p.EnterOuterAlt(localctx, 1)

	case CQLParserK_UINT8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.Match(CQLParserK_UINT8)
		}

	case CQLParserK_UINT16:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(87)
			p.Match(CQLParserK_UINT16)
		}

	case CQLParserK_UINT32:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(88)
			p.Match(CQLParserK_UINT32)
		}

	case CQLParserK_UINT64:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(89)
			p.Match(CQLParserK_UINT64)
		}

	case CQLParserK_FLOAT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(90)
			p.Match(CQLParserK_FLOAT)
		}

	case CQLParserK_ENUM:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(91)
			p.Match(CQLParserK_ENUM)
		}

	case CQLParserK_STRING:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(92)
			p.Match(CQLParserK_STRING)
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
	p.EnterRule(localctx, 20, CQLParserRULE_docId)

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

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(CQLParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CQLParserNUMBER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	p.EnterRule(localctx, 22, CQLParserRULE_value)
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
	p.SetState(97)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CQLParserT__9)|(1<<CQLParserT__10)|(1<<CQLParserT__11)|(1<<CQLParserSTRING)|(1<<CQLParserNUMBER))) != 0) {
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
	p.EnterRule(localctx, 24, CQLParserRULE_predicates)
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
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CQLParserIdentifier {
		{
			p.SetState(99)
			p.Predicate()
		}

		p.SetState(102)
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
	p.EnterRule(localctx, 26, CQLParserRULE_predicate)

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
		p.SetState(104)
		p.Property()
	}
	{
		p.SetState(105)
		p.Relate()
	}
	{
		p.SetState(106)
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
	p.EnterRule(localctx, 28, CQLParserRULE_relate)

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

	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserK_LT, CQLParserK_BT, CQLParserK_EQ, CQLParserK_LE, CQLParserK_BE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Compare()
		}

	case CQLParserK_IN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.Match(CQLParserK_IN)
		}

	case CQLParserK_CONTAINS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
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
	p.EnterRule(localctx, 30, CQLParserRULE_compare)
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
	p.SetState(113)
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
	p.EnterRule(localctx, 32, CQLParserRULE_limit)

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
