package plan

import (
	"github.com/chrislusf/gleam/flow"
	"github.com/chrislusf/gleam/pb"
)

func TranslateToInstructionSet(taskGroups *TaskGroup) (ret *pb.InstructionSet) {
	_ = "STUB: not implemented"
	return nil
}

func translateToInstruction(task *flow.Task) (ret *pb.Instruction) {
	_ = "STUB: not implemented"
	return nil
}

// try to run Instruction first
// if failed, try to run shell scripts

// Command can come from Pipe() directly
// get an exec.Command
// println("processing step:", task.Step.Name)
