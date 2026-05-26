package agent

import (
	"sync"

	"github.com/chrislusf/gleam/pb"
)

var (
	statsChanMap        = make(map[string]chan *pb.ExecutionStat)
	statsChanMapRWMutex sync.Mutex
)

func getStatsChan(flowHashCode uint32, stepId int32, taskId int32) chan *pb.ExecutionStat {
	_ = "STUB: not implemented"
	return nil
}

func deleteStatsChanByInstructionSet(instructionSet *pb.InstructionSet) {
	_ = "STUB: not implemented"
	return
}

func deleteStatsChan(flowHashCode uint32, stepId int32, taskId int32) {
	_ = "STUB: not implemented"
	return
}

func createStatsChanByInstructionSet(instructionSet *pb.InstructionSet) chan *pb.ExecutionStat {
	_ = "STUB: not implemented"
	return nil
}
