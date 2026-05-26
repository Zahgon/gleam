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
)

const (
	netWorkFactor   = 1.5
	memoryFactor    = 5.0
	selectionFactor = 0.8
	distinctFactor  = 0.7
	cpuFactor       = 0.9
	aggFactor       = 0.1
	joinFactor      = 0.3
)

// JoinConcurrency means the number of goroutines that participate in joining.
var JoinConcurrency = 5

func (p *DataSource) convert2TableScan(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *DataSource) convert2IndexScan(prop *requiredProperty, index *model.IndexInfo) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isCoveringIndex(columns []*model.ColumnInfo, indexColumns []*model.IndexColumn, pkIsHandle bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *DataSource) need2ConsiderIndex(prop *requiredProperty) bool {
	_ = "STUB: not implemented"
	return false
}

// convert2PhysicalPlan implements the LogicalPlan convert2PhysicalPlan interface.
// If there is no index that matches the required property, the returned physicalPlanInfo
// will be table scan and has the cost of MaxInt64. But this can be ignored because the parent will call
// convert2PhysicalPlan again with an empty *requiredProperty, so the plan with the lowest
// cost will be chosen.
func (p *DataSource) convert2PhysicalPlan(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// tryToConvert2DummyScan is an optimization which checks if its parent is a selection with a constant condition
// that evaluates to false. If it is, there is no need for a real physical scan, a dummy scan will do.
func (p *DataSource) tryToConvert2DummyScan(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// addPlanToResponse creates a *physicalPlanInfo that adds p as the parent of info.
func addPlanToResponse(parent PhysicalPlan, info *physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// enforceProperty creates a *physicalPlanInfo that satisfies the required property by adding
// sort or limit as the parent of the given physical plan.
func enforceProperty(prop *requiredProperty, info *physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

func sortCost(cnt uint64) float64 {
	_ = "STUB: not implemented"

	// If cnt is 0, the log(cnt) will be NAN.
	return 0
}

// removeLimit removes the limit from prop.
func removeLimit(prop *requiredProperty) *requiredProperty { _ = "STUB: not implemented"; return nil }

// convertLimitOffsetToCount changes the limit(offset, count) in prop to limit(0, offset + count).
func convertLimitOffsetToCount(prop *requiredProperty) *requiredProperty {
	_ = "STUB: not implemented"
	return nil
}

func limitProperty(limit *Limit) *requiredProperty { _ = "STUB: not implemented"; return nil }

// convert2PhysicalPlan implements the LogicalPlan convert2PhysicalPlan interface.
func (p *Limit) convert2PhysicalPlan(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert2PhysicalPlanSemi converts the semi join to *physicalPlanInfo.
func (p *Join) convert2PhysicalPlanSemi(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert2PhysicalPlanLeft converts the left join to *physicalPlanInfo.
func (p *Join) convert2PhysicalPlanLeft(prop *requiredProperty, innerJoin bool) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: decide concurrency by data size.

// replaceColsInPropBySchema replaces the columns in original prop with the columns in schema.
func replaceColsInPropBySchema(prop *requiredProperty, schema expression.Schema) *requiredProperty {
	_ = "STUB: not implemented"
	return nil
}

// convert2PhysicalPlanRight converts the right join to *physicalPlanInfo.
func (p *Join) convert2PhysicalPlanRight(prop *requiredProperty, innerJoin bool) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: decide concurrency by data size.

// convert2PhysicalPlan implements the LogicalPlan convert2PhysicalPlan interface.
func (p *Join) convert2PhysicalPlan(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert2PhysicalPlanStream converts the logical aggregation to the stream aggregation *physicalPlanInfo.
func (p *Aggregation) convert2PhysicalPlanStream(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Consider distinct key.

// group by a + b is not interested in any order.

// We should add columns in aggregation in order to keep index right.

// convert2PhysicalPlanFinalHash converts the logical aggregation to the final hash aggregation *physicalPlanInfo.
func (p *Aggregation) convert2PhysicalPlanFinalHash(x physicalDistSQLPlan, childInfo *physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// if we build the final aggregation, it must be the best plan.

// convert2PhysicalPlanCompleteHash converts the logical aggregation to the complete hash aggregation *physicalPlanInfo.
func (p *Aggregation) convert2PhysicalPlanCompleteHash(childInfo *physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// convert2PhysicalPlanHash converts the logical aggregation to the physical hash aggregation.
func (p *Aggregation) convert2PhysicalPlanHash() (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert2PhysicalPlan implements the LogicalPlan convert2PhysicalPlan interface.
func (p *Aggregation) convert2PhysicalPlan(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert2PhysicalPlan implements the LogicalPlan convert2PhysicalPlan interface.
func (p *Union) convert2PhysicalPlan(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert2PhysicalPlan implements the LogicalPlan convert2PhysicalPlan interface.
func (p *Selection) convert2PhysicalPlan(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Firstly, we try to push order.

// Secondly, we push nothing and enforce this property.

func (p *Selection) convert2PhysicalPlanPushOrder(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert2PhysicalPlanEnforce converts a selection to *physicalPlanInfo which does not push the
// required property to the children, but enforce the property instead.
func (p *Selection) convert2PhysicalPlanEnforce(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert2PhysicalPlan implements the LogicalPlan convert2PhysicalPlan interface.
func (p *Projection) convert2PhysicalPlan(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func matchProp(ctx context.Context, target, new *requiredProperty) bool {
	_ = "STUB: not implemented"
	return false
}

// convert2PhysicalPlan implements the LogicalPlan convert2PhysicalPlan interface.
func (p *Sort) convert2PhysicalPlan(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert2PhysicalPlan implements the LogicalPlan convert2PhysicalPlan interface.
func (p *Apply) convert2PhysicalPlan(prop *requiredProperty) (*physicalPlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// physicalInitialize will set value of some attributes after convert2PhysicalPlan process.
// Currently, only attribute "correlated" is considered.
func physicalInitialize(p PhysicalPlan) { _ = "STUB: not implemented"; return }

// initialize attributes

// addCachePlan will add a Cache plan above the plan whose father's IsCorrelated() is true but its own IsCorrelated() is false.
func addCachePlan(p PhysicalPlan, allocator *idAllocator) { _ = "STUB: not implemented"; return }
