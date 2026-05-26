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

// Bit is for mysql bit type.
type Bit struct {
	// Value holds the value for bit type.
	Value uint64

	// Width is the display with for bit value.
	// e.g, with is 8, 0 is for 0b00000000.
	Width int
}

// String implements fmt.Stringer interface.
func (b Bit) String() string { _ = "STUB: not implemented"; return "" }

// ToNumber changes bit type to float64 for numeric operation.
// MySQL treats bit as double type.
func (b Bit) ToNumber() float64 { _ = "STUB: not implemented"; return 0 }

// ToString returns the binary string for bit type.
func (b Bit) ToString() string { _ = "STUB: not implemented"; return "" }

// Min and Max bit width.
const (
	MinBitWidth = 1
	MaxBitWidth = 64
	// UnspecifiedBitWidth is the unspecified with if you want to calculate bit width dynamically.
	UnspecifiedBitWidth = -1
)

// ParseBit parses bit string.
// The string format can be b'val', B'val' or 0bval, val must be 0 or 1.
// Width is the display width for bit representation. -1 means calculating
// width dynamically, using following algorithm: (len("011101") + 7) & ^7,
// e.g, if bit string is 0b01, the above will return 8 for its bit width.
func ParseBit(s string, width int) (Bit, error) { _ = "STUB: not implemented"; return *new(Bit), nil }

// format is b'val' or B'val'

// here means format is not b'val', B'val' or 0bval.
