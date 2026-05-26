package scheduler

import (
	"github.com/chrislusf/gleam/distributed/driver/scheduler/market"
	"github.com/chrislusf/gleam/distributed/plan"
)

func (s *Scheduler) Score(r market.Requirement, bid float64, obj market.Object) float64 {
	_ = "STUB: not implemented"
	return 0
}

// log.Printf("Strange1: %s not allocated yet.", input.Name())

func memoryCost(tg *plan.TaskGroup) (cost int64) { _ = "STUB: not implemented"; return 0 }
