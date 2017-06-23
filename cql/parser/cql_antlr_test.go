package parser // CQL
import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type VerboseErrorListener struct {
	antlr.DefaultErrorListener
	exp antlr.RecognitionException
	msg string
}

func (el *VerboseErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	parser := recognizer.(antlr.Parser)
	stack := parser.GetRuleInvocationStack(parser.GetParserRuleContext())
	el.msg = fmt.Sprintf("rule stack: %v, line %d:%d at %v: %s\n", stack, line, column, offendingSymbol, msg)
	el.exp = e
}

//revised from antlr.BailErrorStrategy at github.com/antlr/antlr4/runtime/Go/antlr/error_strategy.go
//original antlr.BailErrorStrategy BUG: panic: interface conversion: interface is nil, not antlr.ParserRuleContext
type BailErrorStrategy struct {
	antlr.DefaultErrorStrategy
}

func NewBailErrorStrategy() *BailErrorStrategy {
	b := new(BailErrorStrategy)
	return b
}

// Instead of recovering from exception {@code e}, re-panic it wrapped
// in a {@link ParseCancellationException} so it is not caught by the
// rule func catches. Use {@link Exception//getCause()} to get the
// original {@link RecognitionException}.
//
func (b *BailErrorStrategy) Recover(recognizer antlr.Parser, e antlr.RecognitionException) {
	context := recognizer.GetParserRuleContext()
	for context != nil {
		context.SetException(e)
		if parent := context.GetParent(); parent != nil {
			context = parent.(antlr.ParserRuleContext)
		} else {
			break
		}
	}
	fmt.Printf("Recover %+v\n", e)
	panic(antlr.NewParseCancellationException()) // TODO we don't emit e properly
}

// Make sure we don't attempt to recover inline if the parser
// successfully recovers, it won't panic an exception.
//
func (b *BailErrorStrategy) RecoverInline(recognizer antlr.Parser) antlr.Token {
	b.Recover(recognizer, antlr.NewInputMisMatchException(recognizer))
	return nil
}

// Make sure we don't attempt to recover from problems in subrules.//
func (b *BailErrorStrategy) Sync(recognizer antlr.Parser) {
	// pass
}

type ParserCase struct {
	Input       string
	ExpectError bool
}

func TestCqlParserError(t *testing.T) {
	fmt.Println("================TestCqlParserError================")
	tcs := []ParserCase{
		{"IDX.CREATE orders SCHEMA object UINT64 price FLOAT number UINT32 date UINT64", false},
		{"IDX orders SCHEMA object UINT64 price FLOAT number UINT32 date UINT64", true},
	}
	for i, tc := range tcs {
		input := antlr.NewInputStream(tc.Input)
		lexer := NewCQLLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		parser := NewCQLParser(stream)
		el := new(VerboseErrorListener)
		parser.AddErrorListener(el)
		parser.SetErrorHandler(NewBailErrorStrategy())
		{
			//TODO: how to disable parser recovery?
			defer func() {
				if e := recover(); e != nil {
					if re, ok := e.(*antlr.ParseCancellationException); ok {
						fmt.Printf("catched ParseCancellationException %+v\n", re)
					} else {
						panic(e)
					}
				}
			}()
			_ = parser.Cql()
		}

		//tree.(*CqlContext).(*antlr.BaseParserRuleContext).exc

		if el.exp != nil {
			fmt.Printf("parser raised exception %s\n", el.msg)
		}
		if (tc.ExpectError && el.exp == nil) || (!tc.ExpectError && el.exp != nil) {
			t.Fatalf("case %d failed. has %v, want %v", i, !tc.ExpectError, tc.ExpectError)
		}
	}
}

type CqlTestListener struct {
	BaseCQLListener
}

func NewCqlTestListener() *CqlTestListener {
	return new(CqlTestListener)
}

/*func (l *CqlTestListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}*/

func (l *CqlTestListener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("VisitTerminal: %v, tokenType: %v\n", node.GetText(), node.GetSymbol().GetTokenType())
}

func (l *CqlTestListener) EnterCql(ctx *CqlContext) {
	fmt.Printf("EnterCql: %v\n", ctx.GetText())
}

func (l *CqlTestListener) ExitCreate(ctx *CreateContext) {
	fmt.Printf("ExitCreate: %v\n", ctx.GetText())
	fmt.Printf("create indexName: %v\n", ctx.IndexName().GetText())
}

//POC of listener
func TestCqlListener(t *testing.T) {
	fmt.Println("================TestCqlListener================")
	input := antlr.NewInputStream("IDX.CREATE orders SCHEMA object UINT64 price FLOAT number UINT32 date UINT64")
	lexer := NewCQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewCQLParser(stream)
	//parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	//parser.BuildParseTrees = true
	tree := parser.Cql()
	listener := NewCqlTestListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}

type CqlTestVisitor struct {
	BaseCQLVisitor
	res interface{} //record the result of visitor
}

func (v *CqlTestVisitor) VisitCql(ctx *CqlContext) interface{} {
	fmt.Printf("VisitCql %v...\n", ctx)
	//If there are multiple subrules, then check one by one.
	if create := ctx.Create(); create != nil {
		v.res = v.VisitCreate(create.(*CreateContext))
	} else if destroy := ctx.Destroy(); destroy != nil {
		v.res = v.VisitDestroy(destroy.(*DestroyContext))
	}
	return nil
}

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

func (v *CqlTestVisitor) VisitCreate(ctx *CreateContext) interface{} {
	fmt.Println("VisitCreate...")
	indexName := ctx.IndexName().GetText()
	fmt.Printf("indexName: %s\n", indexName)
	var doc Document
	for _, popDef := range ctx.AllUintPropDef() {
		pop := v.VisitUintPropDef(popDef.(*UintPropDefContext))
		doc.UintProps = append(doc.UintProps, pop.(UintProp))
	}
	for _, popDef := range ctx.AllEnumPropDef() {
		pop := v.VisitEnumPropDef(popDef.(*EnumPropDefContext))
		doc.EnumProps = append(doc.EnumProps, pop.(EnumProp))
	}
	for _, popDef := range ctx.AllStrPropDef() {
		pop := v.VisitStrPropDef(popDef.(*StrPropDefContext))
		doc.StrProps = append(doc.StrProps, pop.(StrProp))
	}
	return fmt.Sprintf("Create index %s schema %v\n", indexName, doc)
}

func (v *CqlTestVisitor) VisitUintPropDef(ctx *UintPropDefContext) interface{} {
	fmt.Println("VisitUintPropDef...")
	pop := UintProp{}
	pop.Name = ctx.Property().GetText()
	uintType := ctx.UintType().(*UintTypeContext)
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
		//panic(fmt.Sprintf("incorrect uintType: %v", ctx.UintType().GetText()))
		fmt.Printf("incorrect uintType: %v\n", ctx.UintType().GetText())
	}
	return pop
}

func (v *CqlTestVisitor) VisitEnumPropDef(ctx *EnumPropDefContext) interface{} {
	fmt.Println("VisitEnumPropDef...")
	pop := EnumProp{}
	pop.Name = ctx.Property().GetText()
	return pop
}

func (v *CqlTestVisitor) VisitStrPropDef(ctx *StrPropDefContext) interface{} {
	fmt.Println("VisitStrPropDef...")
	pop := StrProp{}
	pop.Name = ctx.Property().GetText()
	return pop
}

func (v *CqlTestVisitor) VisitDestroy(ctx *DestroyContext) interface{} {
	fmt.Println("VisitDestroy...")
	indexName := ctx.IndexName().GetText()
	fmt.Printf("indexName: %s\n", indexName)
	return fmt.Sprintf("Destroy %s", indexName)
}

//POC of visitor
func TestCqlVisitor(t *testing.T) {
	fmt.Println("================TestCqlVisitor================")

	input := antlr.NewInputStream("IDX.CREATE orders SCHEMA object UINT64 price FLOAT number UINT32 desc STRING date UINT64")
	//input := antlr.NewInputStream("IDX.DESTROY orders")
	lexer := NewCQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewCQLParser(stream)
	el := new(VerboseErrorListener)
	parser.AddErrorListener(el)
	//parser.BuildParseTrees = true

	tree := parser.Cql()
	if el.exp != nil {
		fmt.Printf("parser raised exception %+v\n", el.exp)
		t.Fatalf(el.msg)
	}

	visitor := new(CqlTestVisitor)
	tree.Accept(visitor)
	fmt.Printf("the result of visitor: %v\n", visitor.res)
}
