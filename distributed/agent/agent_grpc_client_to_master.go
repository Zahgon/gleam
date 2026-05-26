package agent

import (
	"time"

	"github.com/chrislusf/gleam/pb"
)

func (as *AgentServer) heartbeat() { _ = "STUB: not implemented"; return }

func (as *AgentServer) doHeartbeat(sleepInterval time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (as *AgentServer) sendOneHeartbeat(stream pb.GleamMaster_SendHeartbeatClient) error {
	_ = "STUB: not implemented"
	return nil
}

// log.Printf("Reporting allocated %v", as.allocatedResource)
