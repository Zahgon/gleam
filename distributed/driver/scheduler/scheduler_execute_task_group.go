package scheduler

import (
	"context"
	"sync"

	"github.com/chrislusf/gleam/distributed/plan"
	"github.com/chrislusf/gleam/distributed/resource"
	"github.com/chrislusf/gleam/flow"
	"github.com/chrislusf/gleam/pb"
)

// ExecuteTaskGroup wait for inputs and execute the task group remotely.
// If cancelled, the output will be cleaned up.
func (s *Scheduler) ExecuteTaskGroup(ctx context.Context,
	fc *flow.Flow,
	taskGroupStatus *pb.FlowExecutionStatus_TaskGroup,
	wg *sync.WaitGroup,
	taskGroup *plan.TaskGroup,
	bid float64, relatedFiles []resource.FileResource, binaryPath string) {
	_ = "STUB: not implemented"
	return
}

// these should be only one task on the driver side

// wait until inputs are registed

// for non-restartable taskGroup, wait until on disk inputs are completed

// fmt.Printf("inputs of %s is %s\n", tasks[0].Name(), s.allInputLocations(tasks[0]))

// get assigned executor location

// tell the driver to write to me

// println("registering", shard.Name(), "at", allocation.Location.URL())

// println("registering", shard.Name(), "at", allocation.Location.URL(), "onDisk", shard.Dataset.GetIsOnDiskIO())

// send driver code only when using go mapper reducer

// println("deleting", shard.Name(), "from", allocation.Location.URL())
