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
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/expression"
	"github.com/chrislusf/gleam/sql/model"
	"github.com/chrislusf/gleam/sql/util/types"
	"github.com/pingcap/tipb/go-tipb"
)

var (
	_ physicalDistSQLPlan = &PhysicalTableScan{}
	_ physicalDistSQLPlan = &PhysicalIndexScan{}
)

// PhysicalIndexScan represents an index scan plan.
type PhysicalIndexScan struct {
	physicalTableSource

	Table      *model.TableInfo
	Index      *model.IndexInfo
	Ranges     []*IndexRange
	Columns    []*model.ColumnInfo
	DBName     *model.CIStr
	Desc       bool
	OutOfOrder bool
	// DoubleRead means if the index executor will read kv two times.
	// If the query requires the columns that don't belong to index, DoubleRead will be true.
	DoubleRead bool

	// All conditions in AccessCondition[accessEqualCount:accessInAndEqCount] are IN expressions or equal conditions.
	accessInAndEqCount int
	// All conditions in AccessCondition[:accessEqualCount] are equal conditions.
	accessEqualCount int

	TableAsName *model.CIStr
}

// PhysicalMemTable reads memory table.
type PhysicalMemTable struct {
	basePlan

	DBName      *model.CIStr
	Table       *model.TableInfo
	Columns     []*model.ColumnInfo
	Ranges      []TableRange
	TableAsName *model.CIStr
}

// Copy implements the PhysicalPlan Copy interface.
func (p *PhysicalMemTable) Copy() PhysicalPlan {
	_ = "STUB: not implemented"

	// physicalDistSQLPlan means the plan that can be executed distributively.
	// We can push down other plan like selection, limit, aggregation, topn into this plan.
	return *new(PhysicalPlan)
}

type physicalDistSQLPlan interface {
	addAggregation(ctx context.Context, agg *PhysicalAggregation) expression.Schema
	addTopN(ctx context.Context, prop *requiredProperty) bool
	addLimit(limit *Limit)
	// scanCount means the original row count that need to be scanned and resultCount means the row count after scanning.
	calculateCost(resultCount uint64, scanCount uint64) float64
}

func (p *PhysicalIndexScan) calculateCost(resultCount uint64, scanCount uint64) float64 {
	_ = "STUB: not implemented"
	// TODO: Eliminate index cost more precisely.
	return 0
}

// sort cost

func (p *PhysicalTableScan) calculateCost(resultCount uint64, scanCount uint64) float64 {
	_ = "STUB: not implemented"
	return 0
}

type physicalTableSource struct {
	basePlan

	Aggregated bool
	readOnly   bool
	AggFields  []*types.FieldType
	AggFuncsPB []*tipb.Expr
	GbyItemsPB []*tipb.ByItem

	// TableConditionPBExpr is the pb structure of conditions that used in the table scan.
	TableConditionPBExpr *tipb.Expr
	// IndexConditionPBExpr is the pb structure of conditions that used in the index scan.
	IndexConditionPBExpr *tipb.Expr

	// AccessCondition is used to calculate range.
	AccessCondition []expression.Expression

	LimitCount  *int64
	SortItemsPB []*tipb.ByItem

	// The following fields are used for explaining and testing. Because pb structures are not human-readable.
	aggFuncs              []expression.AggregationFunction
	gbyItems              []expression.Expression
	sortItems             []*ByItems
	indexFilterConditions []expression.Expression
	tableFilterConditions []expression.Expression
}

// MarshalJSON implements json.Marshaler interface.
func (p *physicalTableSource) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// print condition infos

func (p *physicalTableSource) clearForAggPushDown() { _ = "STUB: not implemented"; return }

func (p *physicalTableSource) clearForTopnPushDown() { _ = "STUB: not implemented"; return }

func needCount(af expression.AggregationFunction) bool { _ = "STUB: not implemented"; return false }

func needValue(af expression.AggregationFunction) bool { _ = "STUB: not implemented"; return false }

func (p *physicalTableSource) tryToAddUnionScan(resultPlan PhysicalPlan) PhysicalPlan {
	_ = "STUB: not implemented"
	return *new(PhysicalPlan)
}

func (p *physicalTableSource) addLimit(l *Limit) { _ = "STUB: not implemented"; return }

func (p *physicalTableSource) addTopN(ctx context.Context, prop *requiredProperty) bool {
	_ = "STUB: not implemented"
	return false
}

// sc := ctx.GetSessionVars().StmtCtx

// sortByItemToPB(sc, p.client, prop.col, prop.desc)

// When we fail to convert any sortItem to PB struct, we should clear the environments.

func (p *physicalTableSource) addAggregation(ctx context.Context, agg *PhysicalAggregation) expression.Schema {
	_ = "STUB: not implemented"
	return *new(expression.Schema)
}

// sc := ctx.GetSessionVars().StmtCtx

// aggFuncToPBExpr(sc, p.client, f)

// When we fail to convert any agg function to PB struct, we should clear the environments.

// groupByItemToPB(sc, p.client, item)

// When we fail to convert any group-by item to PB struct, we should clear the environments.

// PhysicalTableScan represents a table scan plan.
type PhysicalTableScan struct {
	physicalTableSource

	Table   *model.TableInfo
	Columns []*model.ColumnInfo
	DBName  *model.CIStr
	Desc    bool
	Ranges  []TableRange
	pkCol   *expression.Column

	TableAsName *model.CIStr

	// If sort data by scanning pkcol, KeepOrder should be true.
	KeepOrder bool
}

// PhysicalDummyScan is a dummy table that returns nothing.
type PhysicalDummyScan struct {
	basePlan
}

// PhysicalApply represents apply plan, only used for subquery.
type PhysicalApply struct {
	basePlan

	PhysicalJoin PhysicalPlan
	OuterSchema  []*expression.CorrelatedColumn
}

// PhysicalHashJoin represents hash join for inner/ outer join.
type PhysicalHashJoin struct {
	basePlan

	JoinType JoinType

	EqualConditions []*expression.ScalarFunction
	LeftConditions  []expression.Expression
	RightConditions []expression.Expression
	OtherConditions []expression.Expression
	SmallTable      int
	Concurrency     int

	DefaultValues []types.Datum
}

// PhysicalHashSemiJoin represents hash join for semi join.
type PhysicalHashSemiJoin struct {
	basePlan

	WithAux bool
	Anti    bool

	EqualConditions []*expression.ScalarFunction
	LeftConditions  []expression.Expression
	RightConditions []expression.Expression
	OtherConditions []expression.Expression
}

// AggregationType stands for the mode of aggregation plan.
type AggregationType int

const (
	// StreamedAgg supposes its input is sorted by group by key.
	StreamedAgg AggregationType = iota
	// FinalAgg supposes its input is partial results.
	FinalAgg
	// CompleteAgg supposes its input is original results.
	CompleteAgg
)

// PhysicalAggregation is Aggregation's physical plan.
type PhysicalAggregation struct {
	basePlan

	HasGby       bool
	AggType      AggregationType
	AggFuncs     []expression.AggregationFunction
	GroupByItems []expression.Expression
}

// PhysicalUnionScan represents a union scan operator.
type PhysicalUnionScan struct {
	basePlan

	Condition expression.Expression
}

// Cache plan is a physical plan which stores the result of its child node.
type Cache struct {
	basePlan
}

func (p *PhysicalHashJoin) extractCorrelatedCols() []*expression.CorrelatedColumn {
	_ = "STUB: not implemented"
	return nil
}

func (p *PhysicalHashSemiJoin) extractCorrelatedCols() []*expression.CorrelatedColumn {
	_ = "STUB: not implemented"
	return nil
}

func (p *PhysicalApply) extractCorrelatedCols() []*expression.CorrelatedColumn {
	_ = "STUB: not implemented"
	return nil
}

func (p *PhysicalAggregation) extractCorrelatedCols() []*expression.CorrelatedColumn {
	_ = "STUB: not implemented"
	return nil
}

// Copy implements the PhysicalPlan Copy interface.
func (p *PhysicalIndexScan) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// MarshalJSON implements json.Marshaler interface.
	new(PhysicalPlan)
}

func (p *PhysicalIndexScan) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Copy implements the PhysicalPlan Copy interface.
func (p *PhysicalTableScan) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// MarshalJSON implements json.Marshaler interface.
	new(PhysicalPlan)
}

func (p *PhysicalTableScan) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Copy implements the PhysicalPlan Copy interface.
func (p *PhysicalApply) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// MarshalJSON implements json.Marshaler interface.
	new(PhysicalPlan)
}

func (p *PhysicalApply) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetCorrelated implements Plan interface.
func (p *PhysicalApply) SetCorrelated() { _ = "STUB: not implemented"; return }

// Copy implements the PhysicalPlan Copy interface.
func (p *PhysicalHashSemiJoin) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// MarshalJSON implements json.Marshaler interface.
	new(PhysicalPlan)
}

func (p *PhysicalHashSemiJoin) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetCorrelated implements Plan interface.
func (p *PhysicalHashSemiJoin) SetCorrelated() { _ = "STUB: not implemented"; return }

// Copy implements the PhysicalPlan Copy interface.
func (p *PhysicalHashJoin) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// MarshalJSON implements json.Marshaler interface.
	new(PhysicalPlan)
}

func (p *PhysicalHashJoin) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetCorrelated implements Plan interface.
func (p *PhysicalHashJoin) SetCorrelated() { _ = "STUB: not implemented"; return }

// Copy implements the PhysicalPlan Copy interface.
func (p *Selection) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// MarshalJSON implements json.Marshaler interface.
	new(PhysicalPlan)
}

func (p *Selection) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Copy implements the PhysicalPlan Copy interface.
func (p *Projection) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// MarshalJSON implements json.Marshaler interface.
	new(PhysicalPlan)
}

func (p *Projection) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Copy implements the PhysicalPlan Copy interface.
func (p *Exists) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// Copy implements the PhysicalPlan Copy interface.
	new(PhysicalPlan)
}

func (p *MaxOneRow) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// Copy implements the PhysicalPlan Copy interface.
	new(PhysicalPlan)
}

func (p *Insert) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// Copy implements the PhysicalPlan Copy interface.
	new(PhysicalPlan)
}

func (p *Limit) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// MarshalJSON implements json.Marshaler interface.
	new(PhysicalPlan)
}

func (p *Limit) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Copy implements the PhysicalPlan Copy interface.
func (p *Union) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// Copy implements the PhysicalPlan Copy interface.
	new(PhysicalPlan)
}

func (p *Sort) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// MarshalJSON implements json.Marshaler interface.
	new(PhysicalPlan)
}

func (p *Sort) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Copy implements the PhysicalPlan Copy interface.
func (p *TableDual) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// Copy implements the PhysicalPlan Copy interface.
	new(PhysicalPlan)
}

func (p *Trim) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// Copy implements the PhysicalPlan Copy interface.
	new(PhysicalPlan)
}

func (p *SelectLock) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// Copy implements the PhysicalPlan Copy interface.
	new(PhysicalPlan)
}

func (p *PhysicalAggregation) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// MarshalJSON implements json.Marshaler interface.
	new(PhysicalPlan)
}

func (p *PhysicalAggregation) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetCorrelated implements Plan interface.
func (p *PhysicalAggregation) SetCorrelated() { _ = "STUB: not implemented"; return }

// Copy implements the PhysicalPlan Copy interface.
func (p *Update) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// Copy implements the PhysicalPlan Copy interface.
	new(PhysicalPlan)
}

func (p *PhysicalDummyScan) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// Copy implements the PhysicalPlan Copy interface.
	new(PhysicalPlan)
}

func (p *Delete) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// Copy implements the PhysicalPlan Copy interface.
	new(PhysicalPlan)
}

func (p *Show) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// Copy implements the PhysicalPlan Copy interface.
	new(PhysicalPlan)
}

func (p *PhysicalUnionScan) Copy() PhysicalPlan {
	_ = "STUB: not implemented"
	return *

	// Copy implements the PhysicalPlan Copy interface.
	new(PhysicalPlan)
}

func (p *Cache) Copy() PhysicalPlan { _ = "STUB: not implemented"; return *new(PhysicalPlan) }
