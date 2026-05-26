// Package driver coordinates distributed execution.
package driver

import (
	"github.com/chrislusf/gleam/distributed/plan"
	"github.com/chrislusf/gleam/flow"
	"github.com/chrislusf/gleam/pb"
)

func (fcd *FlowDriver) GetTaskGroupStatus(taskGroup *plan.TaskGroup) *pb.FlowExecutionStatus_TaskGroup {
	_ = "STUB: not implemented"
	return nil
}

func (fcd *FlowDriver) logExecutionPlan(fc *flow.Flow) { _ = "STUB: not implemented"; return }

// find the parent step group from all step groups

// if the first step is the same
