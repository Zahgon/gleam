package executor

import (
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/infoschema"
	"github.com/chrislusf/gleam/sql/plan"
)

// executorBuilder builds an Executor from a Plan.
// The InfoSchema must not change during execution.
type executorBuilder struct {
	ctx context.Context
	is  infoschema.InfoSchema
	// If there is any error during Executor building process, err is set.
	err error
}

func newExecutorBuilder(ctx context.Context, is infoschema.InfoSchema) *executorBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *executorBuilder) build(p plan.Plan) Executor {
	_ = "STUB: not implemented"
	return *new(Executor)
}

// b.buildSemiJoin(v)

func (b *executorBuilder) buildLimit(v *plan.Limit) Executor {
	_ = "STUB: not implemented"
	return *new(Executor)
}

func (b *executorBuilder) buildUnionScanExec(v *plan.PhysicalUnionScan) Executor {
	_ = "STUB: not implemented"
	return *new(Executor)
}

/*
	case *XSelectIndexExec:
		us.desc = x.indexPlan.Desc
		for _, ic := range x.indexPlan.Index.Columns {
			for i, col := range x.indexPlan.GetSchema().Columns {
				if col.ColName.L == ic.Name.L {
					us.usedIndex = append(us.usedIndex, i)
					break
				}
			}
		}
		us.dirty = getDirtyDB(b.ctx).getDirtyTable(x.table.Meta().ID)
		us.condition = v.Condition
		us.buildAndSortAddedRows(x.table, x.asName)
*/

// The mem table will not be written by sql directly, so we can omit the union scan to avoid err reporting.

func (b *executorBuilder) buildJoin(v *plan.PhysicalHashJoin) Executor {
	_ = "STUB: not implemented"
	return *new(Executor)
}

func (b *executorBuilder) buildAggregation(v *plan.PhysicalAggregation) Executor {
	_ = "STUB: not implemented"
	return *new(Executor)
}

func (b *executorBuilder) buildSelection(v *plan.Selection) Executor {
	_ = "STUB: not implemented"
	return *new(Executor)
}

func (b *executorBuilder) buildProjection(v *plan.Projection) Executor {
	_ = "STUB: not implemented"
	return *new(Executor)
}

func (b *executorBuilder) buildTableDual(v *plan.TableDual) Executor {
	_ = "STUB: not implemented"
	return *new(Executor)
}

func (b *executorBuilder) buildTableScan(v *plan.PhysicalTableScan) Executor {
	_ = "STUB: not implemented"
	return *new(Executor)
}

// where:       v.TableConditionPBExpr,

// aggFuncs:    v.AggFuncsPB,

// byItems:     v.GbyItemsPB,
// orderByList: v.SortItemsPB,

func (b *executorBuilder) buildSort(v *plan.Sort) Executor {
	_ = "STUB: not implemented"
	return *new(Executor)
}

func (b *executorBuilder) buildApply(v *plan.PhysicalApply) Executor {
	_ = "STUB: not implemented"
	return *new(Executor)
}

func (b *executorBuilder) buildUnion(v *plan.Union) Executor {
	_ = "STUB: not implemented"
	return *new(Executor)
}
