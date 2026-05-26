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
	"github.com/chrislusf/gleam/sql/expression"
	"github.com/chrislusf/gleam/sql/sessionctx/variable"
)

type idAllocator struct {
	id int
}

func (a *idAllocator) allocID() string { _ = "STUB: not implemented"; return "" }

func (p *Aggregation) collectGroupByColumns() { _ = "STUB: not implemented"; return }

func (b *planBuilder) buildAggregation(p LogicalPlan, aggFuncList []*ast.AggregateFuncExpr, gbyItems []expression.Expression) (LogicalPlan, map[int]int) {
	_ = "STUB: not implemented"
	return *new(LogicalPlan), nil
}

// aggIdxMap maps the old index to new index after applying common aggregation functions elimination.

func (b *planBuilder) buildResultSetNode(node ast.ResultSetNode) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

func extractCorColumns(expr expression.Expression) (cols []*expression.CorrelatedColumn) {
	_ = "STUB: not implemented"
	return nil
}

func extractOnCondition(conditions []expression.Expression, left LogicalPlan, right LogicalPlan) (
	eqCond []*expression.ScalarFunction, leftCond []expression.Expression, rightCond []expression.Expression,
	otherCond []expression.Expression) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (b *planBuilder) buildJoin(join *ast.Join) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

func (b *planBuilder) buildSelection(p LogicalPlan, where ast.ExprNode, AggMapper map[*ast.AggregateFuncExpr]int) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

// buildProjection returns a Projection plan and non-aux columns length.
func (b *planBuilder) buildProjection(p LogicalPlan, fields []*ast.SelectField, mapper map[*ast.AggregateFuncExpr]int) (LogicalPlan, int) {
	_ = "STUB: not implemented"
	return *new(LogicalPlan), 0
}

// When the query is select t.a from t group by a; The Column Name should be a but not t.a;

func (b *planBuilder) buildDistinct(child LogicalPlan, length int) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

func (b *planBuilder) buildUnion(union *ast.UnionStmt) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

/*
 * The lengths of the columns in the UNION result take into account the values retrieved by all of the SELECT statements
 * SELECT REPEAT('a',1) UNION SELECT REPEAT('b',10);
 * +---------------+
 * | REPEAT('a',1) |
 * +---------------+
 * | a             |
 * | bbbbbbbbbb    |
 * +---------------+
 */

// For select nul union select "abc", we should not convert "abc" to nil.
// And the result field type should be VARCHAR.

// ByItems wraps a "by" item.
type ByItems struct {
	Expr expression.Expression
	Desc bool
}

// String implements fmt.Stringer interface.
func (by *ByItems) String() string { _ = "STUB: not implemented"; return "" }

func (b *planBuilder) buildSort(p LogicalPlan, byItems []*ast.ByItem, aggMapper map[*ast.AggregateFuncExpr]int) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

// getUintForLimitOffset gets uint64 value for limit/offset.
// For ordinary statement, limit/offset should be uint64 constant value.
// For prepared statement, limit/offset is string. We should convert it to uint64.
func getUintForLimitOffset(sc *variable.StatementContext, val interface{}) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *planBuilder) buildLimit(src LogicalPlan, limit *ast.Limit) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

// colMatch(a,b) means that if a match b, e.g. t.a can match test.t.a bug test.t.a can't match t.a.
// Because column a want column from database test exactly.
func colMatch(a *ast.ColumnName, b *ast.ColumnName) bool { _ = "STUB: not implemented"; return false }

func matchField(f *ast.SelectField, col *ast.ColumnNameExpr, ignoreAsName bool) bool {
	_ = "STUB: not implemented"
	// if col specify a table name, resolve from table source directly.
	return false
}

// a expression without as name can't be matched.

func resolveFromSelectFields(v *ast.ColumnNameExpr, fields []*ast.SelectField, ignoreAsName bool) (index int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// AggregateFuncExtractor visits Expr tree.
// It converts ColunmNameExpr to AggregateFuncExpr and collects AggregateFuncExpr.
type havingAndOrderbyExprResolver struct {
	inAggFunc    bool
	inExpr       bool
	orderBy      bool
	err          error
	p            LogicalPlan
	selectFields []*ast.SelectField
	aggMapper    map[*ast.AggregateFuncExpr]int
	colMapper    map[*ast.ColumnNameExpr]int
	gbyItems     []*ast.ByItem
	outerSchemas []expression.Schema
}

// Enter implements Visitor interface.
func (a *havingAndOrderbyExprResolver) Enter(n ast.Node) (node ast.Node, skipChildren bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

// Enter a new context, skip it.
// For example: select sum(c) + c + exists(select c from t) from t;

func (a *havingAndOrderbyExprResolver) resolveFromSchema(v *ast.ColumnNameExpr, schema expression.Schema) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Leave implements Visitor interface.
func (a *havingAndOrderbyExprResolver) Leave(n ast.Node) (node ast.Node, ok bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

// We should ignore the err when resolving from schema. Because we could resolve successfully
// when considering select fields.

// If we can't find it any where, it may be a correlated columns.

func (b *planBuilder) resolveHavingAndOrderBy(sel *ast.SelectStmt, p LogicalPlan) (
	map[*ast.AggregateFuncExpr]int, map[*ast.AggregateFuncExpr]int) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract agg funcs from having clause.

// Extract agg funcs from order by clause.

func (b *planBuilder) extractAggFuncs(fields []*ast.SelectField) ([]*ast.AggregateFuncExpr, map[*ast.AggregateFuncExpr]int) {
	_ = "STUB: not implemented"
	return nil, nil
}

// gbyResolver resolves group by items from select fields.
type gbyResolver struct {
	fields []*ast.SelectField
	schema expression.Schema
	err    error
	inExpr bool
}

func (g *gbyResolver) Enter(inNode ast.Node) (ast.Node, bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

func (g *gbyResolver) Leave(inNode ast.Node) (ast.Node, bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

func (b *planBuilder) resolveGbyExprs(p LogicalPlan, gby *ast.GroupByClause, fields []*ast.SelectField) (LogicalPlan, []expression.Expression) {
	_ = "STUB: not implemented"
	return *new(LogicalPlan), nil
}

func (b *planBuilder) unfoldWildStar(p LogicalPlan, selectFields []*ast.SelectField) (resultList []*ast.SelectField) {
	_ = "STUB: not implemented"
	return nil
}

func (b *planBuilder) buildSelect(sel *ast.SelectStmt) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

// We must resolve having and order by clause before build projection,
// because when the query is "select a+1 as b from t having sum(b) < 0", we must replace sum(b) to sum(a+1),
// which only can be done before building projection and extracting Agg functions.

func (b *planBuilder) buildTrim(p LogicalPlan, len int) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

func (b *planBuilder) buildTableDual() LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

func (b *planBuilder) buildDataSource(tn *ast.TableName) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

// Equal condition contains a column from previous joined table.

// ApplyConditionChecker checks whether all or any output of apply matches a condition.
type ApplyConditionChecker struct {
	Condition expression.Expression
	All       bool
}

// buildInnerApply builds apply plan with outerPlan and innerPlan, which apply inner-join for every row from outerPlan and the whole innerPlan.
func (b *planBuilder) buildInnerApply(outerPlan, innerPlan LogicalPlan) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

// buildSemiApply builds apply plan with outerPlan and innerPlan, which apply semi-join for every row from outerPlan and the whole innerPlan.
func (b *planBuilder) buildSemiApply(outerPlan, innerPlan LogicalPlan, condition []expression.Expression, asScalar, not bool) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

func (b *planBuilder) buildExists(p LogicalPlan) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

// This can be removed when in exists clause,
// e.g. exists(select count(*) from t order by a) is equal to exists t.

func (b *planBuilder) buildMaxOneRow(p LogicalPlan) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

func (b *planBuilder) buildSemiJoin(outerPlan, innerPlan LogicalPlan, onCondition []expression.Expression, asScalar bool, not bool) *Join {
	_ = "STUB: not implemented"
	return nil
}

func (b *planBuilder) buildUpdate(update *ast.UpdateStmt) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

func (b *planBuilder) buildUpdateLists(list []*ast.Assignment, p LogicalPlan) ([]*expression.Assignment, LogicalPlan) {
	_ = "STUB: not implemented"
	return nil, *new(LogicalPlan)
}

func (b *planBuilder) buildDelete(delete *ast.DeleteStmt) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

func getInnerFromParentheses(expr ast.ExprNode) ast.ExprNode {
	_ = "STUB: not implemented"
	return *new(ast.ExprNode)
}
