// Generated from CQL.g4 by ANTLR 4.7.

package parser // CQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CQLParser.
type CQLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CQLParser#cql.
	VisitCql(ctx *CqlContext) interface{}

	// Visit a parse tree produced by CQLParser#create.
	VisitCreate(ctx *CreateContext) interface{}

	// Visit a parse tree produced by CQLParser#destroy.
	VisitDestroy(ctx *DestroyContext) interface{}

	// Visit a parse tree produced by CQLParser#insert.
	VisitInsert(ctx *InsertContext) interface{}

	// Visit a parse tree produced by CQLParser#del.
	VisitDel(ctx *DelContext) interface{}

	// Visit a parse tree produced by CQLParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by CQLParser#indexName.
	VisitIndexName(ctx *IndexNameContext) interface{}

	// Visit a parse tree produced by CQLParser#propertyDef.
	VisitPropertyDef(ctx *PropertyDefContext) interface{}

	// Visit a parse tree produced by CQLParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by CQLParser#popType.
	VisitPopType(ctx *PopTypeContext) interface{}

	// Visit a parse tree produced by CQLParser#docId.
	VisitDocId(ctx *DocIdContext) interface{}

	// Visit a parse tree produced by CQLParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by CQLParser#predicates.
	VisitPredicates(ctx *PredicatesContext) interface{}

	// Visit a parse tree produced by CQLParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by CQLParser#relate.
	VisitRelate(ctx *RelateContext) interface{}

	// Visit a parse tree produced by CQLParser#compare.
	VisitCompare(ctx *CompareContext) interface{}

	// Visit a parse tree produced by CQLParser#limit.
	VisitLimit(ctx *LimitContext) interface{}
}
