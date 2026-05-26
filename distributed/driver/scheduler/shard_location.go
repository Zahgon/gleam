package scheduler

import (
	"sync"

	"github.com/chrislusf/gleam/flow"
	"github.com/chrislusf/gleam/pb"
)

type DatasetShardLocator struct {
	sync.Mutex
	datasetShard2Location     map[string]pb.DataLocation
	datasetShard2LocationLock sync.Mutex
	waitForAllInputs          *sync.Cond
}

func NewDatasetShardLocator() *DatasetShardLocator { _ = "STUB: not implemented"; return nil }

func (l *DatasetShardLocator) GetShardLocation(shardName string) (pb.DataLocation, bool) {
	_ = "STUB: not implemented"
	return *new(pb.DataLocation), false
}

func (l *DatasetShardLocator) SetShardLocation(name string, location pb.DataLocation) {
	_ = "STUB: not implemented"
	return
}

// fmt.Printf("shard %s is at %s\n", name, location.URL())

func (l *DatasetShardLocator) isDatasetShardRegistered(shard *flow.DatasetShard) bool {
	_ = "STUB: not implemented"
	return false
}

// fmt.Printf("%s's waiting for %s, but it is not ready\n", shard.Dataset.Step.Name, shard.Name())

// fmt.Printf("%s knows %s is ready\n", shard.Dataset.Step.Name, shard.Name())

func (l *DatasetShardLocator) waitForInputDatasetShardLocations(task *flow.Task) {
	_ = "STUB: not implemented"
	return
}

func (l *DatasetShardLocator) waitForOutputDatasetShardLocations(task *flow.Task) {
	_ = "STUB: not implemented"
	return
}

func (l *DatasetShardLocator) allInputLocations(task *flow.Task) string {
	_ = "STUB: not implemented"
	return ""
}
