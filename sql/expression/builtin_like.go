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

const (
	patMatch = iota + 1
	patOne
	patAny
)

var (
	_ functionClass = &likeFunctionClass{}
	_ functionClass = &regexpFunctionClass{}
)

var (
	_ builtinFunc = &builtinLikeSig{}
	_ builtinFunc = &builtinRegexpSig{}
)

// Handle escapes and wild cards convert pattern characters and pattern types.
func compilePattern(pattern string, escape byte) (patChars, patTypes []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// valid escape.

// invalid escape, fall back to escape byte
// mysql will treat escape character as the origin value even
// the escape sequence is invalid in Go or C.
// e.g, \m is invalid in Go, but in MySQL we will get "m" for select '\m'.
// Following case is correct just for escape \, not for others like +.
// TODO: add more checks for other escapes.

const caseDiff = 'a' - 'A'

func matchByteCI(a, b byte) bool { _ = "STUB: not implemented"; return false }

func doMatch(str string, patChars, patTypes []byte) bool { _ = "STUB: not implemented"; return false }

type likeFunctionClass struct {
	baseFunctionClass
}

func (c *likeFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinLikeSig struct {
	baseBuiltinFunc
}

func (b *builtinLikeSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/string-comparison-functions.html
func builtinLike(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// TODO: We don't need to compile pattern if it has been compiled or it is static.

type regexpFunctionClass struct {
	baseFunctionClass
}

func (c *regexpFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinRegexpSig struct {
	baseBuiltinFunc
}

func (b *builtinRegexpSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/regexp.html#operator_regexp
func builtinRegexp(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	// TODO: We don't need to compile pattern if it has been compiled or it is static.
	return *new(types.Datum), nil
}
