// Copyright 2016 The ql Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSES/QL-LICENSE file.

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

package table

import (
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/model"
	"github.com/chrislusf/gleam/sql/util/types"
)

// Column provides meta data describing a table column.
type Column model.ColumnInfo

// String implements fmt.Stringer interface.
func (c *Column) String() string { _ = "STUB: not implemented"; return "" }

// ToInfo casts Column to model.ColumnInfo
func (c *Column) ToInfo() *model.ColumnInfo { _ = "STUB: not implemented"; return nil }

// FindCol finds column in cols by name.
func FindCol(cols []*Column, name string) *Column { _ = "STUB: not implemented"; return nil }

// ToColumn converts a *model.ColumnInfo to *Column.
func ToColumn(col *model.ColumnInfo) *Column { _ = "STUB: not implemented"; return nil }

// FindCols finds columns in cols by names.
func FindCols(cols []*Column, names []string) ([]*Column, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindOnUpdateCols finds columns which have OnUpdateNow flag.
func FindOnUpdateCols(cols []*Column) []*Column { _ = "STUB: not implemented"; return nil }

// CastValues casts values based on columns type.
func CastValues(ctx context.Context, rec []types.Datum, cols []*Column, ignoreErr bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// CastValue casts a value based on column type.
func CastValue(ctx context.Context, val types.Datum, col *model.ColumnInfo) (casted types.Datum, err error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// TODO: add warnings.

// ColDesc describes column information like MySQL desc and show columns do.
type ColDesc struct {
	Field        string
	Type         string
	Collation    string
	Null         string
	Key          string
	DefaultValue interface{}
	Extra        string
	Privileges   string
	Comment      string
}

const defaultPrivileges string = "select,insert,update,references"

// GetTypeDesc gets the description for column type.
func (c *Column) GetTypeDesc() string { _ = "STUB: not implemented"; return "" }

// NewColDesc returns a new ColDesc for a column.
func NewColDesc(col *Column) *ColDesc {
	_ = "STUB: not implemented"
	// TODO: if we have no primary key and a unique index which's columns are all not null
	// we will set these columns' flag as PriKeyFlag
	// see https://dev.mysql.com/doc/refman/5.7/en/show-columns.html
	// create table
	return nil
}

// ColDescFieldNames returns the fields name in result set for desc and show columns.
func ColDescFieldNames(full bool) []string { _ = "STUB: not implemented"; return nil }

// CheckOnce checks if there are duplicated column names in cols.
func CheckOnce(cols []*Column) error { _ = "STUB: not implemented"; return nil }

// CheckNotNull checks if nil value set to a column with NotNull flag is set.
func (c *Column) CheckNotNull(data types.Datum) error { _ = "STUB: not implemented"; return nil }

// IsPKHandleColumn checks if the column is primary key handle column.
func (c *Column) IsPKHandleColumn(tbInfo *model.TableInfo) bool {
	_ = "STUB: not implemented"
	return false
}

// CheckNotNull checks if row has nil value set to a column with NotNull flag set.
func CheckNotNull(cols []*Column, row []types.Datum) error { _ = "STUB: not implemented"; return nil }

// GetColDefaultValue gets default value of the column.
func GetColDefaultValue(ctx context.Context, col *model.ColumnInfo) (types.Datum, bool, error) {
	_ = "STUB: not implemented"
	// Check no default value flag.
	return *new(types.Datum), false, nil
}

// TODO: add warning.

// Check and get timestamp/datetime default value.

// For enum type, if no default value and not null is set,
// the default value is the first element of the enum list

// GetZeroValue gets zero value for given column type.
func GetZeroValue(col *model.ColumnInfo) types.Datum {
	_ = "STUB: not implemented"
	return *new(types.Datum)
}
