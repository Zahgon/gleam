package scheduler

import (
	"context"
	"sync"

	"github.com/chrislusf/gleam/distributed/plan"
	"github.com/chrislusf/gleam/flow"
	"github.com/chrislusf/gleam/pb"
)

func (s *Scheduler) remoteExecuteOnLocation(ctx context.Context,
	flowContext *flow.Flow,
	taskGroupStatus *pb.FlowExecutionStatus_TaskGroup,
	executionStatus *pb.FlowExecutionStatus_TaskGroup_Execution,
	taskGroup *plan.TaskGroup,
	allocation *pb.Allocation, wg *sync.WaitGroup) error {
	_ = "STUB: not implemented"

	// s.setupInputChannels(flowContext, tasks[0], allocation.Location, wg)
	return nil
}

// fmt.Printf("allocated %s on %v\n", tasks[0].Name(), allocation.Location)
// create reqeust

// println("RequestId:", taskGroup.RequestId, instructions.FlowHashCode)

func (s *Scheduler) localExecute(ctx context.Context,
	flowContext *flow.Flow,
	executionStatus *pb.FlowExecutionStatus_TaskGroup_Execution,
	task *flow.Task,
	wg *sync.WaitGroup) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scheduler) localExecuteSource(ctx context.Context,
	flowContext *flow.Flow,
	executionStatus *pb.FlowExecutionStatus_TaskGroup_Execution,
	task *flow.Task,
	wg *sync.WaitGroup) error {
	_ = "STUB: not implemented"
	return nil
}

// println(task.Step.Name, "writing to", shard.Name(), "at", location.Location.URL())

func (s *Scheduler) localExecuteOutput(ctx context.Context, flowContext *flow.Flow, task *flow.Task, wg *sync.WaitGroup) error {
	_ = "STUB: not implemented"
	return nil
}

// println(task.Step.Name, "reading from", shard.Name(), "at", location.Location.URL(), "to", inChan, "onDisk", shard.Dataset.GetIsOnDiskIO())
