package plan

import (
	"github.com/chrislusf/gleam/flow"
)

func isMergeableDataset(ds *flow.Dataset, taskCount int) bool {
	_ = "STUB: not implemented"
	return false
}

// find mergeable parent step or itself if parent is not mergeable
func findAncestorStepId(step *flow.Step) (int, bool) { _ = "STUB: not implemented"; return 0, false }

// println("find step", step.Name)

// more than 2 dataset inputs

// no dataset inputs

// group local steps into one step group
func translateToStepGroups(fc *flow.Flow) []*StepGroup {
	_ = "STUB: not implemented"
	// use array instead of map to ensure consistent ordering
	return nil
}

// println("step:", step.Name, step.Id, "starting...")

// println("step:", step.Name, step.Id, "ancestorStepId", ancestorStepId)

// since we add steps following the same order as the code

// since we add steps following the same order as the code

// shrink

// println("add step group started by", stepGroup.Steps[0].Name, "with", len(stepGroup.Steps), "steps")
