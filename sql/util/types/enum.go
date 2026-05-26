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

// Enum is for MySQL enum type.
type Enum struct {
	Name  string
	Value uint64
}

// String implements fmt.Stringer interface.
func (e Enum) String() string {
	_ = "STUB: not implemented"

	// ToNumber changes enum index to float64 for numeric operation.
	return ""
}

func (e Enum) ToNumber() float64 { _ = "STUB: not implemented"; return 0 }

// ParseEnumName creates a Enum with item name.
func ParseEnumName(elems []string, name string) (Enum, error) {
	_ = "STUB: not implemented"
	return *new(Enum), nil
}

// name doesn't exist, maybe an integer?

// ParseEnumValue creates a Enum with special number.
func ParseEnumValue(elems []string, number uint64) (Enum, error) {
	_ = "STUB: not implemented"
	return *new(Enum), nil
}
