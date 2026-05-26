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

func getUsedList(usedCols []*expression.Column, schema expression.Schema) []bool {
	_ = "STUB: not implemented"
	return nil
}

// exprHasSetVar checks if the expression has set-var function. If do, we should not prune it.
func exprHasSetVar(expr expression.Expression) bool { _ = "STUB: not implemented"; return false }

// PruneColumns implements LogicalPlan interface.
func (p *Projection) PruneColumns(parentUsedCols []*expression.Column) {
	_ = "STUB: not implemented"
	return
}

// PruneColumns implements LogicalPlan interface.
func (p *Selection) PruneColumns(parentUsedCols []*expression.Column) {
	_ = "STUB: not implemented"
	return
}

// PruneColumns implements LogicalPlan interface.
func (p *Aggregation) PruneColumns(parentUsedCols []*expression.Column) {
	_ = "STUB: not implemented"
	return
}

// PruneColumns implements LogicalPlan interface.
func (p *Sort) PruneColumns(parentUsedCols []*expression.Column) { _ = "STUB: not implemented"; return }

// PruneColumns implements LogicalPlan interface.
func (p *Union) PruneColumns(parentUsedCols []*expression.Column) {
	_ = "STUB: not implemented"
	return
}

// PruneColumns implements LogicalPlan interface.
func (p *DataSource) PruneColumns(parentUsedCols []*expression.Column) {
	_ = "STUB: not implemented"
	return
}

// PruneColumns implements LogicalPlan interface.
func (p *TableDual) PruneColumns(_ []*expression.Column) {
	_ = "STUB: not implemented"

	// PruneColumns implements LogicalPlan interface.
	return
}

func (p *Trim) PruneColumns(parentUsedCols []*expression.Column) { _ = "STUB: not implemented"; return }

// PruneColumns implements LogicalPlan interface.
func (p *Exists) PruneColumns(parentUsedCols []*expression.Column) {
	_ = "STUB: not implemented"
	return
}

// PruneColumns implements LogicalPlan interface.
func (p *Insert) PruneColumns(_ []*expression.Column) { _ = "STUB: not implemented"; return }

// PruneColumns implements LogicalPlan interface.
func (p *Join) PruneColumns(parentUsedCols []*expression.Column) { _ = "STUB: not implemented"; return }

// PruneColumns implements LogicalPlan interface.
func (p *Apply) PruneColumns(parentUseCols []*expression.Column) { _ = "STUB: not implemented"; return }

// PruneColumns implements LogicalPlan interface.
func (p *Update) PruneColumns(parentUsedCols []*expression.Column) {
	_ = "STUB: not implemented"
	return
}

// PruneColumns implements LogicalPlan interface.
func (p *Delete) PruneColumns(parentUsedCols []*expression.Column) {
	_ = "STUB: not implemented"
	return
}
