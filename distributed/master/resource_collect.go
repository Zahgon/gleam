package master

import (
	"github.com/chrislusf/gleam/pb"
)

func (tp *Topology) UpdateAgentInformation(ai *pb.Heartbeat) { _ = "STUB: not implemented"; return }

// fmt.Printf("hasOldInfo %+v, oldInfo %+v\n", hasOldInfo, oldInfo)

// fmt.Printf("deltaAllocated %+v\n", deltaAllocated)

func (tp *Topology) deleteAgentInformation(location *pb.Location) {
	_ = "STUB: not implemented"
	return
}

// fmt.Printf("deleting %+v\n", oldInfo)

func (tp *Topology) findAgentInformation(location *pb.Location) (*AgentInformation, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
