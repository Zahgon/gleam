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
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/model"
	"github.com/chrislusf/gleam/sql/util/types"
)

// CorrelatedColumn stands for a column in a correlated sub query.
type CorrelatedColumn struct {
	Column

	Data *types.Datum
}

// Clone implements Expression interface.
func (col *CorrelatedColumn) Clone() Expression {
	_ = "STUB: not implemented"

	// Eval implements Expression interface.
	return *new(Expression)
}

func (col *CorrelatedColumn) Eval(row []types.Datum, _ context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *

	// Equal implements Expression interface.
	new(types.Datum), nil
}

func (col *CorrelatedColumn) Equal(expr Expression, ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// IsCorrelated implements Expression interface.
func (col *CorrelatedColumn) IsCorrelated() bool {
	_ = "STUB: not implemented"

	// Decorrelate implements Expression interface.
	return false
}

func (col *CorrelatedColumn) Decorrelate(schema Schema) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// ResolveIndices implements Expression interface.
func (col *CorrelatedColumn) ResolveIndices(_ Schema) {
	_ = "STUB: not implemented"

	// Column represents a column.
	return
}

type Column struct {
	FromID  string
	ColName model.CIStr
	DBName  model.CIStr
	TblName model.CIStr
	RetType *types.FieldType
	// Position means the position of this column that appears in the select fields.
	// e.g. SELECT name as id , 1 - id as id , 1 + name as id, name as id from src having id = 1;
	// There are four ids in the same schema, so you can't identify the column through the FromID and ColName.
	Position int
	// IsAggOrSubq means if this column is referenced to a Aggregation column or a Subquery column.
	// If so, this column's name will be the plain sql text.
	IsAggOrSubq bool

	// Only used for execution.
	Index int

	hashcode []byte
}

// Equal implements Expression interface.
func (col *Column) Equal(expr Expression, _ context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// String implements Stringer interface.
func (col *Column) String() string { _ = "STUB: not implemented"; return "" }

// MarshalJSON implements json.Marshaler interface.
func (col *Column) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GetType implements Expression interface.
func (col *Column) GetType() *types.FieldType {
	_ = "STUB: not implemented"

	// Eval implements Expression interface.
	return nil
}

func (col *Column) Eval(row []types.Datum, _ context.Context) (types.Datum, error) {
	_ = "STUB: not implemented"
	return *new(types.Datum), nil
}

// Clone implements Expression interface.
func (col *Column) Clone() Expression { _ = "STUB: not implemented"; return *new(Expression) }

// IsCorrelated implements Expression interface.
func (col *Column) IsCorrelated() bool {
	_ = "STUB: not implemented"

	// Decorrelate implements Expression interface.
	return false
}

func (col *Column) Decorrelate(_ Schema) Expression {
	_ = "STUB: not implemented"

	// HashCode implements Expression interface.
	return *new(Expression)
}

func (col *Column) HashCode() []byte { _ = "STUB: not implemented"; return nil }

// ResolveIndices implements Expression interface.
func (col *Column) ResolveIndices(schema Schema) { _ = "STUB: not implemented"; return }

// If col's index equals to -1, it means a internal logic error happens.

// Column2Exprs will transfer column slice to expression slice.
func Column2Exprs(cols []*Column) []Expression { _ = "STUB: not implemented"; return nil }
