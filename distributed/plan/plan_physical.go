package plan

import (
	"github.com/chrislusf/gleam/flow"
)

// group local tasks into one task group
func translateToTaskGroups(stepId2StepGroup []*StepGroup) (ret []*TaskGroup) {
	_ = "STUB: not implemented"
	return nil
}

// println("dealing with", stepGroup.Steps[0].Name, "tasks:", len(stepGroup.Steps[0].Tasks))

// depends on the previous step group
// MAYBE IMPROVEMENT: depends on a subset of previus shards

func assertSameNumberOfTasks(steps []*flow.Step) { _ = "STUB: not implemented"; return }
