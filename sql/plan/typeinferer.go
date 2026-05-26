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
	"github.com/chrislusf/gleam/sql/sessionctx/variable"
)

// InferType infers result type for ast.ExprNode.
func InferType(sc *variable.StatementContext, node ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: get the default charset from ctx

type typeInferrer struct {
	sc             *variable.StatementContext
	err            error
	defaultCharset string
}

func (v *typeInferrer) Enter(in ast.Node) (out ast.Node, skipChildren bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

func (v *typeInferrer) Leave(in ast.Node) (out ast.Node, ok bool) {
	_ = "STUB: not implemented"
	return *new(ast.Node), false
}

// Copy a new field type.

// TODO: handle all expression types.

func (v *typeInferrer) selectStmt(x *ast.SelectStmt) { _ = "STUB: not implemented"; return }

// column ID is 0 means it is not a real column from table, but a temporary column,
// so its type is not pre-defined, we need to set it.

func (v *typeInferrer) aggregateFunc(x *ast.AggregateFuncExpr) { _ = "STUB: not implemented"; return }

func (v *typeInferrer) binaryOperation(x *ast.BinaryOperationExpr) {
	_ = "STUB: not implemented"
	return
}

// If both operands are unsigned, result is unsigned.

func mergeArithType(a, b byte) byte { _ = "STUB: not implemented"; return 0 }

func (v *typeInferrer) unaryOperation(x *ast.UnaryOperationExpr) { _ = "STUB: not implemented"; return }

func (v *typeInferrer) handleValueExpr(x *ast.ValueExpr) { _ = "STUB: not implemented"; return }

func (v *typeInferrer) handleValuesExpr(x *ast.ValuesExpr) { _ = "STUB: not implemented"; return }

func (v *typeInferrer) getFsp(x *ast.FuncCallExpr) int { _ = "STUB: not implemented"; return 0 }

func (v *typeInferrer) handleFuncCallExpr(x *ast.FuncCallExpr) { _ = "STUB: not implemented"; return }

// TODO: We should cover all types.

// TODO: fix this
// See https://dev.mysql.com/doc/refman/5.5/en/control-flow-functions.html#function_if
// The default return type of IF() (which may matter when it is stored into a temporary table) is calculated as follows.
// Expression	Return Value
// expr2 or expr3 returns a string	string
// expr2 or expr3 returns a floating-point value	floating-point
// expr2 or expr3 returns an integer	integer

// If charset is unspecified.

// The return type of a CASE expression is the compatible aggregated type of all return values,
// but also depends on the context in which it is used.
// If used in a string context, the result is returned as a string.
// If used in a numeric context, the result is returned as a decimal, real, or integer value.
func (v *typeInferrer) handleCaseExpr(x *ast.CaseExpr) { _ = "STUB: not implemented"; return }

// TODO: We need a better way to set charset/collation

// like expression expects the target expression and pattern to be a string, if it's not, we add a cast function.
func (v *typeInferrer) handleLikeExpr(x *ast.PatternLikeExpr) { _ = "STUB: not implemented"; return }

// regexp expression expects the target expression and pattern to be a string, if it's not, we add a cast function.
func (v *typeInferrer) handleRegexpExpr(x *ast.PatternRegexpExpr) {
	_ = "STUB: not implemented"
	return
}

// AddCastToString adds a cast function to string type if the expr charset is not UTF8.
func (v *typeInferrer) addCastToString(expr ast.ExprNode) ast.ExprNode {
	_ = "STUB: not implemented"
	return *new(ast.ExprNode)
}

// ConvertValueToColumnTypeIfNeeded checks if the expr in PatternInExpr is column name,
// and casts function to the items in the list.
func (v *typeInferrer) convertValueToColumnTypeIfNeeded(x *ast.PatternInExpr) {
	_ = "STUB: not implemented"
	return
}

// The value will never match the column, do not set newDatum.

// TODO: Errors should be handled differently according to query context.
