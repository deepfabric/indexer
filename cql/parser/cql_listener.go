// Generated from CQL.g4 by ANTLR 4.7.

package parser // CQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CQLListener is a complete listener for a parse tree produced by CQLParser.
type CQLListener interface {
	antlr.ParseTreeListener

	// EnterCql is called when entering the cql production.
	EnterCql(c *CqlContext)

	// EnterCreate is called when entering the create production.
	EnterCreate(c *CreateContext)

	// EnterDestroy is called when entering the destroy production.
	EnterDestroy(c *DestroyContext)

	// EnterInsert is called when entering the insert production.
	EnterInsert(c *InsertContext)

	// EnterDel is called when entering the del production.
	EnterDel(c *DelContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterIndexName is called when entering the indexName production.
	EnterIndexName(c *IndexNameContext)

	// EnterPropertyDef is called when entering the propertyDef production.
	EnterPropertyDef(c *PropertyDefContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterPopType is called when entering the popType production.
	EnterPopType(c *PopTypeContext)

	// EnterDocId is called when entering the docId production.
	EnterDocId(c *DocIdContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterPredicates is called when entering the predicates production.
	EnterPredicates(c *PredicatesContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterRelate is called when entering the relate production.
	EnterRelate(c *RelateContext)

	// EnterCompare is called when entering the compare production.
	EnterCompare(c *CompareContext)

	// EnterLimit is called when entering the limit production.
	EnterLimit(c *LimitContext)

	// ExitCql is called when exiting the cql production.
	ExitCql(c *CqlContext)

	// ExitCreate is called when exiting the create production.
	ExitCreate(c *CreateContext)

	// ExitDestroy is called when exiting the destroy production.
	ExitDestroy(c *DestroyContext)

	// ExitInsert is called when exiting the insert production.
	ExitInsert(c *InsertContext)

	// ExitDel is called when exiting the del production.
	ExitDel(c *DelContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitIndexName is called when exiting the indexName production.
	ExitIndexName(c *IndexNameContext)

	// ExitPropertyDef is called when exiting the propertyDef production.
	ExitPropertyDef(c *PropertyDefContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitPopType is called when exiting the popType production.
	ExitPopType(c *PopTypeContext)

	// ExitDocId is called when exiting the docId production.
	ExitDocId(c *DocIdContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitPredicates is called when exiting the predicates production.
	ExitPredicates(c *PredicatesContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitRelate is called when exiting the relate production.
	ExitRelate(c *RelateContext)

	// ExitCompare is called when exiting the compare production.
	ExitCompare(c *CompareContext)

	// ExitLimit is called when exiting the limit production.
	ExitLimit(c *LimitContext)
}
