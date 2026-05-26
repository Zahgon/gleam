package executor

import (
	"github.com/chrislusf/gleam/flow"
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/expression"
)

type UnionScanExec struct {
	ctx       context.Context
	Src       Executor
	desc      bool
	condition expression.Expression

	schema expression.Schema
}

// Schema implements the Executor Schema interface.
func (e *UnionScanExec) Schema() expression.Schema {
	_ = "STUB: not implemented"

	// Next implements the Executor Next interface.
	return *new(expression.Schema)
}

func (e *UnionScanExec) Exec() *flow.Dataset { _ = "STUB: not implemented"; return nil }
