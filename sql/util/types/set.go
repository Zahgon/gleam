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

var zeroSet = Set{Name: "", Value: 0}

// Set is for MySQL Set type.
type Set struct {
	Name  string
	Value uint64
}

// String implements fmt.Stringer interface.
func (e Set) String() string {
	_ = "STUB: not implemented"

	// ToNumber changes Set to float64 for numeric operation.
	return ""
}

func (e Set) ToNumber() float64 { _ = "STUB: not implemented"; return 0 }

// ParseSetName creates a Set with name.
func ParseSetName(elems []string, name string) (Set, error) {
	_ = "STUB: not implemented"
	return *new(Set), nil
}

// name doesn't exist, maybe an integer?

var (
	setIndexValue       []uint64
	setIndexInvertValue []uint64
)

func init() {
	setIndexValue = make([]uint64, 64)
	setIndexInvertValue = make([]uint64, 64)

	for i := 0; i < 64; i++ {
		setIndexValue[i] = 1 << uint64(i)
		setIndexInvertValue[i] = ^setIndexValue[i]
	}
}

// ParseSetValue creates a Set with special number.
func ParseSetValue(elems []string, number uint64) (Set, error) {
	_ = "STUB: not implemented"
	return *new(Set), nil
}
