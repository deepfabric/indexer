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
		_ = parser.Cql()
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
}

func (v *CqlTestVisitor) Visit(tree antlr.ParseTree) interface{} {
	fmt.Println("Visit...")
	ctx := tree.(*CqlContext)
	return v.VisitCql(ctx)
}

func (v *CqlTestVisitor) VisitCql(ctx *CqlContext) interface{} {
	fmt.Printf("VisitCql %v...\n", ctx)
	//If there are multiple subrules, than check one by one.
	if create := ctx.Create(); create != nil {
		return v.VisitCreate(create.(*CreateContext))
	} else if destroy := ctx.Destroy(); destroy != nil {
		return v.VisitDestroy(destroy.(*DestroyContext))
	}
	return nil
}

func (v *CqlTestVisitor) VisitCreate(ctx *CreateContext) interface{} {
	fmt.Println("VisitCreate...")
	indexName := ctx.IndexName().GetText()
	fmt.Printf("indexName: %s\n", indexName)
	return nil
}

func (v *CqlTestVisitor) VisitDestroy(ctx *DestroyContext) interface{} {
	fmt.Println("VisitDestroy...")
	indexName := ctx.IndexName().GetText()
	fmt.Printf("indexName: %s\n", indexName)
	return nil
}

//POC of visitor
func TestCqlVisitor(t *testing.T) {
	fmt.Println("================TestCqlVisitor================")

	//input := antlr.NewInputStream("IDX.CREATE orders SCHEMA object UINT64 price FLOAT number UINT32 date UINT64")
	input := antlr.NewInputStream("IDX.DESTROY orders")
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
}
