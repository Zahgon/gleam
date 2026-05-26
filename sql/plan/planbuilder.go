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

package plan

import (
	"github.com/chrislusf/gleam/sql/ast"
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/expression"
	"github.com/chrislusf/gleam/sql/infoschema"
	"github.com/chrislusf/gleam/sql/model"
	"github.com/chrislusf/gleam/sql/mysql"
	"github.com/chrislusf/gleam/sql/table"
	"github.com/chrislusf/gleam/sql/terror"
)

// Error instances.
var (
	ErrUnsupportedType      = terror.ClassOptimizerPlan.New(CodeUnsupportedType, "Unsupported type")
	SystemInternalErrorType = terror.ClassOptimizerPlan.New(SystemInternalError, "System internal error")
	ErrUnknownColumn        = terror.ClassOptimizerPlan.New(CodeUnknownColumn, "Unknown column '%s' in '%s'")
	ErrWrongArguments       = terror.ClassOptimizerPlan.New(CodeWrongArguments, "Incorrect arguments to EXECUTE")
	ErrAmbiguous            = terror.ClassOptimizerPlan.New(CodeAmbiguous, "Column '%s' in field list is ambiguous")
)

// Error codes.
const (
	CodeUnsupportedType terror.ErrCode = 1
	SystemInternalError terror.ErrCode = 2
	CodeAmbiguous       terror.ErrCode = 1052
	CodeUnknownColumn   terror.ErrCode = 1054
	CodeWrongArguments  terror.ErrCode = 1210
)

func init() {
	tableMySQLErrCodes := map[terror.ErrCode]uint16{
		CodeUnknownColumn:  mysql.ErrBadField,
		CodeAmbiguous:      mysql.ErrNonUniq,
		CodeWrongArguments: mysql.ErrWrongArguments,
	}
	terror.ErrClassToMySQLCodes[terror.ClassOptimizerPlan] = tableMySQLErrCodes
}

// planBuilder builds Plan from an ast.Node.
// It just builds the ast node straightforwardly.
type planBuilder struct {
	err          error
	hasAgg       bool
	obj          interface{}
	allocator    *idAllocator
	ctx          context.Context
	is           infoschema.InfoSchema
	outerSchemas []expression.Schema
	inUpdateStmt bool
	// colMapper stores the column that must be pre-resolved.
	colMapper map[*ast.ColumnNameExpr]int
}

func (b *planBuilder) build(node ast.Node) Plan { _ = "STUB: not implemented"; return *new(Plan) }

func (b *planBuilder) buildSet(v *ast.SetStmt) Plan { _ = "STUB: not implemented"; return *new(Plan) }

// Detect aggregate function or groupby clause.
func (b *planBuilder) detectSelectAgg(sel *ast.SelectStmt) bool {
	_ = "STUB: not implemented"
	return false
}

func availableIndices(hints []*ast.IndexHint, tableInfo *model.TableInfo) (indices []*model.IndexInfo, includeTableScan bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Currently we don't distinguish between Force and Use because our cost estimation is not reliable.

// Collect all the ignore index hints.

// If we have got FORCE or USE index hint, table scan is excluded.

// Empty use hint means don't use any index.

func removeIgnores(indices, ignores []*model.IndexInfo) []*model.IndexInfo {
	_ = "STUB: not implemented"
	return nil
}

func findIndexByName(indices []*model.IndexInfo, name model.CIStr) *model.IndexInfo {
	_ = "STUB: not implemented"
	return nil
}

func (b *planBuilder) buildSelectLock(src Plan, lock ast.SelectLockType) *SelectLock {
	_ = "STUB: not implemented"
	return nil
}

func buildColumn(tableName, name string, tp byte, size int) *expression.Column {
	_ = "STUB: not implemented"
	return nil
}

// splitWhere split a where expression to a list of AND conditions.
func splitWhere(where ast.ExprNode) []ast.ExprNode { _ = "STUB: not implemented"; return nil }

func (b *planBuilder) getDefaultValue(col *table.Column) (*expression.Constant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *planBuilder) findDefaultValue(cols []*table.Column, name *ast.ColumnName) (*expression.Constant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getShowColNamesAndTypes gets column names and types. If the `ftypes` is empty, every column is set to varchar type.
func getShowColNamesAndTypes(s *ast.ShowStmt) (names []string, ftypes []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}
