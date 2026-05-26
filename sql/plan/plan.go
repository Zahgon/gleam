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

package plan

import (
	"encoding/json"

	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/expression"
)

const (
	// Sel is the type of Selection.
	Sel = "Selection"
	// St is the type of Set.
	St = "Set"
	// Proj is the type of Projection.
	Proj = "Projection"
	// Agg is the type of Aggregation.
	Agg = "Aggregation"
	// Jn is the type of Join.
	Jn = "Join"
	// Un is the type of Union.
	Un = "Union"
	// Tbl is the type of TableScan.
	Tbl = "TableScan"
	// Idx is the type of IndexScan.
	Idx = "IndexScan"
	// Srt is the type of Sort.
	Srt = "Sort"
	// Lim is the type of Limit.
	Lim = "Limit"
	// App is the type of Apply.
	App = "Apply"
	// Dis is the type of Distinct.
	Dis = "Distinct"
	// Trm is the type of Trim.
	Trm = "Trim"
	// MOR is the type of MaxOneRow.
	MOR = "MaxOneRow"
	// Ext is the type of Exists.
	Ext = "Exists"
	// Dual is the type of TableDual.
	Dual = "TableDual"
	// Lock is the type of SelectLock.
	Lock = "SelectLock"
	// Load is the type of LoadData.
	Load = "LoadData"
	// Ins is the type of Insert
	Ins = "Insert"
	// Up is the type of Update.
	Up = "Update"
	// Del is the type of Delete.
	Del = "Delete"
)

// Plan is the description of an execution flow.
// It is created from ast.Node first, then optimized by the optimizer,
// finally used by the executor to create a Cursor which executes the statement.
type Plan interface {
	// AddParent means appending a parent for plan.
	AddParent(parent Plan)
	// AddChild means appending a child for plan.
	AddChild(children Plan)
	// ReplaceParent means replacing a parent with another one.
	ReplaceParent(parent, newPar Plan) error
	// ReplaceChild means replacing a child with another one.
	ReplaceChild(children, newChild Plan) error
	// Retrieve the parent by index.
	GetParentByIndex(index int) Plan
	// Retrieve the child by index.
	GetChildByIndex(index int) Plan
	// Get all the parents.
	GetParents() []Plan
	// Get all the children.
	GetChildren() []Plan
	// Set the schema.
	SetSchema(schema expression.Schema)
	// Get the schema.
	GetSchema() expression.Schema
	// Get the ID.
	GetID() string
	// Check whether this plan is correlated or not.
	IsCorrelated() bool
	// Set the value of attribute "correlated".
	// A plan will be correlated if one of its expressions or its child plans is correlated, except Apply.
	// As for Apply, it will be correlated if the outer plan is correlated or the inner plan has column that the outer doesn't has.
	// It will be called in the final step of logical plan building and the PhysicalInitialize process after convert2PhysicalPlan process.
	SetCorrelated()
	// SetParents sets the parents for the plan.
	SetParents(...Plan)
	// SetParents sets the children for the plan.
	SetChildren(...Plan)

	context() context.Context

	extractCorrelatedCols() []*expression.CorrelatedColumn
}

type columnProp struct {
	col  *expression.Column
	desc bool
}

func (c *columnProp) equal(nc *columnProp, ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

type requiredProperty struct {
	props      []*columnProp
	sortKeyLen int
	limit      *Limit
}

// getHashKey encodes a requiredProperty to a unique hash code.
func (p *requiredProperty) getHashKey() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type physicalPlanInfo struct {
	p     PhysicalPlan
	cost  float64
	count uint64
}

// LogicalPlan is a tree of logical operators.
// We can do a lot of logical optimizations to it, like predicate pushdown and column pruning.
type LogicalPlan interface {
	Plan

	// PredicatePushDown pushes down the predicates in the where/on/having clauses as deeply as possible.
	// It will accept a predicate that is an expression slice, and return the expressions that can't be pushed.
	// Because it might change the root if the having clause exists, we need to return a plan that represents a new root.
	PredicatePushDown([]expression.Expression) ([]expression.Expression, LogicalPlan, error)

	// PruneColumns prunes the unused columns.
	PruneColumns([]*expression.Column)

	// ResolveIndicesAndCorCols resolves the index for columns and initializes the correlated columns.
	ResolveIndicesAndCorCols()

	// convert2PhysicalPlan converts the logical plan to the physical plan.
	// It is called recursively from the parent to the children to create the result physical plan.
	// Some logical plans will convert the children to the physical plans in different ways, and return the one
	// with the lowest cost.
	convert2PhysicalPlan(prop *requiredProperty) (*physicalPlanInfo, error)

	// buildKeyInfo will collect the information of unique keys into schema.
	buildKeyInfo()
}

// PhysicalPlan is a tree of the physical operators.
type PhysicalPlan interface {
	json.Marshaler
	Plan

	// matchProperty calculates the cost of the physical plan if it matches the required property.
	// It's usually called at the end of convert2PhysicalPlan. Some physical plans do not implement it because there is
	// no property to match, these plans just do the cost calculation directly.
	// If the cost of the physical plan does not match the required property, the cost will be set to MaxInt64
	// so it will not be chosen as the result physical plan.
	// childrenPlanInfo are used to calculate the result cost of the plan.
	// The returned *physicalPlanInfo will be chosen as the final plan if it has the lowest cost.
	// For the lowest level *PhysicalTableScan and *PhysicalIndexScan, even though it doesn't have childPlanInfo, we
	// create an initial *physicalPlanInfo to pass the row count.
	matchProperty(prop *requiredProperty, childPlanInfo ...*physicalPlanInfo) *physicalPlanInfo

	// Copy copies the current plan.
	Copy() PhysicalPlan
}

type baseLogicalPlan struct {
	basePlan
	planMap map[string]*physicalPlanInfo
	self    LogicalPlan
}

func (p *baseLogicalPlan) getPlanInfo(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *baseLogicalPlan) convert2PhysicalPlan(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *baseLogicalPlan) storePlanInfo(prop *requiredProperty, info *physicalPlanInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// copy it

func (p *baseLogicalPlan) buildKeyInfo() { _ = "STUB: not implemented"; return }

func newBaseLogicalPlan(tp string, a *idAllocator) baseLogicalPlan {
	_ = "STUB: not implemented"
	return *new(baseLogicalPlan)
}

// PredicatePushDown implements LogicalPlan interface.
func (p *baseLogicalPlan) PredicatePushDown(predicates []expression.Expression) ([]expression.Expression, LogicalPlan, error) {
	_ = "STUB: not implemented"
	return nil, *new(LogicalPlan), nil
}

func (p *basePlan) extractCorrelatedCols() []*expression.CorrelatedColumn {
	_ = "STUB: not implemented"
	return nil
}

// ResolveIndicesAndCorCols implements LogicalPlan interface.
func (p *baseLogicalPlan) ResolveIndicesAndCorCols() { _ = "STUB: not implemented"; return }

// PruneColumns implements LogicalPlan interface.
func (p *baseLogicalPlan) PruneColumns(parentUsedCols []*expression.Column) {
	_ = "STUB: not implemented"
	return
}

func (p *basePlan) initIDAndContext(ctx context.Context) { _ = "STUB: not implemented"; return }

// basePlan implements base Plan interface.
// Should be used as embedded struct in Plan implementations.
type basePlan struct {
	correlated bool

	parents  []Plan
	children []Plan

	schema    expression.Schema
	tp        string
	id        string
	allocator *idAllocator
	ctx       context.Context
}

// MarshalJSON implements json.Marshaler interface.
func (p *basePlan) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// IsCorrelated implements Plan IsCorrelated interface.
func (p *basePlan) IsCorrelated() bool { _ = "STUB: not implemented"; return false }

func (p *basePlan) SetCorrelated() { _ = "STUB: not implemented"; return }

// GetID implements Plan GetID interface.
func (p *basePlan) GetID() string {
	_ = "STUB: not implemented"

	// SetSchema implements Plan SetSchema interface.
	return ""
}

func (p *basePlan) SetSchema(schema expression.Schema) {
	_ = "STUB: not implemented"

	// GetSchema implements Plan GetSchema interface.
	return
}

func (p *basePlan) GetSchema() expression.Schema {
	_ = "STUB: not implemented"

	// AddParent implements Plan AddParent interface.
	return *new(expression.Schema)
}

func (p *basePlan) AddParent(parent Plan) { _ = "STUB: not implemented"; return }

// AddChild implements Plan AddChild interface.
func (p *basePlan) AddChild(child Plan) { _ = "STUB: not implemented"; return }

// ReplaceParent means replace a parent for another one.
func (p *basePlan) ReplaceParent(parent, newPar Plan) error { _ = "STUB: not implemented"; return nil }

// ReplaceChild means replace a child with another one.
func (p *basePlan) ReplaceChild(child, newChild Plan) error { _ = "STUB: not implemented"; return nil }

// GetParentByIndex implements Plan GetParentByIndex interface.
func (p *basePlan) GetParentByIndex(index int) (parent Plan) {
	_ = "STUB: not implemented"
	return *new(Plan)
}

// GetChildByIndex implements Plan GetChildByIndex interface.
func (p *basePlan) GetChildByIndex(index int) (parent Plan) {
	_ = "STUB: not implemented"
	return *new(Plan)
}

// GetParents implements Plan GetParents interface.
func (p *basePlan) GetParents() []Plan {
	_ = "STUB: not implemented"

	// GetChildren implements Plan GetChildren interface.
	return nil
}

func (p *basePlan) GetChildren() []Plan {
	_ = "STUB: not implemented"

	// RemoveAllParents implements Plan RemoveAllParents interface.
	return nil
}

func (p *basePlan) SetParents(pars ...Plan) {
	_ = "STUB: not implemented"

	// RemoveAllParents implements Plan RemoveAllParents interface.
	return
}

func (p *basePlan) SetChildren(children ...Plan) { _ = "STUB: not implemented"; return }

func (p *basePlan) context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
