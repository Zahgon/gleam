package executor

import (
	"github.com/chrislusf/gleam/flow"
	"github.com/chrislusf/gleam/sql/expression"
)

type LimitExec struct {
	Src    Executor
	Offset uint64
	Count  uint64
	schema expression.Schema
}

// Schema implements the Executor Schema interface.
func (e *LimitExec) Schema() expression.Schema {
	_ = "STUB: not implemented"

	// Next implements the Executor Next interface.
	return *new(expression.Schema)
}

func (e *LimitExec) Exec() *flow.Dataset { _ = "STUB: not implemented"; return nil }
