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
	"github.com/chrislusf/gleam/sql/ast"
)

// KeyInfo stores the columns of one unique key or primary key.
type KeyInfo []*Column

// Clone copies the entire UniqueKey.
func (ki KeyInfo) Clone() KeyInfo { _ = "STUB: not implemented"; return *new(KeyInfo) }

// Schema stands for the row schema and unique key information get from input.
type Schema struct {
	Columns []*Column
	Keys    []KeyInfo
}

// String implements fmt.Stringer interface.
func (s Schema) String() string { _ = "STUB: not implemented"; return "" }

// Clone copies the total schema.
func (s Schema) Clone() Schema { _ = "STUB: not implemented"; return *new(Schema) }

// FindColumn finds an Column from schema for a ast.ColumnName. It compares the db/table/column names.
// If there are more than one result, it will raise ambiguous error.
func (s Schema) FindColumn(astCol *ast.ColumnName) (*Column, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RetrieveColumn retrieves column in expression from the columns in schema.
func (s Schema) RetrieveColumn(col *Column) *Column { _ = "STUB: not implemented"; return nil }

// GetColumnIndex finds the index for a column.
func (s Schema) GetColumnIndex(col *Column) int { _ = "STUB: not implemented"; return 0 }

// Len returns the number of columns in schema.
func (s Schema) Len() int { _ = "STUB: not implemented"; return 0 }

// Append append new column to the columns stored in schema.
func (s *Schema) Append(col *Column) { _ = "STUB: not implemented"; return }

// SetUniqueKeys will set the value of Schema.Keys.
func (s *Schema) SetUniqueKeys(keys []KeyInfo) {
	_ = "STUB: not implemented"

	// GetColumnsIndices will return a slice which contains the position of each column in schema.
	// If there is one column that doesn't match, nil will be returned.
	return
}

func (s Schema) GetColumnsIndices(cols []*Column) (ret []int) {
	_ = "STUB: not implemented"
	return nil
}

// MergeSchema will merge two schema into one schema.
func MergeSchema(lSchema, rSchema Schema) Schema { _ = "STUB: not implemented"; return *new(Schema) }

// NewSchema returns a schema made by its parameter.
func NewSchema(cols []*Column) Schema { _ = "STUB: not implemented"; return *new(Schema) }
