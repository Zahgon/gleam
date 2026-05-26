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
	"time"

	"github.com/chrislusf/gleam/sql/ast"
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/util/types"
	"github.com/juju/errors"
)

const (
	zeroI64 int64 = 0
	oneI64  int64 = 1
)

func boolToInt64(v bool) int64 { _ = "STUB: not implemented"; return 0 }

var (
	// CurrentTimestamp is the keyword getting default value for datetime and timestamp type.
	CurrentTimestamp  = "CURRENT_TIMESTAMP"
	currentTimestampL = "current_timestamp"
	// ZeroTimestamp shows the zero datetime and timestamp.
	ZeroTimestamp = "0000-00-00 00:00:00"
)

var (
	errDefaultValue = errors.New("invalid default value")
)

// GetTimeValue gets the time value with type tp.
func GetTimeValue(ctx context.Context, v interface{}, tp byte, fsp int) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

func getTimeValue(ctx context.Context, v interface{}, tp byte, fsp int) (d types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// support some expression, like `-1`

// IsCurrentTimeExpr returns whether e is CurrentTimeExpr.
func IsCurrentTimeExpr(e ast.ExprNode) bool { _ = "STUB: not implemented"; return false }

func getSystemTimestamp(ctx context.Context) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// check whether use timestamp varibale
