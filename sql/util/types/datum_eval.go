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

package types

import (
	"github.com/chrislusf/gleam/sql/sessionctx/variable"
)

// CoerceArithmetic converts datum to appropriate datum for arithmetic computing.
func CoerceArithmetic(sc *variable.StatementContext, a Datum) (d Datum, err error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// MySQL will convert string to float for arithmetic operation

// if time has no precision, return int64

// if duration has no precision, return int64

// ComputePlus computes the result of a+b.
func ComputePlus(a, b Datum) (d Datum, err error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// ComputeMinus computes the result of a-b.
func ComputeMinus(a, b Datum) (d Datum, err error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// ComputeMul computes the result of a*b.
func ComputeMul(a, b Datum) (d Datum, err error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// ComputeDiv computes the result of a/b.
func ComputeDiv(sc *variable.StatementContext, a, b Datum) (d Datum, err error) {
	_ = "STUB: not implemented"
	// MySQL support integer division Div and division operator /
	// we use opcode.Div for division operator and will use another for integer division later.
	// for division operator, we will use float64 for calculation.
	return *new(Datum), nil
}

// the scale of the result is the scale of the first operand plus
// the value of the div_precision_increment system variable (which is 4 by default)
// we will use 4 here

// division by zero return null

// ComputeMod computes the result of a mod b.
func ComputeMod(sc *variable.StatementContext, a, b Datum) (d Datum, err error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// first is int64, return int64.

// first is uint64, return uint64.

// div by zero returns nil without error.

// ComputeIntDiv computes the result of a / b, both a and b are integer.
func ComputeIntDiv(sc *variable.StatementContext, a, b Datum) (d Datum, err error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// If either is not integer, use decimal to calculate

// decimal2RoundUint converts a MyDecimal to an uint64 after rounding.
func decimal2RoundUint(x *MyDecimal) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// ComputeBitAnd computes the result of a & b.
func ComputeBitAnd(sc *variable.StatementContext, a, b Datum) (d Datum, err error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// If either is not integer, we round the operands and then use uint64 to calculate.

// ComputeBitOr computes the result of a | b.
func ComputeBitOr(sc *variable.StatementContext, a, b Datum) (d Datum, err error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// If either is not integer, we round the operands and then use uint64 to calculate.

// ComputeBitNeg computes the result of ~a.
func ComputeBitNeg(sc *variable.StatementContext, a Datum) (d Datum, err error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// If either is not integer, we round the operands and then use uint64 to calculate.

// ComputeBitXor computes the result of a ^ b.
func ComputeBitXor(sc *variable.StatementContext, a, b Datum) (d Datum, err error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// If either is not integer, we round the operands and then use uint64 to calculate.

// ComputeLeftShift computes the result of a >> b.
func ComputeLeftShift(sc *variable.StatementContext, a, b Datum) (d Datum, err error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// If either is not integer, we round the operands and then use uint64 to calculate.

// ComputeRightShift computes the result of a << b.
func ComputeRightShift(sc *variable.StatementContext, a, b Datum) (d Datum, err error) {
	_ = "STUB: not implemented"
	return *new(Datum), nil
}

// If either is not integer, we round the operands and then use uint64 to calculate.

// covertNonIntegerToUint64 coverts a non-integer to an uint64
func convertNonInt2RoundUint64(sc *variable.StatementContext, x Datum) (d uint64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
