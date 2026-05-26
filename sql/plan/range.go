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

package plan

import (
	"github.com/chrislusf/gleam/sql/expression"
	"github.com/chrislusf/gleam/sql/sessionctx/variable"
	"github.com/chrislusf/gleam/sql/util/types"
)

type rangePoint struct {
	value types.Datum
	excl  bool // exclude
	start bool
}

func (rp rangePoint) String() string { _ = "STUB: not implemented"; return "" }

type rangePointSorter struct {
	points []rangePoint
	err    error
	sc     *variable.StatementContext
}

func (r *rangePointSorter) Len() int { _ = "STUB: not implemented"; return 0 }

func (r *rangePointSorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func rangePointLess(sc *variable.StatementContext, a, b rangePoint) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func rangePointEqualValueLess(a, b rangePoint) bool { _ = "STUB: not implemented"; return false }

func (r *rangePointSorter) Swap(i, j int) { _ = "STUB: not implemented"; return }

type rangeBuilder struct {
	err error
	sc  *variable.StatementContext
}

func (r *rangeBuilder) build(expr expression.Expression) []rangePoint {
	_ = "STUB: not implemented"
	return nil
}

func (r *rangeBuilder) buildFromConstant(expr *expression.Constant) []rangePoint {
	_ = "STUB: not implemented"
	return nil
}

func (r *rangeBuilder) buildFromColumn(expr *expression.Column) []rangePoint {
	_ = "STUB: not implemented"
	// column name expression is equivalent to column name is true.
	return nil
}

func (r *rangeBuilder) buildFormBinOp(expr *expression.ScalarFunction) []rangePoint {
	_ = "STUB: not implemented"
	// This has been checked that the binary operation is comparison operation, and one of
	// the operand is column name expression.
	return nil
}

func (r *rangeBuilder) buildFromIsTrue(expr *expression.ScalarFunction, isNot int) []rangePoint {
	_ = "STUB: not implemented"

	// NOT TRUE range is {[null null] [0, 0]}
	return nil
}

// TRUE range is {[-inf 0) (0 +inf]}

func (r *rangeBuilder) buildFromIsFalse(expr *expression.ScalarFunction, isNot int) []rangePoint {
	_ = "STUB: not implemented"

	// NOT FALSE range is {[-inf, 0), (0, +inf], [null, null]}
	return nil
}

// FALSE range is {[0, 0]}

func (r *rangeBuilder) newBuildFromIn(expr *expression.ScalarFunction) []rangePoint {
	_ = "STUB: not implemented"
	return nil
}

// check duplicates

// remove duplicates

func (r *rangeBuilder) newBuildFromPatternLike(expr *expression.ScalarFunction) []rangePoint {
	_ = "STUB: not implemented"
	return nil
}

// Get the prefix.

// Get the prefix, but exclude the prefix.
// e.g., "abc_x", the start point exclude "abc",
// because the string length is more than 3.

// Make the end point value more than the start point value,
// and the length of the end point value is the same as the length of the start point value.
// e.g., the start point value is "abc", so the end point value is "abd".

// If highValue[i] is 255 and highValue[i]++ is 0, then the end point value is max value.

func (r *rangeBuilder) buildFromNot(expr *expression.ScalarFunction) []rangePoint {
	_ = "STUB: not implemented"
	return nil
}

// Pattern not in is not supported.

// Pattern not like is not supported.

func (r *rangeBuilder) buildFromScalarFunc(expr *expression.ScalarFunction) []rangePoint {
	_ = "STUB: not implemented"
	return nil
}

func (r *rangeBuilder) intersection(a, b []rangePoint) []rangePoint {
	_ = "STUB: not implemented"
	return nil
}

func (r *rangeBuilder) union(a, b []rangePoint) []rangePoint { _ = "STUB: not implemented"; return nil }

func (r *rangeBuilder) merge(a, b []rangePoint, union bool) []rangePoint {
	_ = "STUB: not implemented"
	return nil
}

// just reached the required in range count, a new range started.

// just about to leave the required in range count, the range is ended.

// buildIndexRanges build index ranges from range points.
// Only the first column in the index is built, extra column ranges will be appended by
// appendIndexRanges.
func (r *rangeBuilder) buildIndexRanges(rangePoints []rangePoint, tp *types.FieldType) []*IndexRange {
	_ = "STUB: not implemented"
	return nil
}

func (r *rangeBuilder) convertPoint(point rangePoint, tp *types.FieldType) rangePoint {
	_ = "STUB: not implemented"
	return *new(rangePoint)
}

// e.g. "a > 1.9" convert to "a >= 2".

// e.g. "a >= 1.1 convert to "a > 1"

// e.g. "a < 1.1" convert to "a <= 1"

// e.g. "a <= 1.9" convert to "a < 2"

// appendIndexRanges appends additional column ranges for multi-column index.
// The additional column ranges can only be appended to point ranges.
// for example we have an index (a, b), if the condition is (a > 1 and b = 2)
// then we can not build a conjunctive ranges for this index.
func (r *rangeBuilder) appendIndexRanges(origin []*IndexRange, rangePoints []rangePoint, ft *types.FieldType) []*IndexRange {
	_ = "STUB: not implemented"
	return nil
}

func (r *rangeBuilder) appendIndexRange(origin *IndexRange, rangePoints []rangePoint, ft *types.FieldType) []*IndexRange {
	_ = "STUB: not implemented"
	return nil
}

func (r *rangeBuilder) buildTableRanges(rangePoints []rangePoint) []TableRange {
	_ = "STUB: not implemented"
	return nil
}
