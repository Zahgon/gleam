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

package plan

import (
	"github.com/chrislusf/gleam/sql/ast"
	"github.com/chrislusf/gleam/sql/expression"
	"github.com/chrislusf/gleam/sql/model"
	"github.com/chrislusf/gleam/sql/util/types"
)

// JoinType contains CrossJoin, InnerJoin, LeftOuterJoin, RightOuterJoin, FullOuterJoin, SemiJoin.
type JoinType int

const (
	// InnerJoin means inner join.
	InnerJoin JoinType = iota
	// LeftOuterJoin means left join.
	LeftOuterJoin
	// RightOuterJoin means right join.
	RightOuterJoin
	// SemiJoin means if row a in table A matches some rows in B, just output a.
	SemiJoin
	// LeftOuterSemiJoin means if row a in table A matches some rows in B, output (a, true), otherwise, output (a, false).
	LeftOuterSemiJoin
)

// Join is the logical join plan.
type Join struct {
	baseLogicalPlan

	JoinType      JoinType
	anti          bool
	reordered     bool
	cartesianJoin bool

	EqualConditions []*expression.ScalarFunction
	LeftConditions  []expression.Expression
	RightConditions []expression.Expression
	OtherConditions []expression.Expression

	// DefaultValues is only used for outer join, which stands for the default values when the outer table cannot find join partner
	// instead of null padding.
	DefaultValues []types.Datum
}

func (p *Join) attachOnConds(onConds []expression.Expression) { _ = "STUB: not implemented"; return }

func (p *Join) extractCorrelatedCols() []*expression.CorrelatedColumn {
	_ = "STUB: not implemented"
	return nil
}

// SetCorrelated implements Plan interface.
func (p *Join) SetCorrelated() { _ = "STUB: not implemented"; return }

// Projection represents a select fields plan.
type Projection struct {
	baseLogicalPlan
	Exprs []expression.Expression
}

func (p *Projection) extractCorrelatedCols() []*expression.CorrelatedColumn {
	_ = "STUB: not implemented"
	return nil
}

// SetCorrelated implements Plan interface.
func (p *Projection) SetCorrelated() { _ = "STUB: not implemented"; return }

// Aggregation represents an aggregate plan.
type Aggregation struct {
	baseLogicalPlan

	AggFuncs     []expression.AggregationFunction
	GroupByItems []expression.Expression

	// groupByCols stores the columns that are group-by items.
	groupByCols []*expression.Column
}

func (p *Aggregation) extractCorrelatedCols() []*expression.CorrelatedColumn {
	_ = "STUB: not implemented"
	return nil
}

// SetCorrelated implements Plan interface.
func (p *Aggregation) SetCorrelated() { _ = "STUB: not implemented"; return }

// Selection means a filter.
type Selection struct {
	baseLogicalPlan

	// Originally the WHERE or ON condition is parsed into a single expression,
	// but after we converted to CNF(Conjunctive normal form), it can be
	// split into a list of AND conditions.
	Conditions []expression.Expression

	// onTable means if this selection's child is a table scan or index scan.
	onTable bool
}

func (p *Selection) extractCorrelatedCols() []*expression.CorrelatedColumn {
	_ = "STUB: not implemented"
	return nil
}

// SetCorrelated implements Plan interface.
func (p *Selection) SetCorrelated() { _ = "STUB: not implemented"; return }

// Apply gets one row from outer executor and gets one row from inner executor according to outer row.
type Apply struct {
	Join

	corCols []*expression.CorrelatedColumn
}

// SetCorrelated implements Plan interface.
func (p *Apply) SetCorrelated() { _ = "STUB: not implemented"; return }

// If the outer column can't be resolved from this outer schema, it should be resolved by outer schema.

// Exists checks if a query returns result.
type Exists struct {
	baseLogicalPlan
}

// MaxOneRow checks if a query returns no more than one row.
type MaxOneRow struct {
	baseLogicalPlan
}

// TableDual represents a dual table plan.
type TableDual struct {
	baseLogicalPlan
}

// DataSource represents a tablescan without condition push down.
type DataSource struct {
	baseLogicalPlan

	indexHints []*ast.IndexHint
	tableInfo  *model.TableInfo
	Columns    []*model.ColumnInfo
	DBName     *model.CIStr

	TableAsName *model.CIStr

	LimitCount *int64
}

// Trim trims extra columns in src rows.
type Trim struct {
	baseLogicalPlan
}

// Union represents Union plan.
type Union struct {
	baseLogicalPlan
}

// Sort stands for the order by plan.
type Sort struct {
	baseLogicalPlan

	ByItems   []*ByItems
	ExecLimit *Limit
}

func (p *Sort) extractCorrelatedCols() []*expression.CorrelatedColumn {
	_ = "STUB: not implemented"
	return nil
}

// SetCorrelated implements Plan interface.
func (p *Sort) SetCorrelated() { _ = "STUB: not implemented"; return }

// Update represents Update plan.
type Update struct {
	baseLogicalPlan

	OrderedList []*expression.Assignment
}

// Delete represents a delete plan.
type Delete struct {
	baseLogicalPlan

	Tables       []*ast.TableName
	IsMultiTable bool
}

// AddChild for parent.
func addChild(parent Plan, child Plan) { _ = "STUB: not implemented"; return }

// InsertPlan means inserting plan between two plans.
func InsertPlan(parent Plan, child Plan, insert Plan) error { _ = "STUB: not implemented"; return nil }

// RemovePlan means removing a plan.
func RemovePlan(p Plan) error { _ = "STUB: not implemented"; return nil }
