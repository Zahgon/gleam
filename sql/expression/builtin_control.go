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

package expression

import (
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/util/types"
)

var (
	_ functionClass = &caseWhenFunctionClass{}
	_ functionClass = &ifFunctionClass{}
	_ functionClass = &ifNullFunctionClass{}
	_ functionClass = &nullIfFunctionClass{}
)

var (
	_ builtinFunc = &builtinCaseWhenSig{}
	_ builtinFunc = &builtinIfSig{}
	_ builtinFunc = &builtinIfNullSig{}
	_ builtinFunc = &builtinNullIfSig{}
)

type caseWhenFunctionClass struct {
	baseFunctionClass
}

func (c *caseWhenFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinCaseWhenSig struct {
	baseBuiltinFunc
}

func (b *builtinCaseWhenSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/case.html
func builtinCaseWhen(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// when clause(condition, result) -> args[i], args[i+1]; (i >= 0 && i+1 < l-1)
// else clause -> args[l-1]
// If case clause has else clause, l%2 == 1.

type ifFunctionClass struct {
	baseFunctionClass
}

func (c *ifFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinIfSig struct {
	baseBuiltinFunc
}

func (b *builtinIfSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/control-flow-functions.html#function_if
func builtinIf(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	// if(expr1, expr2, expr3)
	// if expr1 is true, return expr2, otherwise, return expr3
	return *new(types.Datum), nil
}

// TODO: check return type, must be numeric or string

type ifNullFunctionClass struct {
	baseFunctionClass
}

func (c *ifNullFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinIfNullSig struct {
	baseBuiltinFunc
}

func (b *builtinIfNullSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/control-flow-functions.html#function_ifnull
func builtinIfNull(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	// ifnull(expr1, expr2)
	// if expr1 is not null, return expr1, otherwise, return expr2
	return *new(types.Datum), nil
}

type nullIfFunctionClass struct {
	baseFunctionClass
}

func (c *nullIfFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinNullIfSig struct {
	baseBuiltinFunc
}

func (b *builtinNullIfSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/control-flow-functions.html#function_nullif
func builtinNullIf(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	// nullif(expr1, expr2)
	// returns null if expr1 = expr2 is true, otherwise returns expr1
	return *new(types.Datum), nil
}
