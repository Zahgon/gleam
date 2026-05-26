package sql

import (
	"github.com/chrislusf/gleam/flow"
	"github.com/chrislusf/gleam/sql/executor"
	"github.com/chrislusf/gleam/sql/model"
	"github.com/chrislusf/gleam/sql/plan"
)

func RegisterTable(dataset *flow.Dataset, tableName string, columns []executor.TableColumn) {
	_ = "STUB: not implemented"
	return
}

func tableInfoList() (infos []*model.TableInfo) { _ = "STUB: not implemented"; return nil }

func Query(sql string) (*flow.Dataset, plan.Plan, error) {
	_ = "STUB: not implemented"
	return nil, *new(plan.Plan), nil
}
