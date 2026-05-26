// Package flow contains data structure for computation.
// Mostly Dataset operations such as Map/Reduce/Join/Sort etc.
package flow

import (
	"context"
)

func New(name string) (fc *Flow) { _ = "STUB: not implemented"; return nil }

func (fc *Flow) Run(options ...FlowOption) { _ = "STUB: not implemented"; return }

func (fc *Flow) RunContext(ctx context.Context, options ...FlowOption) {
	_ = "STUB: not implemented"
	return
}

func (fc *Flow) NewNextDataset(shardSize int) (ret *Dataset) { _ = "STUB: not implemented"; return nil }

// the tasks should run on the source dataset shard
func (f *Flow) AddOneToOneStep(input *Dataset, output *Dataset) (step *Step) {
	_ = "STUB: not implemented"
	return nil
}

// setup the network

// the task should run on the destination dataset shard
func (f *Flow) AddAllToOneStep(input *Dataset, output *Dataset) (step *Step) {
	_ = "STUB: not implemented"
	return nil
}

// setup the network

// the task should run on the source dataset shard
// input is nil for initial source dataset
func (f *Flow) AddOneToAllStep(input *Dataset, output *Dataset) (step *Step) {
	_ = "STUB: not implemented"
	return nil
}

// setup the network

func (f *Flow) AddAllToAllStep(input *Dataset, output *Dataset) (step *Step) {
	_ = "STUB: not implemented"
	return nil
}

// setup the network

func (f *Flow) AddOneToEveryNStep(input *Dataset, n int, output *Dataset) (step *Step) {
	_ = "STUB: not implemented"
	return nil
}

// setup the network

func (f *Flow) AddLinkedNToOneStep(input *Dataset, m int, output *Dataset) (step *Step) {
	_ = "STUB: not implemented"
	return nil
}

// setup the network

// All dataset should have the same number of shards.
func (f *Flow) MergeDatasets1ShardTo1Step(inputs []*Dataset, output *Dataset) (step *Step) {
	_ = "STUB: not implemented"
	return nil
}

// setup the network

func fromStepToDataset(step *Step, output *Dataset) { _ = "STUB: not implemented"; return }

func fromDatasetToStep(input *Dataset, step *Step) { _ = "STUB: not implemented"; return }

func setupDatasetShard(d *Dataset, n int) { _ = "STUB: not implemented"; return }

func fromDatasetShardToTask(shard *DatasetShard, task *Task) { _ = "STUB: not implemented"; return }

func fromTaskToDatasetShard(task *Task, shard *DatasetShard) { _ = "STUB: not implemented"; return }
