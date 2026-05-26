// Copyright 2016 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
// // Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package plan

import (
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/expression"
	"github.com/chrislusf/gleam/sql/util/types"
)

type aggPushDownSolver struct {
	alloc *idAllocator
	ctx   context.Context
}

// isDecomposable checks if an aggregate function is decomposable. An aggregation function $F$ is decomposable
// if there exist aggregation functions F_1 and F_2 such that F(S_1 union all S_2) = F_2(F_1(S_1),F_1(S_2)),
// where S_1 and S_2 are two sets of values. We call S_1 and S_2 partial groups.
// It's easy to see that max, min, first row is decomposable, no matter whether it's distinct, but sum(distinct) and
// count(distinct) is not.
// Currently we don't support avg and concat.
func (a *aggPushDownSolver) isDecomposable(fun expression.AggregationFunction) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO: Support avg push down.

// getAggFuncChildIdx gets which children it belongs to, 0 stands for left, 1 stands for right, -1 stands for both.
func (a *aggPushDownSolver) getAggFuncChildIdx(aggFunc expression.AggregationFunction, schema expression.Schema) int {
	_ = "STUB: not implemented"
	return 0
}

// collectAggFuncs collects all aggregate functions and splits them into two parts: "leftAggFuncs" and "rightAggFuncs" whose
// arguments are all from left child or right child separately. If some aggregate functions have the arguments that have
// columns both from left and right children, the whole aggregation is forbidden to push down.
func (a *aggPushDownSolver) collectAggFuncs(agg *Aggregation, join *Join) (valid bool, leftAggFuncs, rightAggFuncs []expression.AggregationFunction) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// collectGbyCols collects all columns from gby-items and join-conditions and splits them into two parts: "leftGbyCols" and
// "rightGbyCols". e.g. For query "SELECT SUM(B.id) FROM A, B WHERE A.c1 = B.c1 AND A.c2 != B.c2 GROUP BY B.c3" , the optimized
// query should be "SELECT SUM(B.agg) FROM A, (SELECT SUM(id) as agg, c1, c2, c3 FROM B GROUP BY id, c1, c2, c3) as B
// WHERE A.c1 = B.c1 AND A.c2 != B.c2 GROUP BY B.c3". As you see, all the columns appearing in join-conditions should be
// treated as group by columns in join subquery.
func (a *aggPushDownSolver) collectGbyCols(agg *Aggregation, join *Join) (leftGbyCols, rightGbyCols []*expression.Column) {
	_ = "STUB: not implemented"
	return nil, nil
}

// extract equal conditions

func (a *aggPushDownSolver) splitAggFuncsAndGbyCols(agg *Aggregation, join *Join) (valid bool,
	leftAggFuncs, rightAggFuncs []expression.AggregationFunction,
	leftGbyCols, rightGbyCols []*expression.Column) {
	_ = "STUB: not implemented"
	return false, nil, nil, nil, nil
}

// addGbyCol adds a column to gbyCols. If a group by column has existed, it will not be added repeatedly.
func (a *aggPushDownSolver) addGbyCol(gbyCols []*expression.Column, cols ...*expression.Column) []*expression.Column {
	_ = "STUB: not implemented"
	return nil
}

// checkValidJoin checks if this join should be pushed across.
func (a *aggPushDownSolver) checkValidJoin(join *Join) bool {
	_ = "STUB: not implemented"
	return false
}

// decompose splits an aggregate function to two parts: a final mode function and a partial mode function. Currently
// there are no differences between partial mode and complete mode, so we can confuse them.
func (a *aggPushDownSolver) decompose(aggFunc expression.AggregationFunction, schema expression.Schema, id string) ([]expression.AggregationFunction, expression.Schema) {
	_ = "STUB: not implemented"
	// Result is a slice because avg should be decomposed to sum and count. Currently we don't process this case.
	return nil, *new(expression.Schema)
}

// useless but for debug

func (a *aggPushDownSolver) allFirstRow(aggFuncs []expression.AggregationFunction) bool {
	_ = "STUB: not implemented"
	return false
}

// tryToPushDownAgg tries to push down an aggregate function into a join path. If all aggFuncs are first row, we won't
// process it temporarily. If not, We will add additional group by columns and first row functions. We make a new aggregation
// operator.
func (a *aggPushDownSolver) tryToPushDownAgg(aggFuncs []expression.AggregationFunction, gbyCols []*expression.Column, join *Join, childIdx int) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

// If agg has no group-by item, it will return a default value, which may cause some bugs.
// So here we add a group-by item forcely.

func (a *aggPushDownSolver) getDefaultValues(agg *Aggregation) ([]types.Datum, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (a *aggPushDownSolver) checkAnyCountAndSum(aggFuncs []expression.AggregationFunction) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *aggPushDownSolver) makeNewAgg(aggFuncs []expression.AggregationFunction, gbyCols []*expression.Column) *Aggregation {
	_ = "STUB: not implemented"
	return nil
}

func (a *aggPushDownSolver) pushAggCrossUnion(agg *Aggregation, unionSchema expression.Schema, unionChild LogicalPlan) LogicalPlan {
	_ = "STUB: not implemented"
	return *new(LogicalPlan)
}

// aggPushDown tries to push down aggregate functions to join paths.
func (a *aggPushDownSolver) aggPushDown(p LogicalPlan) { _ = "STUB: not implemented"; return }

// If there exist count or sum functions in left join path, we can't push any
// aggregate function into right join path.

// TODO: This optimization is not always reasonable. We have not supported pushing projection to kv layer yet,
// so we must do this optimization.
