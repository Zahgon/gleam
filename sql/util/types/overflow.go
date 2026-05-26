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
	"github.com/juju/errors"
)

// ErrArithOverflow is the error for arthimetic operation overflow.
var ErrArithOverflow = errors.New("operation overflow")

// AddUint64 adds uint64 a and b if no overflow, else returns error.
func AddUint64(a uint64, b uint64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// AddInt64 adds int64 a and b if no overflow, otherwise returns error.
func AddInt64(a int64, b int64) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// AddInteger adds uint64 a and int64 b and returns uint64 if no overflow error.
func AddInteger(a uint64, b int64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// SubUint64 subtracts uint64 a with b and returns uint64 if no overflow error.
func SubUint64(a uint64, b uint64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// SubInt64 subtracts int64 a with b and returns int64 if no overflow error.
func SubInt64(a int64, b int64) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// SubUintWithInt subtracts uint64 a with int64 b and returns uint64 if no overflow error.
func SubUintWithInt(a uint64, b int64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// SubIntWithUint subtracts int64 a with uint64 b and returns uint64 if no overflow error.
func SubIntWithUint(a int64, b uint64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// MulUint64 multiplies uint64 a and b and returns uint64 if no overflow error.
func MulUint64(a uint64, b uint64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// MulInt64 multiplies int64 a and b and returns int64 if no overflow error.
func MulInt64(a int64, b int64) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// negative result

// positive result

// MulInteger multiplies uint64 a and int64 b, and returns uint64 if no overflow error.
func MulInteger(a uint64, b int64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// DivInt64 divides int64 a with b, returns int64 if no overflow error.
// It just checks overflow, if b is zero, a "divide by zero" panic throws.
func DivInt64(a int64, b int64) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// DivUintWithInt divides uint64 a with int64 b, returns uint64 if no overflow error.
// It just checks overflow, if b is zero, a "divide by zero" panic throws.
func DivUintWithInt(a uint64, b int64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// DivIntWithUint divides int64 a with uint64 b, returns uint64 if no overflow error.
// It just checks overflow, if b is zero, a "divide by zero" panic throws.
func DivIntWithUint(a int64, b uint64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
