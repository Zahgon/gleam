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
	"github.com/chrislusf/gleam/sql/util/types"
)

var (
	_ functionClass = &sleepFunctionClass{}
	_ functionClass = &inFunctionClass{}
	_ functionClass = &rowFunctionClass{}
	_ functionClass = &castFunctionClass{}
	_ functionClass = &setVarFunctionClass{}
	_ functionClass = &getVarFunctionClass{}
	_ functionClass = &lockFunctionClass{}
	_ functionClass = &releaseLockFunctionClass{}
	_ functionClass = &valuesFunctionClass{}
)

var (
	_ builtinFunc = &builtinSleepSig{}
	_ builtinFunc = &builtinInSig{}
	_ builtinFunc = &builtinRowSig{}
	_ builtinFunc = &builtinCastSig{}
	_ builtinFunc = &builtinSetVarSig{}
	_ builtinFunc = &builtinGetVarSig{}
	_ builtinFunc = &builtinLockSig{}
	_ builtinFunc = &builtinReleaseLockSig{}
	_ builtinFunc = &builtinValuesSig{}
)

type sleepFunctionClass struct {
	baseFunctionClass
}

func (c *sleepFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinSleepSig struct {
	baseBuiltinFunc
}

func (b *builtinSleepSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/miscellaneous-functions.html#function_sleep
func builtinSleep(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// processing argument is negative

// TODO: consider it's interrupted using KILL QUERY from other session, or
// interrupted by time out.

type inFunctionClass struct {
	baseFunctionClass
}

func (c *inFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinInSig struct {
	baseBuiltinFunc
}

func (b *builtinInSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/any-in-some-subqueries.html
func builtinIn(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// If it's no matched but we get null in In, returns null.
// e.g 1 in (null, 2, 3) returns null.

type rowFunctionClass struct {
	baseFunctionClass
}

func (c *rowFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinRowSig struct {
	baseBuiltinFunc
}

func (b *builtinRowSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func builtinRow(row []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type castFunctionClass struct {
	baseFunctionClass

	tp *types.FieldType
}

func (c *castFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinCastSig struct {
	baseBuiltinFunc

	tp *types.FieldType
}

func (b *builtinCastSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// CastFuncFactory produces builtin function according to field types.
// See https://dev.mysql.com/doc/refman/5.7/en/cast-functions.html
func CastFuncFactory(tp *types.FieldType) (BuiltinFunc, error) {
	_ = "STUB: not implemented"

	// Parser has restricted this.
	return *new(BuiltinFunc), nil
}

type setVarFunctionClass struct {
	baseFunctionClass
}

func (c *setVarFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinSetVarSig struct {
	baseBuiltinFunc
}

func (b *builtinSetVarSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func builtinSetVar(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type getVarFunctionClass struct {
	baseFunctionClass
}

func (c *getVarFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinGetVarSig struct {
	baseBuiltinFunc
}

func (b *builtinGetVarSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func builtinGetVar(args []types.Datum, ctx context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type lockFunctionClass struct {
	baseFunctionClass
}

func (c *lockFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinLockSig struct {
	baseBuiltinFunc
}

func (b *builtinLockSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// The lock function will do nothing.
// Warning: get_lock() function is parsed but ignored.
func builtinLock(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type releaseLockFunctionClass struct {
	baseFunctionClass
}

func (c *releaseLockFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinReleaseLockSig struct {
	baseBuiltinFunc
}

func (b *builtinReleaseLockSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// The release lock function will do nothing.
// Warning: release_lock() function is parsed but ignored.
func builtinReleaseLock(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type valuesFunctionClass struct {
	baseFunctionClass

	offset int
}

func (c *valuesFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinValuesSig struct {
	baseBuiltinFunc

	offset int
}

func (b *builtinValuesSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// BuiltinValuesFactory generates values builtin function.
func BuiltinValuesFactory(offset int) BuiltinFunc {
	_ = "STUB: not implemented"
	return *new(BuiltinFunc)
}
