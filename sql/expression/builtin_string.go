// Copyright 2013 The ql Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSES/QL-LICENSE file.

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
	_ functionClass = &lengthFunctionClass{}
	_ functionClass = &asciiFunctionClass{}
	_ functionClass = &concatFunctionClass{}
	_ functionClass = &concatWSFunctionClass{}
	_ functionClass = &leftFunctionClass{}
	_ functionClass = &repeatFunctionClass{}
	_ functionClass = &lowerFunctionClass{}
	_ functionClass = &reverseFunctionClass{}
	_ functionClass = &spaceFunctionClass{}
	_ functionClass = &upperFunctionClass{}
	_ functionClass = &strcmpFunctionClass{}
	_ functionClass = &replaceFunctionClass{}
	_ functionClass = &convertFunctionClass{}
	_ functionClass = &substringFunctionClass{}
	_ functionClass = &substringIndexFunctionClass{}
	_ functionClass = &locateFunctionClass{}
	_ functionClass = &hexFunctionClass{}
	_ functionClass = &unhexFunctionClass{}
	_ functionClass = &trimFunctionClass{}
	_ functionClass = &lTrimFunctionClass{}
	_ functionClass = &rTrimFunctionClass{}
	_ functionClass = &rpadFunctionClass{}
	_ functionClass = &bitLengthFunctionClass{}
	_ functionClass = &charFunctionClass{}
	_ functionClass = &charLengthFunctionClass{}
	_ functionClass = &findInSetFunctionClass{}
	_ functionClass = &fieldFunctionClass{}
)

var (
	_ builtinFunc = &builtinLengthSig{}
	_ builtinFunc = &builtinASCIISig{}
	_ builtinFunc = &builtinConcatSig{}
	_ builtinFunc = &builtinConcatWSSig{}
	_ builtinFunc = &builtinLeftSig{}
	_ builtinFunc = &builtinRepeatSig{}
	_ builtinFunc = &builtinLowerSig{}
	_ builtinFunc = &builtinReverseSig{}
	_ builtinFunc = &builtinSpaceSig{}
	_ builtinFunc = &builtinUpperSig{}
	_ builtinFunc = &builtinStrcmpSig{}
	_ builtinFunc = &builtinReplaceSig{}
	_ builtinFunc = &builtinConvertSig{}
	_ builtinFunc = &builtinSubstringSig{}
	_ builtinFunc = &builtinSubstringIndexSig{}
	_ builtinFunc = &builtinLocateSig{}
	_ builtinFunc = &builtinHexSig{}
	_ builtinFunc = &builtinUnHexSig{}
	_ builtinFunc = &builtinTrimSig{}
	_ builtinFunc = &builtinLTrimSig{}
	_ builtinFunc = &builtinRTrimSig{}
	_ builtinFunc = &builtinRpadSig{}
	_ builtinFunc = &builtinBitLengthSig{}
	_ builtinFunc = &builtinCharSig{}
	_ builtinFunc = &builtinCharLengthSig{}
	_ builtinFunc = &builtinFindInSetSig{}
)

type lengthFunctionClass struct {
	baseFunctionClass
}

func (c *lengthFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinLengthSig struct {
	baseBuiltinFunc
}

func (b *builtinLengthSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html
func builtinLength(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type asciiFunctionClass struct {
	baseFunctionClass
}

func (c *asciiFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinASCIISig struct {
	baseBuiltinFunc
}

func (b *builtinASCIISig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_ascii
func builtinASCII(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type concatFunctionClass struct {
	baseFunctionClass
}

func (c *concatFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinConcatSig struct {
	baseBuiltinFunc
}

func (b *builtinConcatSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_concat
func builtinConcat(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type concatWSFunctionClass struct {
	baseFunctionClass
}

func (c *concatWSFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinConcatWSSig struct {
	baseBuiltinFunc
}

func (b *builtinConcatWSSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_concat-ws
func builtinConcatWS(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type leftFunctionClass struct {
	baseFunctionClass
}

func (c *leftFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinLeftSig struct {
	baseBuiltinFunc
}

func (b *builtinLeftSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_left
func builtinLeft(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type repeatFunctionClass struct {
	baseFunctionClass
}

func (c *repeatFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinRepeatSig struct {
	baseBuiltinFunc
}

func (b *builtinRepeatSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_repeat
func builtinRepeat(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type lowerFunctionClass struct {
	baseFunctionClass
}

func (c *lowerFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinLowerSig struct {
	baseBuiltinFunc
}

func (b *builtinLowerSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_lower
func builtinLower(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type reverseFunctionClass struct {
	baseFunctionClass
}

func (c *reverseFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinReverseSig struct {
	baseBuiltinFunc
}

func (b *builtinReverseSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_reverse
func builtinReverse(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type spaceFunctionClass struct {
	baseFunctionClass
}

func (c *spaceFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinSpaceSig struct {
	baseBuiltinFunc
}

func (b *builtinSpaceSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_space
func builtinSpace(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type upperFunctionClass struct {
	baseFunctionClass
}

func (c *upperFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinUpperSig struct {
	baseBuiltinFunc
}

func (b *builtinUpperSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_upper
func builtinUpper(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type strcmpFunctionClass struct {
	baseFunctionClass
}

func (c *strcmpFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinStrcmpSig struct {
	baseBuiltinFunc
}

func (b *builtinStrcmpSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-comparison-functions.html
func builtinStrcmp(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type replaceFunctionClass struct {
	baseFunctionClass
}

func (c *replaceFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinReplaceSig struct {
	baseBuiltinFunc
}

func (b *builtinReplaceSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_replace
func builtinReplace(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type convertFunctionClass struct {
	baseFunctionClass
}

func (c *convertFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinConvertSig struct {
	baseBuiltinFunc
}

func (b *builtinConvertSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/cast-functions.html#function_convert
func builtinConvert(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	// Casting nil to any type returns nil
	return *new(types.Datum), nil
}

type substringFunctionClass struct {
	baseFunctionClass
}

func (c *substringFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinSubstringSig struct {
	baseBuiltinFunc
}

func (b *builtinSubstringSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func builtinSubstring(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	// The meaning of the elements of args.
	// arg[0] -> StrExpr
	// arg[1] -> Pos
	// arg[2] -> Len (Optional)
	return *new(types.Datum), nil
}

// The forms without a len argument return a substring from string str starting at position pos.
// The forms with a len argument return a substring len characters long from string str, starting at position pos.
// The forms that use FROM are standard SQL syntax. It is also possible to use a negative value for pos.
// In this case, the beginning of the substring is pos characters from the end of the string, rather than the beginning.
// A negative value may be used for pos in any of the forms of this function.

type substringIndexFunctionClass struct {
	baseFunctionClass
}

func (c *substringIndexFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinSubstringIndexSig struct {
	baseBuiltinFunc
}

func (b *builtinSubstringIndexSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_substring-index
func builtinSubstringIndex(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	// The meaning of the elements of args.
	// args[0] -> StrExpr
	// args[1] -> Delim
	// args[2] -> Count
	return *new(types.Datum), nil
}

// If count is positive, everything to the left of the final delimiter (counting from the left) is returned.

// If count is negative, everything to the right of the final delimiter (counting from the right) is returned.

type locateFunctionClass struct {
	baseFunctionClass
}

func (c *locateFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinLocateSig struct {
	baseBuiltinFunc
}

func (b *builtinLocateSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_locate
func builtinLocate(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	// The meaning of the elements of args.
	// args[0] -> SubStr
	// args[1] -> Str
	// args[2] -> Pos
	// eval str
	return *new(types.Datum), nil
}

// eval substr

// eval pos

const spaceChars = "\n\t\r "

type hexFunctionClass struct {
	baseFunctionClass
}

func (c *hexFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinHexSig struct {
	baseBuiltinFunc
}

func (b *builtinHexSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_hex
func builtinHex(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type unhexFunctionClass struct {
	baseFunctionClass
}

func (c *unhexFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinUnHexSig struct {
	baseBuiltinFunc
}

func (b *builtinUnHexSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_unhex
func builtinUnHex(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type trimFunctionClass struct {
	baseFunctionClass
}

func (c *trimFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinTrimSig struct {
	baseBuiltinFunc
}

func (b *builtinTrimSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_trim
func builtinTrim(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	// args[0] -> Str
	// args[1] -> RemStr
	// args[2] -> Direction
	// eval str
	return *new(types.Datum), nil
}

// eval remstr

// do trim

type lTrimFunctionClass struct {
	baseFunctionClass
}

func (c *lTrimFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinLTrimSig struct {
	baseBuiltinFunc
}

func (b *builtinLTrimSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type rTrimFunctionClass struct {
	baseFunctionClass
}

func (c *rTrimFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinRTrimSig struct {
	baseBuiltinFunc
}

func (b *builtinRTrimSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// For LTRIM & RTRIM
// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_ltrim
// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_rtrim
func trimFn(fn func(string, string) string, cutset string) BuiltinFunc {
	_ = "STUB: not implemented"
	return *new(BuiltinFunc)
}

func trimLeft(str, remstr string) string { _ = "STUB: not implemented"; return "" }

func trimRight(str, remstr string) string { _ = "STUB: not implemented"; return "" }

type rpadFunctionClass struct {
	baseFunctionClass
}

func (c *rpadFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinRpadSig struct {
	baseBuiltinFunc
}

func (b *builtinRpadSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_rpad
func builtinRpad(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	// RPAD(str,len,padstr)
	// args[0] string, args[1] int, args[2] string
	return *new(types.Datum), nil
}

type bitLengthFunctionClass struct {
	baseFunctionClass
}

func (c *bitLengthFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinBitLengthSig struct {
	baseBuiltinFunc
}

func (b *builtinBitLengthSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_bit-length
func builtinBitLength(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type charFunctionClass struct {
	baseFunctionClass
}

func (c *charFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinCharSig struct {
	baseBuiltinFunc
}

func (b *builtinCharSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_char
func builtinChar(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	// The kinds of args are int or string, and the last one represents charset.
	return *new(types.Datum), nil
}

// The last argument represents the charset name after "using".
// If it is nil, the default charset utf8 is used.

func convertInt64ToBytes(ints []int64) []byte { _ = "STUB: not implemented"; return nil }

func reverseByteSlice(slice []byte) { _ = "STUB: not implemented"; return }

type charLengthFunctionClass struct {
	baseFunctionClass
}

func (c *charLengthFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinCharLengthSig struct {
	baseBuiltinFunc
}

func (b *builtinCharLengthSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See https://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_char-length
func builtinCharLength(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

type findInSetFunctionClass struct {
	baseFunctionClass
}

func (c *findInSetFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinFindInSetSig struct {
	baseBuiltinFunc
}

func (b *builtinFindInSetSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_find-in-set
// TODO: This function can be optimized by using bit arithmetic when the first argument is
// a constant string and the second is a column of type SET.
func builtinFindInSet(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	// args[0] -> Str
	// args[1] -> StrList
	return *new(types.Datum), nil
}

type fieldFunctionClass struct {
	baseFunctionClass
}

func (c *fieldFunctionClass) getFunction(args []Expression, ctx context.Context) (builtinFunc, error) {
	_ = "STUB: not implemented"
	return *new(builtinFunc), nil
}

type builtinFieldSig struct {
	baseBuiltinFunc
}

func (b *builtinFieldSig) eval(row []types.Datum) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// See http://dev.mysql.com/doc/refman/5.7/en/string-functions.html#function_field
// Returns the index (position) of arg0 in the arg1, arg2, arg3, ... list.
// Returns 0 if arg0 is not found.
// If arg0 is NULL, the return value is 0 because NULL fails equality comparison with any value.
// If all arguments are strings, all arguments are compared as strings.
// If all arguments are numbers, they are compared as numbers.
// Otherwise, the arguments are compared as double.
func builtinField(args []types.Datum, ctx context.Context) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// argsToSpecifiedType converts the type of all arguments in args into string type or double type.
func argsToSpecifiedType(args []types.Datum, allString bool, allNumber bool, ctx context.Context) (newArgs []types.Datum, err error) {
	_ = "STUB: not implemented"
	// If all arguments are numbers, they can be compared directly without type converting.
	return nil, nil
}

// If error occurred when convert arg to float64, ignore it and set f as 0.
