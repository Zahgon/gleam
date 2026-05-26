package plan

import (
	"sync"
	"time"

	"github.com/chrislusf/gleam/flow"
	"github.com/chrislusf/gleam/pb"
)

type TaskGroup struct {
	Id              int
	Tasks           []*flow.Task
	Parents         []*TaskGroup
	ParentStepGroup *StepGroup
	RequestId       uint32 // id for actual request when running
	WaitAt          time.Time
	StartAt         time.Time
	StopAt          time.Time
	Error           error
}

type StepGroup struct {
	Steps      []*flow.Step
	Parents    []*StepGroup
	TaskGroups []*TaskGroup
	sync.Mutex
	waitForAllTasks *sync.Cond
}

func GroupTasks(fc *flow.Flow) ([]*StepGroup, []*TaskGroup) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewStepGroup() *StepGroup { _ = "STUB: not implemented"; return nil }

func (s *StepGroup) AddStep(Step *flow.Step) *StepGroup { _ = "STUB: not implemented"; return nil }

func (s *StepGroup) AddParent(parent *StepGroup) *StepGroup { _ = "STUB: not implemented"; return nil }

func NewTaskGroup() *TaskGroup { _ = "STUB: not implemented"; return nil }

func (t *TaskGroup) AddTask(task *flow.Task) *TaskGroup { _ = "STUB: not implemented"; return nil }

func (t *TaskGroup) AddParent(parent *TaskGroup) *TaskGroup { _ = "STUB: not implemented"; return nil }

func (t *TaskGroup) String() string { _ = "STUB: not implemented"; return "" }

func (t *TaskGroup) RequiredResources() *pb.ComputeResource { _ = "STUB: not implemented"; return nil }

// log.Printf("  %s : %s (%d MB)\n", t.String(), task.Step.Name, taskMemSize)

func (t *TaskGroup) MarkStop(err error) { _ = "STUB: not implemented"; return }

func (s *StepGroup) WaitForAllTasksToComplete() { _ = "STUB: not implemented"; return }
