package distributed

import (
	"context"

	"github.com/chrislusf/gleam/flow"
)

type DistributedPlanner struct {
}

func Planner() *DistributedPlanner { _ = "STUB: not implemented"; return nil }

func (o *DistributedPlanner) GetFlowRunner() flow.FlowRunner {
	_ = "STUB: not implemented"

	// driver runs on local, controlling all tasks
	return *new(flow.FlowRunner)
}

func (fcd *DistributedPlanner) RunFlowContext(ctx context.Context, fc *flow.Flow) {
	_ = "STUB: not implemented"
	return
}
