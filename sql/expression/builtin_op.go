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
	"github.com/chrislusf/gleam/sql/parser/opcode"
	"github.com/chrislusf/gleam/sql/util/types"
)

var (
	_ functionClass = &andandFunctionClass{}
	_ functionClass = &ororFunctionClass{}
	_ functionClass = &logicXorFunctionClass{}
	_ functionClass = &bitOpFunctionClass{}
	_ functionClass = &isTrueOpFunctionClass{}
	_ functionClass = &unaryOpFunctionClass{}
	_ functionClass = &isNullFunctionClass{}
)

var (
	_ builtinFunc = &builtinAndAndSig{}
	_ builtinFunc = &builtinOrOrSig{}
	_ builtinFunc = &builtinLogicXorSig{}
	_ builtinFunc = &builtinBitOpSig{}
	_ builtinFunc = &builtinIsTrueOpSig{}
	_ builtinFunc = &builtinUnaryOpSig{}
	_ builtinFunc = &builtinIsNullSig{}
)

type andandFunctionClass struct {
	baseFunctionClass
}

func (c *andandFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinAndAndSig struct {
	baseBuiltinFunc
}

func (b *builtinAndAndSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func builtinAndAnd(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// false && any other types is false

type ororFunctionClass struct {
	baseFunctionClass
}

func (c *ororFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinOrOrSig struct {
	baseBuiltinFunc
}

func (b *builtinOrOrSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func builtinOrOr(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// false && any other types is false

type logicXorFunctionClass struct {
	baseFunctionClass
}

func (c *logicXorFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinLogicXorSig struct {
	baseBuiltinFunc
}

func (b *builtinLogicXorSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func builtinLogicXor(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type bitOpFunctionClass struct {
	baseFunctionClass

	op opcode.Op
}

func (c *bitOpFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinBitOpSig struct {
	baseBuiltinFunc

	op opcode.Op
}

func (b *builtinBitOpSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func bitOpFactory(op opcode.Op) BuiltinFunc { _ = "STUB: not implemented"; return *new(BuiltinFunc) }

// use a int64 for bit operator, return uint64

type isTrueOpFunctionClass struct {
	baseFunctionClass

	op opcode.Op
}

func (c *isTrueOpFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinIsTrueOpSig struct {
	baseBuiltinFunc

	op opcode.Op
}

func (b *builtinIsTrueOpSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func isTrueOpFactory(op opcode.Op) BuiltinFunc { _ = "STUB: not implemented"; return *new(BuiltinFunc) }

type unaryOpFunctionClass struct {
	baseFunctionClass

	op opcode.Op
}

func (c *unaryOpFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinUnaryOpSig struct {
	baseBuiltinFunc

	op opcode.Op
}

func (b *builtinUnaryOpSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func unaryOpFactory(op opcode.Op) BuiltinFunc { _ = "STUB: not implemented"; return *new(BuiltinFunc) }

// for bit operation, we will use int64 first, then return uint64

type isNullFunctionClass struct {
	baseFunctionClass
}

func (c *isNullFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinIsNullSig struct {
	baseBuiltinFunc
}

func (b *builtinIsNullSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/comparison-operators.html#function_isnull
func builtinIsNull(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}
