package executor

import (
	"time"

	"github.com/chrislusf/gleam/flow"
	"github.com/chrislusf/gleam/sql/context"
	"github.com/chrislusf/gleam/sql/infoschema"
	"github.com/chrislusf/gleam/sql/plan"
)

// statement implements the ast.Statement interface, it builds a plan.Plan to an ast.Statement.
type Statement struct {
	// The InfoSchema cannot change during execution, so we hold a reference to it.
	InfoSchema infoschema.InfoSchema
	ctx        context.Context
	Text       string
	Plan       plan.Plan
	startTime  time.Time
}

func (a *Statement) OriginText() string {
	_ = "STUB: not implemented"

	// Exec implements the ast.Statement Exec interface.
	// This function builds an Executor from a plan. If the Executor doesn't return result,
	// like the INSERT, UPDATE statements, it executes in this function, if the Executor returns
	// result, execution is done after this function returns, in the returned ast.RecordSet Next method.
	return ""
}

func (a *Statement) Exec(ctx context.Context) (*flow.Dataset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
