// Copyright 2014 The ql Authors. All rights reserved.
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

package types

import (
	"math"

	"github.com/chrislusf/gleam/sql/mysql"
	"github.com/chrislusf/gleam/sql/sessionctx/variable"
)

func truncateStr(str string, flen int) string { _ = "STUB: not implemented"; return "" }

var unsignedUpperBound = map[byte]uint64{
	mysql.TypeTiny:     math.MaxUint8,
	mysql.TypeShort:    math.MaxUint16,
	mysql.TypeInt24:    mysql.MaxUint24,
	mysql.TypeLong:     math.MaxUint32,
	mysql.TypeLonglong: math.MaxUint64,
	mysql.TypeBit:      math.MaxUint64,
	mysql.TypeEnum:     math.MaxUint64,
	mysql.TypeSet:      math.MaxUint64,
}

var signedUpperBound = map[byte]int64{
	mysql.TypeTiny:     math.MaxInt8,
	mysql.TypeShort:    math.MaxInt16,
	mysql.TypeInt24:    mysql.MaxInt24,
	mysql.TypeLong:     math.MaxInt32,
	mysql.TypeLonglong: math.MaxInt64,
}

var signedLowerBound = map[byte]int64{
	mysql.TypeTiny:     math.MinInt8,
	mysql.TypeShort:    math.MinInt16,
	mysql.TypeInt24:    mysql.MinInt24,
	mysql.TypeLong:     math.MinInt32,
	mysql.TypeLonglong: math.MinInt64,
}

func convertFloatToInt(sc *variable.StatementContext, fval float64, lowerBound, upperBound int64, tp byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func convertIntToInt(val int64, lowerBound int64, upperBound int64, tp byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func convertUintToInt(val uint64, upperBound int64, tp byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func convertIntToUint(val int64, upperBound uint64, tp byte) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func convertUintToUint(val uint64, upperBound uint64, tp byte) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func convertFloatToUint(sc *variable.StatementContext, fval float64, upperBound uint64, tp byte) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func isCastType(tp byte) bool { _ = "STUB: not implemented"; return false }

// StrToInt converts a string to an integer at the best-effort.
func StrToInt(sc *variable.StatementContext, str string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// StrToUint converts a string to an unsigned interger at the best-effortt.
func StrToUint(sc *variable.StatementContext, str string) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// getValidIntPrefix gets prefix of the string which can be successfully parsed as int.
func getValidIntPrefix(sc *variable.StatementContext, str string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// floatStrToIntStr converts a valid float string into valid integer string which can be parsed by
// strconv.ParseInt, we can't parse float first then convert it to string because precision will
// be lost.
func floatStrToIntStr(validFloat string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// (exp + incCnt) overflows MaxInt64.

// Return overflow to avoid allocating too much memory.

// StrToFloat converts a string to a float64 at the best-effort.
func StrToFloat(sc *variable.StatementContext, str string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// getValidFloatPrefix gets prefix of string which can be successfully parsed as float.
func getValidFloatPrefix(sc *variable.StatementContext, s string) (valid string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// "1e+1" is valid.

// "1.1." or "1e1.1"

// "123." is valid.

// "+.e"

// "1e5e"

// ToString converts an interface to a string.
func ToString(value interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }
