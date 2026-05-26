// Copyright 2017 PingCAP, Inc.
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

// extractCorColumnsBySchema only extracts the correlated columns that match the outer plan's schema.
// e.g. If the correlated columns from inner plan are [t1.a, t2.a, t3.a] and outer plan's schema is [t2.a, t2.b, t2.c],
// only [t2.a] is treated as this apply's correlated column.
func (a *Apply) extractCorColumnsBySchema() { _ = "STUB: not implemented"; return }

// Shrink slice. e.g. [col1, nil, col2, nil] will be changed to [col1, col2].

// decorrelate function tries to convert apply plan to join plan.
func decorrelate(p LogicalPlan) LogicalPlan { _ = "STUB: not implemented"; return *new(LogicalPlan) }

// If the inner plan is non-correlated, the apply will be simplified to join.

// If the inner plan is a selection, we add this condition to join predicates.
// Notice that no matter what kind of join is, it's always right.

// TODO: Deal with aggregation and projection.
