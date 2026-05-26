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
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/model"
	"github.com/chrislusf/gleam/sql/util/types"
)

// ScalarFunction is the function that returns a value.
type ScalarFunction struct {
	FuncName model.CIStr
	// TODO: Implement type inference here, now we use ast's return type temporarily.
	RetType  *types.FieldType
	Function builtinFunc
}

// GetArgs gets arguments of function.
func (sf *ScalarFunction) GetArgs() []Expression { _ = "STUB: not implemented"; return nil }

// GetCtx gets the context of function.
func (sf *ScalarFunction) GetCtx() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// String implements fmt.Stringer interface.
func (sf *ScalarFunction) String() string { _ = "STUB: not implemented"; return "" }

// MarshalJSON implements json.Marshaler interface.
func (sf *ScalarFunction) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// NewFunction creates a new scalar function or constant.
func NewFunction(ctx context.Context, funcName string, retType *types.FieldType, args ...Expression) (Expression, error) {
	_ = "STUB: not implemented"
	return *new(Expression), nil
}

// ScalarFuncs2Exprs converts []*ScalarFunction to []Expression.
func ScalarFuncs2Exprs(funcs []*ScalarFunction) []Expression { _ = "STUB: not implemented"; return nil }

// Clone implements Expression interface.
func (sf *ScalarFunction) Clone() Expression { _ = "STUB: not implemented"; return *new(Expression) }

// GetType implements Expression interface.
func (sf *ScalarFunction) GetType() *types.FieldType {
	_ = "STUB: not implemented"

	// Equal implements Expression interface.
	return nil
}

func (sf *ScalarFunction) Equal(e Expression, ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// IsCorrelated implements Expression interface.
func (sf *ScalarFunction) IsCorrelated() bool { _ = "STUB: not implemented"; return false }

// Decorrelate implements Expression interface.
func (sf *ScalarFunction) Decorrelate(schema Schema) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// Eval implements Expression interface.
func (sf *ScalarFunction) Eval(row []types.Datum, _ context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// HashCode implements Expression interface.
func (sf *ScalarFunction) HashCode() []byte { _ = "STUB: not implemented"; return nil }

// ResolveIndices implements Expression interface.
func (sf *ScalarFunction) ResolveIndices(schema Schema) { _ = "STUB: not implemented"; return }
