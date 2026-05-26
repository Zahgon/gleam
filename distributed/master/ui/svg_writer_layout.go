package ui

import (
	"github.com/chrislusf/gleam/pb"
)

// separate step groups into layers of step group ids via depencency analysis
func toStepGroupLayers(status *pb.FlowExecutionStatus) (layers [][]int) {
	_ = "STUB: not implemented"

	// how many step groups depenend on this step group
	return nil
}

// maintain dependencyCount after one layer

func checkAllDepencies(dep []int, isUsed []bool) (noDependencyStepGroupIds []int) {
	_ = "STUB: not implemented"
	return nil
}
