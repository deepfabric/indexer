// Generated from /home/zhichyu/src/github.com/deepfabric/indexer/cql/parser/CQL.g4 by ANTLR 4.7.

package parser // CQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCQLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCQLVisitor) VisitCql(ctx *CqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitCreate(ctx *CreateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitDestroy(ctx *DestroyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitInsert(ctx *InsertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitDel(ctx *DelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitIndexName(ctx *IndexNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitPropertyDef(ctx *PropertyDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitPopType(ctx *PopTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitDocId(ctx *DocIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitPredicates(ctx *PredicatesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitRelate(ctx *RelateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitCompare(ctx *CompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCQLVisitor) VisitLimit(ctx *LimitContext) interface{} {
	return v.VisitChildren(ctx)
}
