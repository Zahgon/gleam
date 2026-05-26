// Package driver coordinates distributed execution.
package driver

import (
	"context"
	"sync"

	"github.com/chrislusf/gleam/distributed/driver/scheduler"
	"github.com/chrislusf/gleam/distributed/plan"
	"github.com/chrislusf/gleam/distributed/resource"
	"github.com/chrislusf/gleam/flow"
	"github.com/chrislusf/gleam/pb"
)

type Option struct {
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

type FlowDriver struct {
	Option *Option

	stepGroups []*plan.StepGroup
	taskGroups []*plan.TaskGroup

	status *pb.FlowExecutionStatus
}

func NewFlowDriver(option *Option) *FlowDriver { _ = "STUB: not implemented"; return nil }

// driver runs on local, controlling all tasks
func (fcd *FlowDriver) RunFlowContext(parentCtx context.Context, fc *flow.Flow) {
	_ = "STUB: not implemented"

	// task fusion to minimize disk IO
	return
}

// create the scheduler

// best effort to clean data on agent disk
// this may need more improvements

// schedule to run the steps

func (fcd *FlowDriver) cleanup(sched *scheduler.Scheduler, fc *flow.Flow) {
	_ = "STUB: not implemented"
	return
}

// TODO send the pprof files back to driver

func (fcd *FlowDriver) reportStatus(ctx context.Context, wg *sync.WaitGroup, master string, stopChan chan bool) {
	_ = "STUB: not implemented"
	return
}

// println("grpc closing....")
