// Copyright 2015 PingCAP, Inc.
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

package resolver

import (
	"github.com/chrislusf/gleam/sql/ast"
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/infoschema"
	"github.com/chrislusf/gleam/sql/model"
)

// ResolveName resolves table name and column name.
// It generates ResultFields for ResultSetNode and resolves ColumnNameExpr to a ResultField.
func ResolveName(node ast.Node, info infoschema.InfoSchema, ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// nameResolver is the visitor to resolve table name and column name.
// In general, a reference can only refer to information that are available for it.
// So children elements are visited in the order that previous elements make information
// available for following elements.
//
// During visiting, information are collected and stored in resolverContext.
// When we enter a subquery, a new resolverContext is pushed to the contextStack, so subquery
// information can overwrite outer query information. When we look up for a column reference,
// we look up from top to bottom in the contextStack.
type nameResolver struct {
	Info            infoschema.InfoSchema
	Ctx             context.Context
	DefaultSchema   model.CIStr
	Err             error
	useOuterContext bool

	contextStack []*resolverContext
}

// resolverContext stores information in a single level of select statement
// that table name and column name can be resolved.
type resolverContext struct {
	/* For Select Statement. */
	// table map to lookup and check table name conflict.
	tableMap map[string]int
	// table map to lookup and check derived-table(subselect) name conflict.
	derivedTableMap map[string]int
	// tableSources collected in from clause.
	tables []*ast.TableSource
	// result fields collected in select field list.
	fieldList []*ast.ResultField
	// result fields collected in group by clause.
	groupBy []*ast.ResultField

	// The join node stack is used by on condition to find out
	// available tables to reference. On condition can only
	// refer to tables involved in current join.
	joinNodeStack []*ast.Join

	// When visiting TableRefs, tables in this context are not available
	// because it is being collected.
	inTableRefs bool
	// When visiting on condition only tables in current join node are available.
	inOnCondition bool
	// When visiting field list, fieldList in this context are not available.
	inFieldList bool
	// When visiting group by, groupBy fields are not available.
	inGroupBy bool
	// When visiting having, only fieldList and groupBy fields are available.
	inHaving bool
	// When visiting having, checks if the expr is an aggregate function expr.
	inHavingAgg bool
	// OrderBy clause has different resolving rule than group by.
	inOrderBy bool
	// When visiting column name in ByItem, we should know if the column name is in an expression.
	inByItemExpression bool
	// If subquery use outer context.
	useOuterContext bool
	// When visiting multi-table delete stmt table list.
	inDeleteTableList bool
	// When visiting create/drop table statement.
	inCreateOrDropTable bool
	// When visiting show statement.
	inShow bool
}

// currentContext gets the current resolverContext.
func (nr *nameResolver) currentContext() *resolverContext { _ = "STUB: not implemented"; return nil }

// pushContext is called when we enter a statement.
func (nr *nameResolver) pushContext() { _ = "STUB: not implemented"; return }

// popContext is called when we leave a statement.
func (nr *nameResolver) popContext() { _ = "STUB: not implemented"; return }

// pushJoin is called when we enter a join node.
func (nr *nameResolver) pushJoin(j *ast.Join) { _ = "STUB: not implemented"; return }

// popJoin is called when we leave a join node.
func (nr *nameResolver) popJoin() { _ = "STUB: not implemented"; return }

// Enter implements ast.Visitor interface.
func (nr *nameResolver) Enter(inNode ast.Node) (outNode ast.Node, skipChildren bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

// If ByItem is not a single column name expression,
// the resolving rule is different from order by clause.

// make sure item is not aggregate function

// Convert column name expression to string value expression.

// Leave implements ast.Visitor interface.
func (nr *nameResolver) Leave(inNode ast.Node) (node ast.Node, ok bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

// TODO: check this
// If there is a deep nest of subquery, there may be something wrong.

// handleTableName looks up and sets the schema information and result fields for table name.
func (nr *nameResolver) handleTableName(tn *ast.TableName) { _ = "STUB: not implemented"; return }

// The table may not exist in create table or drop table statement.
// Skip resolving the table to avoid error.

// handleTableSources checks name duplication
// and puts the table source in current resolverContext.
// Note:
// "select * from t as a join (select 1) as a;" is not duplicate.
// "select * from t as a join t as a;" is duplicate.
// "select * from (select 1) as a join (select 1) as a;" is duplicate.
func (nr *nameResolver) handleTableSource(ts *ast.TableSource) { _ = "STUB: not implemented"; return }

// duplicate column name in one table is not allowed.
// "select * from (select 1, 1) as a;" is duplicate.

// handleJoin sets result fields for join.
func (nr *nameResolver) handleJoin(j *ast.Join) { _ = "STUB: not implemented"; return }

// handleColumnName looks up and sets ResultField for
// the column name.
func (nr *nameResolver) handleColumnName(cn *ast.ColumnNameExpr) { _ = "STUB: not implemented"; return }

// In on condition, only tables within current join is available.

// Try to resolve the column name form top to bottom in the context stack.

// Column is already resolved or encountered an error.

// If in subselect, the query use outer query.

// resolveColumnNameInContext looks up and sets ResultField for a column with the ctx.
func (nr *nameResolver) resolveColumnNameInContext(ctx *resolverContext, cn *ast.ColumnNameExpr) bool {
	_ = "STUB: not implemented"
	return false

	// In TableRefsClause, column reference only in join on condition which is handled before.
}

// only resolve column using tables.

// From tables first, then field list.
// If ctx.InByItemExpression is true, the item is not an identifier.
// Otherwise it is an identifier.

// From table first, then field list.

// Check if resolved refer is an aggregate function expr.

// Resolve from table first, then from select list.

// We should copy the refer here.
// Because if the ByItem is an identifier, we should check if it
// is ambiguous even it is already resolved from table source.
// If the ByItem is not an identifier, we do not need the second check.

// It is not ambiguous and already resolved from table source.
// We should restore its Refer.

// First group by, then field list.

// If cn is in an aggregate function in having clause, check tablesource first.

// From table first, then field list.

// Field list first, then from table.

// In where clause.

// resolveColumnNameInOnCondition resolves the column name in current join.
func (nr *nameResolver) resolveColumnNameInOnCondition(cn *ast.ColumnNameExpr) {
	_ = "STUB: not implemented"
	return
}

func (nr *nameResolver) resolveColumnInTableSources(cn *ast.ColumnNameExpr, tableSources []*ast.TableSource) (done bool) {
	_ = "STUB: not implemented"
	return false
}

// different table name.

// Table as name shadows table real name.

// resolve column.

// Bind column.

func (nr *nameResolver) resolveColumnInResultFields(ctx *resolverContext, cn *ast.ColumnNameExpr, rfs []*ast.ResultField) bool {
	_ = "STUB: not implemented"
	return false
}

// Check table name

// This is not a real table column, resolve it directly.

// If in GroupBy, we clone the ResultField

// Bind column.

// handleFieldList expands wild card field and sets fieldList in current context.
func (nr *nameResolver) handleFieldList(fieldList *ast.FieldList) {
	_ = "STUB: not implemented"
	return
}

func getInnerFromParentheses(expr ast.ExprNode) ast.ExprNode {
	_ = "STUB: not implemented"
	return *new(ast.ExprNode)
}

// createResultFields creates result field list for a single select field.
func (nr *nameResolver) createResultFields(field *ast.SelectField) (rfs []*ast.ResultField) {
	_ = "STUB: not implemented"
	return nil
}

// Convert it to ColumnNameExpr

// The column is visited before so it must has been resolved already.

// Empty column info.
// Empty table info.

func appendTableSources(in []*ast.TableSource, resultSetNode ast.ResultSetNode) (out []*ast.TableSource) {
	_ = "STUB: not implemented"
	return nil
}

func (nr *nameResolver) tableUniqueName(schema, table model.CIStr) string {
	_ = "STUB: not implemented"
	return ""
}

func (nr *nameResolver) handlePosition(pos *ast.PositionExpr) { _ = "STUB: not implemented"; return }

// make sure item is not aggregate function

func (nr *nameResolver) handleUnionSelectList(u *ast.UnionSelectList) {
	_ = "STUB: not implemented"
	return
}

// Copy first result fields, because we may change the result field type.

func (nr *nameResolver) fillShowFields(s *ast.ShowStmt) { _ = "STUB: not implemented"; return }

// Empty column info.
// Empty table info.

// use varchar as the default return column type
