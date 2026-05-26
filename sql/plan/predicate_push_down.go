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
)

func addSelection(p Plan, child LogicalPlan, conditions []expression.Expression, allocator *idAllocator) error {
	_ = "STUB: not implemented"
	return nil
}

// PredicatePushDown implements LogicalPlan PredicatePushDown interface.
func (p *Selection) PredicatePushDown(predicates []expression.Expression) ([]expression.Expression, LogicalPlan, error) {
	_ = "STUB: not implemented"
	return nil, *new(LogicalPlan), nil
}

// PredicatePushDown implements LogicalPlan PredicatePushDown interface.
func (p *DataSource) PredicatePushDown(predicates []expression.Expression) ([]expression.Expression, LogicalPlan, error) {
	_ = "STUB: not implemented"
	return nil,

		// PredicatePushDown implements LogicalPlan PredicatePushDown interface.
		*new(LogicalPlan), nil
}

func (p *TableDual) PredicatePushDown(predicates []expression.Expression) ([]expression.Expression, LogicalPlan, error) {
	_ = "STUB: not implemented"
	return nil,

		// PredicatePushDown implements LogicalPlan PredicatePushDown interface.
		*new(LogicalPlan), nil
}

func (p *Join) PredicatePushDown(predicates []expression.Expression) (ret []expression.Expression, retPlan LogicalPlan, err error) {
	_ = "STUB: not implemented"
	return nil, *new(LogicalPlan), nil
}

// outerJoinSimplify simplifies outer join.
func outerJoinSimplify(p *Join, predicates []expression.Expression) error {
	_ = "STUB: not implemented"
	return nil
}

// first simplify embedded outer join.
// When trying to simplify an embedded outer join operation in a query,
// we must take into account the join condition for the embedding outer join together with the WHERE condition.

// then simplify embedding outer join.

// isNullRejected check whether a condition is null-rejected
// A condition would be null-rejected in one of following cases:
// If it is a predicate containing a reference to an inner table that evaluates to UNKNOWN or FALSE when one of its arguments is NULL.
// If it is a conjunction containing a null-rejected condition as a conjunct.
// If it is a disjunction of null-rejected conditions.
func isNullRejected(ctx context.Context, schema expression.Schema, expr expression.Expression) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// concatOnAndWhereConds concatenate ON conditions with WHERE conditions.
func concatOnAndWhereConds(join *Join, predicates []expression.Expression) []expression.Expression {
	_ = "STUB: not implemented"
	return nil
}

// PredicatePushDown implements LogicalPlan PredicatePushDown interface.
func (p *Projection) PredicatePushDown(predicates []expression.Expression) (ret []expression.Expression, retPlan LogicalPlan, err error) {
	_ = "STUB: not implemented"
	return nil, *new(LogicalPlan), nil
}

// PredicatePushDown implements LogicalPlan PredicatePushDown interface.
func (p *Union) PredicatePushDown(predicates []expression.Expression) (ret []expression.Expression, retPlan LogicalPlan, err error) {
	_ = "STUB: not implemented"
	return nil, *new(LogicalPlan), nil
}

// getGbyColIndex gets the column's index in the group-by columns.
func (p *Aggregation) getGbyColIndex(col *expression.Column) int {
	_ = "STUB: not implemented"
	return 0
}

// PredicatePushDown implements LogicalPlan PredicatePushDown interface.
func (p *Aggregation) PredicatePushDown(predicates []expression.Expression) (ret []expression.Expression, retPlan LogicalPlan, err error) {
	_ = "STUB: not implemented"
	return nil, *new(LogicalPlan), nil
}

// Consider SQL list "select sum(b) from t group by a having 1=0". "1=0" is a constant predicate which should be
// retained and pushed down at the same time. Because we will get a wrong query result that contains one column
// with value 0 rather than an empty query result.

// PredicatePushDown implements LogicalPlan PredicatePushDown interface.
func (p *Limit) PredicatePushDown(predicates []expression.Expression) ([]expression.Expression, LogicalPlan, error) {
	_ = "STUB: not implemented"
	// Limit forbids any condition to push down.
	return nil, *new(LogicalPlan), nil
}

// PredicatePushDown implements LogicalPlan PredicatePushDown interface.
func (p *MaxOneRow) PredicatePushDown(predicates []expression.Expression) ([]expression.Expression, LogicalPlan, error) {
	_ = "STUB: not implemented"
	// MaxOneRow forbids any condition to push down.
	return nil, *new(LogicalPlan), nil
}
