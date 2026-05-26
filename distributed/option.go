package distributed

import (
	"github.com/chrislusf/gleam/distributed/resource"
	"github.com/chrislusf/gleam/flow"
)

type DistributedOption struct {
	RequiredFiles []resource.FileResource
	Master        string
	DataCenter    string
	Rack          string
	TaskMemoryMB  int
	FlowBid       float64
	Module        string
	IsProfiling   bool
	BinaryPath    string
}

func Option() *DistributedOption { _ = "STUB: not implemented"; return nil }

func (o *DistributedOption) GetFlowRunner() flow.FlowRunner {
	_ = "STUB: not implemented"
	return *new(flow.FlowRunner)
}

func (o *DistributedOption) SetDataCenter(dataCenter string) *DistributedOption {
	_ = "STUB: not implemented"
	return nil
}

func (o *DistributedOption) SetMaster(master string) *DistributedOption {
	_ = "STUB: not implemented"
	return nil
}

// SetProfiling profiling will generate cpu and memory profile files when the executors are completed.
func (o *DistributedOption) SetProfiling(isProfiling bool) *DistributedOption {
	_ = "STUB: not implemented"
	return nil
}

// WithFile sends any related file over to gleam agents
// so the task can still access these files on gleam agents.
// The files are placed on the executed task's current working directory.
func (o *DistributedOption) WithFile(relatedFile, toFolder string) *DistributedOption {
	_ = "STUB: not implemented"
	return nil
}

// WithBinary allows you to provide a binary built to run
// on the gleam agents' architecture.
func (o *DistributedOption) WithBinary(binaryPath string) *DistributedOption {
	_ = "STUB: not implemented"
	return nil
}
