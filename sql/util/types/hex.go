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

// Hex is for mysql hexadecimal literal type.
type Hex struct {
	// Value holds numeric value for hexadecimal literal.
	Value int64
}

// String implements fmt.Stringer interface.
func (h Hex) String() string { _ = "STUB: not implemented"; return "" }

// ToNumber changes hexadecimal type to float64 for numeric operation.
// MySQL treats hexadecimal literal as double type.
func (h Hex) ToNumber() float64 { _ = "STUB: not implemented"; return 0 }

// ToString returns the string representation for hexadecimal literal.
func (h Hex) ToString() string { _ = "STUB: not implemented"; return "" }

// should never error.

func uniformHexStrLit(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// format is x'val' or X'val'

// here means format is not x'val', X'val' or 0xval.

// ParseHex parses hexadecimal literal string.
// The string format can be X'val', x'val' or 0xval.
// val must in (0...9, a...f, A...F).
func ParseHex(s string) (Hex, error) { _ = "STUB: not implemented"; return *new(Hex), nil }

// ParseHexStr parses hexadecimal literal as string.
// See https://dev.mysql.com/doc/refman/5.7/en/hexadecimal-literals.html
func ParseHexStr(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }
