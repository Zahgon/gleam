package executor

import (
	"net"

	"github.com/chrislusf/gleam/pb"
)

func (exe *Executor) serveGrpc(listener net.Listener) { _ = "STUB: not implemented"; return }

// Collect stat from "gleam execute" started mapper reducer process
func (exe *Executor) CollectExecutionStatistics(stream pb.GleamExecutor_CollectExecutionStatisticsServer) error {
	_ = "STUB: not implemented"
	return nil
}

// fmt.Printf("executor received stat: %+v\n", stat)
