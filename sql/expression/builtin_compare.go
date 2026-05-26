// Copyright 2017 PingCAP, Inc.
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
	_ functionClass = &coalesceFunctionClass{}
	_ functionClass = &greatestFunctionClass{}
	_ functionClass = &leastFunctionClass{}
	_ functionClass = &intervalFunctionClass{}
	_ functionClass = &compareFunctionClass{}
)

var (
	_ builtinFunc = &builtinCoalesceSig{}
	_ builtinFunc = &builtinGreatestSig{}
	_ builtinFunc = &builtinLeastSig{}
	_ builtinFunc = &builtinIntervalSig{}
	_ builtinFunc = &builtinCompareSig{}
)

type coalesceFunctionClass struct {
	baseFunctionClass
}

func (c *coalesceFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinCoalesceSig struct {
	baseBuiltinFunc
}

func (b *builtinCoalesceSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/comparison-operators.html#function_coalesce
func builtinCoalesce(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type greatestFunctionClass struct {
	baseFunctionClass
}

func (c *greatestFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinGreatestSig struct {
	baseBuiltinFunc
}

func (b *builtinGreatestSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/comparison-operators.html#function_greatest
func builtinGreatest(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type leastFunctionClass struct {
	baseFunctionClass
}

func (c *leastFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinLeastSig struct {
	baseBuiltinFunc
}

func (b *builtinLeastSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/comparison-operators.html#function_least
func builtinLeast(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type intervalFunctionClass struct {
	baseFunctionClass
}

func (c *intervalFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinIntervalSig struct {
	baseBuiltinFunc
}

func (b *builtinIntervalSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/comparison-operators.html#function_interval
func builtinInterval(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type compareFunctionClass struct {
	baseFunctionClass

	op opcode.Op
}

func (c *compareFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinCompareSig struct {
	baseBuiltinFunc

	op opcode.Op
}

func (b *builtinCompareSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func compareFuncFactory(op opcode.Op) BuiltinFunc {
	_ = "STUB: not implemented"
	return *new(BuiltinFunc)
}

// for <=>, if a and b are both nil, return true.
// if a or b is nil, return false.
