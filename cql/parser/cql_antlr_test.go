package parser // CQL
import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

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
func TestCqlListenerCreate(t *testing.T) {
	fmt.Println("================TestCqlListenerCreate================")
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

func (v *CqlTestVisitor) Visit(localctx ICqlContext) interface{} {
	fmt.Println("Visit...")
	ctx := localctx.(*CqlContext)
	return v.VisitChildren(ctx)
}

func (v *CqlTestVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	fmt.Printf("VisitChildren %v\n", node)
	return nil
}

func (v *CqlTestVisitor) VisitCql(ctx *CqlContext) interface{} {
	fmt.Println("VisitCql...")
	return v.VisitChildren(ctx)
}

func (v *CqlTestVisitor) VisitCreate(ctx *CreateContext) interface{} {
	fmt.Println("VisitCreate...")
	return v.VisitChildren(ctx)
}

func TestCqlVisitorCreate(t *testing.T) {
	fmt.Println("================TestCqlVisitorCreate================")
	input := antlr.NewInputStream("IDX.CREATE orders SCHEMA object UINT64 price FLOAT number UINT32 date UINT64")
	lexer := NewCQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewCQLParser(stream)
	//parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	//parser.BuildParseTrees = true
	tree := parser.Cql()
	visitor := new(CqlTestVisitor)
	visitor.Visit(tree)
	//visitor.VisitCql(tree.(CqlContext))
}
