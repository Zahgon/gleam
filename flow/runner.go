package flow

import (
	"context"
	"sync"
)

type FlowRunner interface {
	RunFlowContext(context.Context, *Flow)
}

type FlowOption interface {
	GetFlowRunner() FlowRunner
}

type localDriver struct {
	ctx context.Context
}

var (
	Local *localDriver
)

func init() {
	Local = &localDriver{}
}

func (r *localDriver) GetFlowRunner() FlowRunner {
	_ = "STUB: not implemented"
	return *new(FlowRunner)
}

func (r *localDriver) RunFlowContext(ctx context.Context, fc *Flow) {
	_ = "STUB: not implemented"
	return
}

func (r *localDriver) RunFlowAsync(wg *sync.WaitGroup, fc *Flow) { _ = "STUB: not implemented"; return }

func (r *localDriver) runDataset(wg *sync.WaitGroup, d *Dataset) { _ = "STUB: not implemented"; return }

func (r *localDriver) runDatasetShard(wg *sync.WaitGroup, shard *DatasetShard) {
	_ = "STUB: not implemented"
	return
}

// println("shard", shard.Name(), "moved", n, "bytes.")

func (r *localDriver) runStep(wg *sync.WaitGroup, step *Step) { _ = "STUB: not implemented"; return }

func (r *localDriver) runTask(wg *sync.WaitGroup, task *Task) {
	_ = "STUB: not implemented"

	// try to run Function first
	// if failed, try to run shell scripts
	return
}

// each function should close its own Piper output writer
// and close it's own Piper input reader

// get an exec.Command

// fmt.Printf("execCommand: %+v\n", execCommand)
