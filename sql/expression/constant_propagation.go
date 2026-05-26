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
	"github.com/chrislusf/gleam/sql/context"
)

// MaxPropagateColsCnt means the max number of columns that can participate propagation.
var MaxPropagateColsCnt = 100

var eqFuncNameMap = map[string]bool{
	ast.EQ: true,
}

// inEqFuncNameMap stores all the in-equal operators that can be propagated.
var inEqFuncNameMap = map[string]bool{
	ast.LT: true,
	ast.GT: true,
	ast.LE: true,
	ast.GE: true,
	ast.NE: true,
}

type multiEqualSet struct {
	parent []int
}

func (m *multiEqualSet) init(l int) {
	m.parent = make([]int, l)
	for i := range m.parent {
		m.parent[i] = i
	}
}

func (m *multiEqualSet) addRelation(a int, b int) { _ = "STUB: not implemented"; return }

func (m *multiEqualSet) findRoot(a int) int { _ = "STUB: not implemented"; return 0 }

type propagateConstantSolver struct {
	colMapper  map[string]int // colMapper maps column to its index
	unionSet   *multiEqualSet // unionSet stores the relations like col_i = col_j
	eqList     []*Constant    // if eqList[i] != nil, it means col_i = eqList[i]
	columns    []*Column      // columns stores all columns appearing in the conditions
	conditions []Expression
	ctx        context.Context
}

// propagateInEQ propagates all in-equal conditions.
// e.g. For expression a = b and b = c and c = d and c < 1 , we can get extra a < 1 and b < 1 and d < 1.
// We maintain a unionSet representing the equivalent for every two columns.
func (s *propagateConstantSolver) propagateInEQ() { _ = "STUB: not implemented"; return }

// propagatesEQ propagates equal expression multiple times. An example runs as following:
// a = d & b * 2 = c & c = d + 2 & b = 1 & a = 4, we pick eq cond b = 1 and a = 4
// d = 4 & 2 = c & c = d + 2 & b = 1 & a = 4, we propagate b = 1 and a = 4 and pick eq cond c = 2 and d = 4
// d = 4 & 2 = c & false & b = 1 & a = 4, we propagate c = 2 and d = 4, and do constant folding: c = d + 2 will be folded as false.
func (s *propagateConstantSolver) propagateEQ() { _ = "STUB: not implemented"; return }

// validPropagateCond checks if the cond is an expression like [column op constant] and op is in the funNameMap.
func (s *propagateConstantSolver) validPropagateCond(cond Expression, funNameMap map[string]bool) (*Column, *Constant) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *propagateConstantSolver) setConds2ConstFalse() { _ = "STUB: not implemented"; return }

// pickNewEQConds tries to pick new equal conds and puts them to retMapper.
func (s *propagateConstantSolver) pickNewEQConds(visited []bool) (retMapper map[int]*Constant) {
	_ = "STUB: not implemented"
	return nil
}

// Then we check if this CNF item is a false constant. If so, we will set the whole condition to false.

// tryToUpdateEQList tries to update the eqList. When the eqList has store this column with a different constant, like
// a = 1 and a = 2, we set the second return value to false.
func (s *propagateConstantSolver) tryToUpdateEQList(col *Column, con *Constant) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (s *propagateConstantSolver) solve(conditions []Expression) []Expression {
	_ = "STUB: not implemented"
	return nil
}

func (s *propagateConstantSolver) getColID(col *Column) int { _ = "STUB: not implemented"; return 0 }

func (s *propagateConstantSolver) insertCol(col *Column) { _ = "STUB: not implemented"; return }

// PropagateConstant propagate constant values of equality predicates and inequality predicates in a condition.
func PropagateConstant(ctx context.Context, conditions []Expression) []Expression {
	_ = "STUB: not implemented"
	return nil
}
