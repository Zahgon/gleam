package flow

import (
	"context"
	"time"
)

func newDataset(context *Flow) *Dataset { _ = "STUB: not implemented"; return nil }

func (d *Dataset) GetShards() []*DatasetShard {
	_ = "STUB: not implemented"

	// Run starts the whole flow. This is a convenient method, same as *Flow.Run()
	return nil
}

func (d *Dataset) Run(option ...FlowOption) { _ = "STUB: not implemented"; return }

// Run starts the whole flow. This is a convenient method, same as *Flow.RunContext()
func (d *Dataset) RunContext(ctx context.Context, option ...FlowOption) {
	_ = "STUB: not implemented"
	return
}

func (s *DatasetShard) Closed() bool { _ = "STUB: not implemented"; return false }

func (s *DatasetShard) TimeTaken() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (s *DatasetShard) Name() string { _ = "STUB: not implemented"; return "" }
