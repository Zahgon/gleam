// Schedule tasks to run on available resources assigned by master.
package scheduler

import (
	"sync"
	"time"

	"github.com/chrislusf/gleam/distributed/driver/scheduler/market"
	"github.com/chrislusf/gleam/pb"
)

type Scheduler struct {
	sync.Mutex

	Master       string
	EventChan    chan interface{}
	Market       *market.Market
	Option       *Option
	shardLocator *DatasetShardLocator
}

type RemoteExecutorStatus struct {
	Request      *pb.ExecutionRequest
	Allocation   *pb.Allocation
	RequestTime  time.Time
	InputLength  int
	OutputLength int
	ReadyTime    time.Time
	RunTime      time.Time
	StopTime     time.Time
}

type Option struct {
	Username     string
	Hostname     string
	FlowHashcode uint32
	DataCenter   string
	Rack         string
	TaskMemoryMB int
	Module       string
	IsProfiling  bool
}

func New(leader string, option *Option) *Scheduler { _ = "STUB: not implemented"; return nil }
