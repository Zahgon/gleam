package master

import (
	"time"

	"context"

	"github.com/chrislusf/gleam/pb"
)

type MasterServer struct {
	Topology     *Topology
	statusCache  *lru.Cache
	logDirectory string
	startTime    time.Time
}

func newMasterServer(logDirectory string) *MasterServer { _ = "STUB: not implemented"; return nil }

func (s *MasterServer) GetResources(ctx context.Context, in *pb.ComputeRequest) (*pb.AllocationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *MasterServer) SendHeartbeat(stream pb.GleamMaster_SendHeartbeatServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *MasterServer) SendFlowExecutionStatus(stream pb.GleamMaster_SendFlowExecutionStatusServer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *MasterServer) onStartup() { _ = "STUB: not implemented"; return }

// println("loading", f, "for", status.GetId())

func (s *MasterServer) onCacheEvict(key interface{}, value interface{}) {
	_ = "STUB: not implemented"
	return
}
