package scheduler

import (
	"github.com/chrislusf/gleam/distributed/driver/scheduler/market"
)

// Requirement is TaskGroup
// Object is Agent's Location
func (s *Scheduler) Fetch(demands []market.Demand) { _ = "STUB: not implemented"; return }

// log.Printf("%s No more new executors.", s.Master)

// log.Printf("%s allocated %d executors with %d MB memory.", s.Master, len(result.Allocations), allocatedMemory)
