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

import (
	"github.com/chrislusf/gleam/sql/expression"
)

func (p *Aggregation) buildKeyInfo() { _ = "STUB: not implemented"; return }

// dealing with p.GroupbyCols
// This is only used for optimization and needn't to be pushed up, so only one is enough.

// A bijection exists between columns of a projection's schema and this projection's Exprs.
// Sometimes we need a schema made by expr of Exprs to convert a column in child's schema to a column in this projection's Schema.
func (p *Projection) buildSchemaByExprs() expression.Schema {
	_ = "STUB: not implemented"
	return *new(expression.Schema)
}

// If the expression is not a column, we add a column to occupy the position.

func (p *Projection) buildKeyInfo() { _ = "STUB: not implemented"; return }

func (p *Trim) buildKeyInfo() { _ = "STUB: not implemented"; return }

func (p *Join) buildKeyInfo() { _ = "STUB: not implemented"; return }

// If there is no equal conditions, then cartesian product can't be prevented and unique key information will destroy.

// Such as 'select * from t1 join t2 where t1.a = t2.a and t1.b = t2.b'.
// If one sides (a, b) is a unique key, then the unique key information is remained.
// But we don't consider this situation currently.
// Only key made by one column is considered now.

// For inner join, if one side of one equal condition is unique key,
// another side's unique key information will all be reserved.
// If it's an outer join, NULL value will fill some position, which will destroy the unique key information.

func (p *DataSource) buildKeyInfo() { _ = "STUB: not implemented"; return }

// The columns of this index should all occur in column schema.
// Since null value could be duplicate in unique key. So we check NotNull flag of every column.
