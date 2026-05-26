package scheduler

import (
	"github.com/chrislusf/gleam/distributed/plan"
	"github.com/chrislusf/gleam/flow"
	"github.com/chrislusf/gleam/pb"
)

func (s *Scheduler) DeleteOutput(taskGroup *plan.TaskGroup) { _ = "STUB: not implemented"; return }

// println("deleting", shard.Name(), "on", location.GetLocation().URL())

func needsInputFromDriver(task *flow.Task) bool { _ = "STUB: not implemented"; return false }

func isInputOnDisk(task *flow.Task) bool { _ = "STUB: not implemented"; return false }

func isRestartableTasks(tasks []*flow.Task) bool { _ = "STUB: not implemented"; return false }

func (s *Scheduler) GetShardLocation(shard *flow.DatasetShard) (pb.DataLocation, bool) {
	_ = "STUB: not implemented"
	return *new(pb.DataLocation), false
}

func (s *Scheduler) setShardLocation(shard *flow.DatasetShard, loc pb.DataLocation) {
	_ = "STUB: not implemented"
	return
}
