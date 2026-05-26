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

package varsutil

import (
	"time"

	"github.com/chrislusf/gleam/sql/sessionctx/variable"
	"github.com/chrislusf/gleam/sql/util/types"
)

// GetSessionSystemVar gets a system variable.
// If it is a session only variable, use the default value defined in code.
// Returns error if there is no such variable.
func GetSessionSystemVar(s *variable.SessionVars, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// None-Global variable can use pre-defined default value.

// GetGlobalSystemVar gets a global system variable.
func GetGlobalSystemVar(s *variable.SessionVars, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// epochShiftBits is used to reserve logical part of the timestamp.
const epochShiftBits = 18

// SetSessionSystemVar sets system variable and updates SessionVars states.
func SetSessionSystemVar(vars *variable.SessionVars, name string, value types.Datum) error {
	_ = "STUB: not implemented"
	return nil
}

func parseTimeZone(s string) *time.Location {
	_ = "STUB: not implemented"

	// TODO: Support global time_zone variable, it should be set to global time_zone value.
	return nil
}

// The value can be given as a string indicating an offset from UTC, such as '+10:00' or '-6:00'.
