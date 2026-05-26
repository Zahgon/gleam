// Copyright 2016 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package plan

import (
	"github.com/chrislusf/gleam/sql/ast"
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/expression"
	"github.com/chrislusf/gleam/sql/infoschema"
	"github.com/chrislusf/gleam/sql/util/types"
)

// EvalSubquery evaluates incorrelated subqueries once.
var EvalSubquery func(p PhysicalPlan, is infoschema.InfoSchema, ctx context.Context) ([]types.Datum, error)

// evalAstExpr evaluates ast expression directly.
func evalAstExpr(expr ast.ExprNode, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// rewrite function rewrites ast expr to expression.Expression.
// aggMapper maps ast.AggregateFuncExpr to the columns offset in p's output schema.
// asScalar means whether this expression must be treated as a scalar expression.
// And this function returns a result expression, a new plan that may have apply or semi-join.
func (b *planBuilder) rewrite(expr ast.ExprNode, p LogicalPlan, aggMapper map[*ast.AggregateFuncExpr]int, asScalar bool) (
	expression.Expression, LogicalPlan, error) {
	_ = "STUB: not implemented"
	return *new(expression.Expression), *new(LogicalPlan), nil
}

type expressionRewriter struct {
	ctxStack []expression.Expression
	p        LogicalPlan
	schema   expression.Schema
	err      error
	aggrMap  map[*ast.AggregateFuncExpr]int
	b        *planBuilder
	ctx      context.Context
	// asScalar means the return value must be a scalar value.
	asScalar bool
}

func getRowLen(e expression.Expression) int { _ = "STUB: not implemented"; return 0 }

func getRowArg(e expression.Expression, idx int) expression.Expression {
	_ = "STUB: not implemented"
	return *new(expression.Expression)
}

// constructBinaryOpFunctions converts (a0,a1,a2) op (b0,b1,b2) to (a0 op b0) and (a1 op b1) and (a2 op b2).
func (er *expressionRewriter) constructBinaryOpFunction(l expression.Expression, r expression.Expression, op string) (expression.Expression, error) {
	_ = "STUB: not implemented"
	return *new(expression.Expression), nil
}

func (er *expressionRewriter) buildSubquery(subq *ast.SubqueryExpr) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

// Enter implements Visitor interface.
func (er *expressionRewriter) Enter(inNode ast.Node) (ast.Node, bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

// For 10 in ((select * from t)), the parser won't set v.Sel.
// So we must process this case here.

func (er *expressionRewriter) handleCompareSubquery(v *ast.CompareSubqueryExpr) (ast.Node, bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

// Only (a,b,c) = all (...) and (a,b,c) != any () can use row expression.

// Only EQ, NE and NullEQ can be composed with and.

// TODO: Support this in future.

// When < all or > any , the agg function should use min.

// The parent expression only use the last column in schema, which represents whether the condition is matched.

// handleOtherComparableSubq handles the queries like < any, < max, etc. For example, if the query is t.id < any (select s.id from s),
// it will be rewrote to t.id < (select max(s.id) from s).
func (er *expressionRewriter) handleOtherComparableSubq(lexpr, rexpr expression.Expression, np LogicalPlan, useMin bool, cmpFunc string, all bool) {
	_ = "STUB: not implemented"
	return
}

// buildQuantifierPlan adds extra condition for any / all subquery.
func (er *expressionRewriter) buildQuantifierPlan(agg *Aggregation, cond, rexpr expression.Expression, all bool) {
	_ = "STUB: not implemented"
	return
}

// All of the inner record set should not contain null value. So for t.id < all(select s.id from s), it
// should be rewrote to t.id < min(s.id) and if(sum(s.id is null) = 0, true, null).

// If the set is empty, it should always return true.

// For "any" expression, if the record set has null and the cond return false, the result should be NULL.

// For Semi Apply without aux column, the result is no matter false or null. So we can add it to join predicate.

// If we treat the result as a scalar value, we will add a projection with a extra column to output true, false or null.

// handleNEAny handles the case of != any. For exmaple, if the query is t.id != any (select s.id from s), it will be rewrote to
// t.id != s.id or count(distinct s.id) > 1 or [any checker]. If there are two different values in s.id ,
// there must exist a s.id that doesn't equal to t.id.
func (er *expressionRewriter) handleNEAny(lexpr, rexpr expression.Expression, np LogicalPlan) {
	_ = "STUB: not implemented"
	return
}

// handleEQAll handles the case of = all. For example, if the query is t.id = all (select s.id from s), it will be rewrote to
// t.id = (select s.id from s having count(distinct s.id) <= 1 and [all checker]).
func (er *expressionRewriter) handleEQAll(lexpr, rexpr expression.Expression, np LogicalPlan) {
	_ = "STUB: not implemented"
	return
}

func (er *expressionRewriter) handleExistSubquery(v *ast.ExistsSubqueryExpr) (ast.Node, bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

func (er *expressionRewriter) handleInSubquery(v *ast.PatternInExpr) (ast.Node, bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

// a in (subq) will be rewrote as a = any(subq).
// a not in (subq) will be rewrote as a != all(subq).

func (er *expressionRewriter) handleScalarSubquery(v *ast.SubqueryExpr) (ast.Node, bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

// Leave implements Visitor interface.
func (er *expressionRewriter) Leave(inNode ast.Node) (retNode ast.Node, ok bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

func datumToConstant(d types.Datum, tp byte) *expression.Constant {
	_ = "STUB: not implemented"
	return nil
}

func (er *expressionRewriter) rewriteVariable(v *ast.VariableExpr) {
	_ = "STUB: not implemented"
	return
}

// TODO: Here is wrong, the sessionVars should store a name -> Datum map. Will fix it later.

// select null user vars is permitted.

func (er *expressionRewriter) unaryOpToExpression(v *ast.UnaryOperationExpr) {
	_ = "STUB: not implemented"
	return
}

// expression (+ a) is equal to a

func (er *expressionRewriter) binaryOpToExpression(v *ast.BinaryOperationExpr) {
	_ = "STUB: not implemented"
	return
}

func (er *expressionRewriter) notToExpression(hasNot bool, op string, tp *types.FieldType,
	args ...expression.Expression) expression.Expression {
	_ = "STUB: not implemented"
	return *new(expression.Expression)
}

func (er *expressionRewriter) isNullToExpression(v *ast.IsNullExpr) {
	_ = "STUB: not implemented"
	return
}

func (er *expressionRewriter) positionToScalarFunc(v *ast.PositionExpr) {
	_ = "STUB: not implemented"
	return
}

func (er *expressionRewriter) isTrueToScalarFunc(v *ast.IsTruthExpr) {
	_ = "STUB: not implemented"
	return
}

func (er *expressionRewriter) inToExpression(v *ast.PatternInExpr) {
	_ = "STUB: not implemented"
	return
}

func (er *expressionRewriter) caseToExpression(v *ast.CaseExpr) { _ = "STUB: not implemented"; return }

// value                          -> ctxStack[stkLen-argsLen-1]
// when clause(condition, result) -> ctxStack[stkLen-argsLen:stkLen-1];
// else clause                    -> ctxStack[stkLen-1]

// args:  eq scalar func(args: value, condition1), result1,
//        eq scalar func(args: value, condition2), result2,
//        ...
//        else clause

// for trimming the value element later

// args:  condition1, result1,
//        condition2, result2,
//        ...
//        else clause

func (er *expressionRewriter) likeToScalarFunc(v *ast.PatternLikeExpr) {
	_ = "STUB: not implemented"
	return
}

func (er *expressionRewriter) regexpToScalarFunc(v *ast.PatternRegexpExpr) {
	_ = "STUB: not implemented"
	return
}

func (er *expressionRewriter) rowToScalarFunc(v *ast.RowExpr) { _ = "STUB: not implemented"; return }

func (er *expressionRewriter) betweenToExpression(v *ast.BetweenExpr) {
	_ = "STUB: not implemented"
	return
}

func (er *expressionRewriter) checkArgsOneColumn(args ...expression.Expression) {
	_ = "STUB: not implemented"
	return
}

func (er *expressionRewriter) funcCallToExpression(v *ast.FuncCallExpr) {
	_ = "STUB: not implemented"
	return
}

func (er *expressionRewriter) toColumn(v *ast.ColumnName) { _ = "STUB: not implemented"; return }
