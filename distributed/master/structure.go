package master

import (
	"sync"
	"time"

	"github.com/chrislusf/gleam/pb"
)

type AgentInformation struct {
	Location      pb.Location
	LastHeartBeat time.Time
	Resource      pb.ComputeResource
	Allocated     pb.ComputeResource
}

type Rack struct {
	sync.RWMutex
	Name      string
	Agents    map[string]*AgentInformation
	Resource  pb.ComputeResource
	Allocated pb.ComputeResource
}

type DataCenter struct {
	sync.RWMutex
	Name      string
	Racks     map[string]*Rack
	Resource  pb.ComputeResource
	Allocated pb.ComputeResource
}

type Topology struct {
	Resource  pb.ComputeResource
	Allocated pb.ComputeResource
	sync.RWMutex
	DataCenters map[string]*DataCenter
}

func NewTopology() *Topology { _ = "STUB: not implemented"; return nil }

func NewDataCenter(name string) *DataCenter { _ = "STUB: not implemented"; return nil }

func NewRack(name string) *Rack { _ = "STUB: not implemented"; return nil }

func (tp *Topology) GetDataCenter(name string) (*DataCenter, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (dc *DataCenter) GetRack(name string) (*Rack, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (rack *Rack) GetAgent(name string) (*AgentInformation, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (tp *Topology) AddDataCenter(dc *DataCenter) { _ = "STUB: not implemented"; return }

func (tp *Topology) GetDataCenters() map[string]*DataCenter { _ = "STUB: not implemented"; return nil }

func (dc *DataCenter) GetRacks() (ret []*Rack) { _ = "STUB: not implemented"; return nil }

func (dc *DataCenter) AddRack(rack *Rack) { _ = "STUB: not implemented"; return }

func (rack *Rack) AddAgent(a *AgentInformation) { _ = "STUB: not implemented"; return }

func (rack *Rack) DropAgent(location *pb.Location) { _ = "STUB: not implemented"; return }

func (rack *Rack) GetAgents() (ret []*AgentInformation) { _ = "STUB: not implemented"; return nil }
