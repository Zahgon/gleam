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

package model

import (
	"github.com/chrislusf/gleam/sql/util/types"
)

// ColumnInfo provides meta data describing of a table column.
type ColumnInfo struct {
	Name            CIStr       `json:"name"`
	Offset          int         `json:"offset"`
	DefaultValue    interface{} `json:"default"`
	types.FieldType `json:"type"`
	Comment         string `json:"comment"`
}

// Clone clones ColumnInfo.
func (c *ColumnInfo) Clone() *ColumnInfo { _ = "STUB: not implemented"; return nil }

// TableInfo provides meta data describing a DB table.
type TableInfo struct {
	Name    CIStr         `json:"name"`
	Columns []*ColumnInfo `json:"cols"` // Columns are listed in the order in which they appear in the schema.
	Indices []*IndexInfo  `json:"index_info"`
	Comment string        `json:"comment"`
}

// Clone clones TableInfo.
func (t *TableInfo) Clone() *TableInfo { _ = "STUB: not implemented"; return nil }

// IndexColumn provides index column info.
type IndexColumn struct {
	Name CIStr `json:"name"` // Index name
}

// Clone clones IndexColumn.
func (i *IndexColumn) Clone() *IndexColumn { _ = "STUB: not implemented"; return nil }

// IndexType is the type of index
type IndexType int

// String implements Stringer interface.
func (t IndexType) String() string { _ = "STUB: not implemented"; return "" }

// IndexTypes
const (
	IndexTypeBtree IndexType = iota + 1
	IndexTypeHash
)

// IndexInfo provides meta data describing a DB index.
// It corresponds to the statement `CREATE INDEX Name ON Table (Column);`
// See https://dev.mysql.com/doc/refman/5.7/en/create-index.html
type IndexInfo struct {
	Name    CIStr          `json:"idx_name"`   // Index name.
	Table   CIStr          `json:"tbl_name"`   // Table name.
	Columns []*IndexColumn `json:"idx_cols"`   // Index columns.
	Unique  bool           `json:"is_unique"`  // Whether the index is unique.
	Primary bool           `json:"is_primary"` // Whether the index is primary key.
	Comment string         `json:"comment"`    // Comment
	Tp      IndexType      `json:"index_type"` // Index type: Btree or Hash
}

// Clone clones IndexInfo.
func (index *IndexInfo) Clone() *IndexInfo { _ = "STUB: not implemented"; return nil }

// DBInfo provides meta data describing a DB.
type DBInfo struct {
	Name   CIStr        `json:"db_name"` // DB name.
	Tables []*TableInfo `json:"-"`       // Tables in the DB.
}

// Clone clones DBInfo.
func (db *DBInfo) Clone() *DBInfo { _ = "STUB: not implemented"; return nil }

// CIStr is case insensitive string.
type CIStr struct {
	O string `json:"O"` // Original string.
	L string `json:"L"` // Lower case string.
}

// String implements fmt.Stringer interface.
func (cis CIStr) String() string {
	_ = "STUB: not implemented"

	// NewCIStr creates a new CIStr.
	return ""
}

func NewCIStr(s string) (cs CIStr) { _ = "STUB: not implemented"; return *new(CIStr) }
