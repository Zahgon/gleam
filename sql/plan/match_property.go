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
	"github.com/chrislusf/gleam/sql/model"
)

// matchProperty implements PhysicalPlan matchProperty interface.
func (ts *PhysicalTableScan) matchProperty(prop *requiredProperty, infos ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// If there exists a table filter, we should calculate the filter scan cost.

func allMatch(matchedList []bool) bool { _ = "STUB: not implemented"; return false }

// matchPropColumn checks if the idxCol match one of columns in required property and return the matched index.
// If no column is matched, return -1.
func matchPropColumn(prop *requiredProperty, matchedIdx int, idxCol *model.IndexColumn) int {
	_ = "STUB: not implemented"
	return 0
}

// When walking through the first sorKeyLen column,
// we should make sure to match them as the columns order exactly.
// So we must check the column in position of matchedIdx.

// When walking outside the first sorKeyLen column, we can match the columns as any order.

// matchProperty implements PhysicalPlan matchProperty interface.
func (is *PhysicalIndexScan) matchProperty(prop *requiredProperty, infos ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *PhysicalHashSemiJoin) matchProperty(_ *requiredProperty, childPlanInfo ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *PhysicalApply) matchProperty(_ *requiredProperty, childPlanInfo ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

func estimateJoinCount(lc uint64, rc uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *PhysicalHashJoin) matchProperty(prop *requiredProperty, childPlanInfo ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *Union) matchProperty(_ *requiredProperty, childPlanInfo ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *Selection) matchProperty(prop *requiredProperty, childPlanInfo ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *PhysicalUnionScan) matchProperty(prop *requiredProperty, childPlanInfo ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *Projection) matchProperty(_ *requiredProperty, childPlanInfo ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *Cache) matchProperty(prop *requiredProperty, childPlanInfo ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *MaxOneRow) matchProperty(_ *requiredProperty, _ ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *Exists) matchProperty(_ *requiredProperty, _ ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *Trim) matchProperty(_ *requiredProperty, _ ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *PhysicalAggregation) matchProperty(prop *requiredProperty, _ ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *Limit) matchProperty(_ *requiredProperty, _ ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *TableDual) matchProperty(_ *requiredProperty, _ ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *Sort) matchProperty(_ *requiredProperty, _ ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *Insert) matchProperty(_ *requiredProperty, _ ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *SelectLock) matchProperty(_ *requiredProperty, _ ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *Update) matchProperty(_ *requiredProperty, _ ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *PhysicalDummyScan) matchProperty(_ *requiredProperty, _ ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *Delete) matchProperty(_ *requiredProperty, _ ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *Show) matchProperty(_ *requiredProperty, _ ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}

// matchProperty implements PhysicalPlan matchProperty interface.
func (p *PhysicalMemTable) matchProperty(_ *requiredProperty, _ ...*physicalPlanInfo) *physicalPlanInfo {
	_ = "STUB: not implemented"
	return nil
}
