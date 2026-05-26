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
	"github.com/chrislusf/gleam/sql/expression"
)

// tryToGetJoinGroup tries to fetch a whole join group, which all joins is cartesian join.
func tryToGetJoinGroup(j *Join) ([]LogicalPlan, bool) { _ = "STUB: not implemented"; return nil, false }

func findColumnIndexByGroup(groups []LogicalPlan, col *expression.Column) int {
	_ = "STUB: not implemented"
	return 0
}

type joinReOrderSolver struct {
	graph      []edgeList
	group      []LogicalPlan
	visited    []bool
	resultJoin LogicalPlan
	groupRank  []*rankInfo
	allocator  *idAllocator
}

type edgeList []*rankInfo

func (l edgeList) Len() int { _ = "STUB: not implemented"; return 0 }

func (l edgeList) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (l edgeList) Swap(i, j int) { _ = "STUB: not implemented"; return }

type rankInfo struct {
	nodeID int
	rate   float64
}

func (e *joinReOrderSolver) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (e *joinReOrderSolver) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (e *joinReOrderSolver) Len() int { _ = "STUB: not implemented"; return 0 }

// reorderJoin implements a simple join reorder algorithm. It will extract all the equal conditions and compose them to a graph.
// Then walk through the graph and pick the nodes connected by some edges to compose a join tree.
// We will pick the node with least result set as early as possible.
func (e *joinReOrderSolver) reorderJoin(group []LogicalPlan, conds []expression.Expression) {
	_ = "STUB: not implemented"
	return
}

// TODO: Estimate it more precisely in future.

// Make cartesian join as bushy tree.
func (e *joinReOrderSolver) makeBushyJoin(cartesianJoinGroup []LogicalPlan) {
	_ = "STUB: not implemented"
	return
}

func (e *joinReOrderSolver) newJoin(lChild, rChild LogicalPlan) *Join {
	_ = "STUB: not implemented"
	return nil
}

// walkGraph implements a dfs algorithm. Each time it picks a edge with lowest rate, which has been sorted before.
func (e *joinReOrderSolver) walkGraphAndComposeJoin(u int) { _ = "STUB: not implemented"; return }
