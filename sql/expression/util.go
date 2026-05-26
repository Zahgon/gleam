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
	"github.com/chrislusf/gleam/sql/sessionctx/variable"
	"github.com/chrislusf/gleam/sql/util/types"
)

// ExtractColumns extracts all columns from an expression.
func ExtractColumns(expr Expression) (cols []*Column) { _ = "STUB: not implemented"; return nil }

// ColumnSubstitute substitutes the columns in filter to expressions in select fields.
// e.g. select * from (select b as a from t) k where a < 10 => select * from (select b as a from t where b < 10) k.
func ColumnSubstitute(expr Expression, schema Schema, newExprs []Expression) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

func datumsToConstants(datums []types.Datum) []Expression { _ = "STUB: not implemented"; return nil }

// calculateSum adds v to sum.
func calculateSum(sc *variable.StatementContext, sum, v types.Datum) (data types.Datum, err error) {
	_ = "STUB: not implemented"
	// for avg and sum calculation
	// avg and sum use decimal for integer and decimal type, use float for others
	// see https://dev.mysql.com/doc/refman/5.7/en/group-by-functions.html
	return *new(types.Datum), nil
}

// getValidPrefix gets a prefix of string which can parsed to a number with base. the minimun base is 2 and the maximum is 36.
func getValidPrefix(s string, base int64) string { _ = "STUB: not implemented"; return "" }
