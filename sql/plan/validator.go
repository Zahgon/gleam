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
)

// Validate checkes whether the node is valid.
func Validate(node ast.Node, inPrepare bool) error { _ = "STUB: not implemented"; return nil }

// validator is an ast.Visitor that validates
// ast Nodes parsed from parser.
type validator struct {
	err           error
	wildCardCount int
	inPrepare     bool
	inAggregate   bool
}

func (v *validator) Enter(in ast.Node) (out ast.Node, skipChildren bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

// Aggregate function can not contain aggregate function.

func (v *validator) Leave(in ast.Node) (out ast.Node, ok bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

// We only accept ? and uint64 for count/offset in parser.y

func checkAutoIncrementOp(colDef *ast.ColumnDef, num int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func isConstraintKeyTp(constraints []*ast.Constraint, colDef *ast.ColumnDef) bool {
	_ = "STUB: not implemented"
	return false
}

// If the constraint as follows: primary key(c1, c2)
// we only support c1 column can be auto_increment.

func (v *validator) checkAutoIncrement(stmt *ast.CreateTableStmt) {
	_ = "STUB: not implemented"
	return
}

func (v *validator) checkCreateTableGrammar(stmt *ast.CreateTableStmt) {
	_ = "STUB: not implemented"
	return
}

func isPrimary(ops []*ast.ColumnOption) int { _ = "STUB: not implemented"; return 0 }

func (v *validator) checkCreateIndexGrammar(stmt *ast.CreateIndexStmt) {
	_ = "STUB: not implemented"
	return
}

func (v *validator) checkAlterTableGrammar(stmt *ast.AlterTableStmt) {
	_ = "STUB: not implemented"
	return
}

// Nothing to do now.

// Nothing to do now.

// checkDuplicateColumnName checks if index exists duplicated columns.
func checkDuplicateColumnName(indexColNames []*ast.IndexColName) error {
	_ = "STUB: not implemented"
	return nil
}
