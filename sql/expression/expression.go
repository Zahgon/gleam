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

package expression

import (
	"encoding/json"
	"fmt"

	"github.com/chrislusf/gleam/sql/ast"
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/model"
	"github.com/chrislusf/gleam/sql/mysql"
	"github.com/chrislusf/gleam/sql/terror"
	"github.com/chrislusf/gleam/sql/util/types"
)

// Error instances.
var (
	errInvalidOperation        = terror.ClassExpression.New(codeInvalidOperation, "invalid operation")
	errIncorrectParameterCount = terror.ClassExpression.New(codeIncorrectParameterCount, "Incorrect parameter count in the call to native function '%s'")
	errFunctionNotExists       = terror.ClassExpression.New(codeFunctionNotExists, "FUNCTION %s does not exist")
)

// Error codes.
const (
	codeInvalidOperation        terror.ErrCode = 1
	codeIncorrectParameterCount                = 1582
	codeFunctionNotExists                      = 1305
)

// EvalAstExpr evaluates ast expression directly.
var EvalAstExpr func(expr ast.ExprNode, ctx context.Context) (types.Datum, error)

// Expression represents all scalar expression in SQL.
type Expression interface {
	fmt.Stringer
	json.Marshaler
	// Eval evaluates an expression through a row.
	Eval(row []types.Datum, ctx context.Context) (types.Datum, error)

	// Get the expression return type.
	GetType() *types.FieldType

	// Clone copies an expression totally.
	Clone() Expression

	// HashCode create the hashcode for expression
	HashCode() []byte

	// Equal checks whether two expressions are equal.
	Equal(e Expression, ctx context.Context) bool

	// IsCorrelated checks if this expression has correlated key.
	IsCorrelated() bool

	// Decorrelate try to decorrelate the expression by schema.
	Decorrelate(schema Schema) Expression

	// ResolveIndices resolves indices by the given schema.
	ResolveIndices(schema Schema)
}

// EvalBool evaluates expression to a boolean value.
func EvalBool(expr Expression, row []types.Datum, ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// One stands for a number 1.
var One = &Constant{
	Value:   types.NewDatum(1),
	RetType: types.NewFieldType(mysql.TypeTiny),
}

// Zero stands for a number 0.
var Zero = &Constant{
	Value:   types.NewDatum(0),
	RetType: types.NewFieldType(mysql.TypeTiny),
}

// Null stands for null constant.
var Null = &Constant{
	Value:   types.NewDatum(nil),
	RetType: types.NewFieldType(mysql.TypeTiny),
}

// Constant stands for a constant value.
type Constant struct {
	Value   types.Datum
	RetType *types.FieldType
}

// String implements fmt.Stringer interface.
func (c *Constant) String() string { _ = "STUB: not implemented"; return "" }

// MarshalJSON implements json.Marshaler interface.
func (c *Constant) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Clone implements Expression interface.
func (c *Constant) Clone() Expression {
	_ = "STUB: not implemented"
	return *

	// GetType implements Expression interface.
	new(Expression)
}

func (c *Constant) GetType() *types.FieldType {
	_ = "STUB: not implemented"

	// Eval implements Expression interface.
	return nil
}

func (c *Constant) Eval(_ []types.Datum, _ context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *

	// Equal implements Expression interface.
	new(types.Datum), nil
}

func (c *Constant) Equal(b Expression, ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// IsCorrelated implements Expression interface.
func (c *Constant) IsCorrelated() bool {
	_ = "STUB: not implemented"

	// Decorrelate implements Expression interface.
	return false
}

func (c *Constant) Decorrelate(_ Schema) Expression {
	_ = "STUB: not implemented"

	// HashCode implements Expression interface.
	return *new(Expression)
}

func (c *Constant) HashCode() []byte { _ = "STUB: not implemented"; return nil }

// ResolveIndices implements Expression interface.
func (c *Constant) ResolveIndices(_ Schema) {
	_ = "STUB: not implemented"

	// composeConditionWithBinaryOp composes condition with binary operator into a balance deep tree, which benefits a lot for pb decoder/encoder.
	return
}

func composeConditionWithBinaryOp(ctx context.Context, conditions []Expression, funcName string) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// ComposeCNFCondition composes CNF items into a balance deep CNF tree, which benefits a lot for pb decoder/encoder.
func ComposeCNFCondition(ctx context.Context, conditions ...Expression) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// ComposeDNFCondition composes DNF items into a balance deep DNF tree.
func ComposeDNFCondition(ctx context.Context, conditions ...Expression) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// Assignment represents a set assignment in Update, such as
// Update t set c1 = hex(12), c2 = c3 where c2 = 1
type Assignment struct {
	Col  *Column
	Expr Expression
}

// VarAssignment represents a variable assignment in Set, such as set global a = 1.
type VarAssignment struct {
	Name        string
	Expr        Expression
	IsDefault   bool
	IsGlobal    bool
	IsSystem    bool
	ExtendValue *Constant
}

// splitNormalFormItems split CNF(conjunctive normal form) like "a and b and c", or DNF(disjunctive normal form) like "a or b or c"
func splitNormalFormItems(onExpr Expression, funcName string) []Expression {
	_ = "STUB: not implemented"
	return nil
}

// SplitCNFItems splits CNF items.
// CNF means conjunctive normal form, e.g. "a and b and c".
func SplitCNFItems(onExpr Expression) []Expression { _ = "STUB: not implemented"; return nil }

// SplitDNFItems splits DNF items.
// DNF means disjunctive normal form, e.g. "a or b or c".
func SplitDNFItems(onExpr Expression) []Expression { _ = "STUB: not implemented"; return nil }

// EvaluateExprWithNull sets columns in schema as null and calculate the final result of the scalar function.
// If the Expression is a non-constant value, it means the result is unknown.
func EvaluateExprWithNull(ctx context.Context, schema Schema, expr Expression) (Expression, error) {
	_ = "STUB: not implemented"
	return *new(Expression), nil
}

// TableInfo2Schema converts table info to schema.
func TableInfo2Schema(tbl *model.TableInfo) Schema { _ = "STUB: not implemented"; return *new(Schema) }

// NewCastFunc creates a new cast function.
func NewCastFunc(tp *types.FieldType, arg Expression, ctx context.Context) *ScalarFunction {
	_ = "STUB: not implemented"
	return nil
}

// NewValuesFunc creates a new values function.
func NewValuesFunc(offset int, retTp *types.FieldType, ctx context.Context) *ScalarFunction {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	expressionMySQLErrCodes := map[terror.ErrCode]uint16{
		codeIncorrectParameterCount: mysql.ErrWrongParamcountToNativeFct,
		codeFunctionNotExists:       mysql.ErrSpDoesNotExist,
	}
	terror.ErrClassToMySQLCodes[terror.ClassExpression] = expressionMySQLErrCodes
}
