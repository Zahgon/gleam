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
	"github.com/chrislusf/gleam/sql/model"
	"github.com/chrislusf/gleam/sql/util/types"
)

var fullRange = []rangePoint{
	{start: true},
	{value: types.MaxValueDatum()},
}

// getEQFunctionOffset judge if the expression is a eq function like A = 1 where a is an index.
// If so, it will return the offset of A in index columns. e.g. for index(C,B,A), A's offset is 2.
func getEQFunctionOffset(expr expression.Expression, cols []*model.IndexColumn) int {
	_ = "STUB: not implemented"
	return 0
}

func removeAccessConditions(conditions, accessConds []expression.Expression) []expression.Expression {
	_ = "STUB: not implemented"
	return nil
}

// detachTableScanConditions distinguishes between access conditions and filter conditions from conditions.
func detachTableScanConditions(conditions []expression.Expression, table *model.TableInfo) ([]expression.Expression, []expression.Expression) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: it will lead to repeated computation cost.

// conditionChecker checks if this condition can be pushed to index plan.
type conditionChecker struct {
	tableName     model.CIStr
	idx           *model.IndexInfo
	columnOffset  int // the offset of the indexed column to be checked.
	pkName        model.CIStr
	shouldReserve bool // check if a access condition should be reserved in filter conditions.
}

func (c *conditionChecker) check(condition expression.Expression) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *conditionChecker) extractAccessAndFilterConds(conditions, accessConds, filterConds []expression.Expression) ([]expression.Expression, []expression.Expression) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *conditionChecker) findEqOrInFunc(conditions []expression.Expression) int {
	_ = "STUB: not implemented"
	return 0
}

func (c *conditionChecker) checkScalarFunction(scalar *expression.ScalarFunction) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO: support "not like" and "not in" convert to access conditions.

// "not column" or "not constant" can't lead to a range.

func (c *conditionChecker) checkLikeFunc(scalar *expression.ScalarFunction) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *conditionChecker) checkColumn(expr expression.Expression) bool {
	_ = "STUB: not implemented"
	return false
}

var oppositeOp = map[string]string{
	ast.LT: ast.GE,
	ast.GE: ast.LT,
	ast.GT: ast.LE,
	ast.LE: ast.GT,
	ast.EQ: ast.NE,
	ast.NE: ast.EQ,
}

func pushDownNot(expr expression.Expression, not bool, ctx context.Context) expression.Expression {
	_ = "STUB: not implemented"
	return *new(expression.Expression)
}
