package master

import (
	"github.com/chrislusf/gleam/pb"
)

func (tp *Topology) allocateDataCenter(requests []*pb.ComputeResource) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (tp *Topology) allocateServersOnRack(dc *DataCenter, rack *Rack, requests []*pb.ComputeResource) (
	allocated []*pb.Allocation, remainingRequests []*pb.ComputeResource) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fmt.Printf("available %v, requested %v\n", available, request.GetMemoryMb())

func (tp *Topology) findServers(dc *DataCenter, requests []*pb.ComputeResource) (ret []*pb.Allocation) {
	_ = "STUB: not implemented"

	// sort racks by unallocated resources
	return nil
}

type byAvailableResources []*Rack

func (s byAvailableResources) Len() int           { _ = "STUB: not implemented"; return 0 }
func (s byAvailableResources) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (s byAvailableResources) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type byRequestedResources []*pb.ComputeResource

func (s byRequestedResources) Len() int           { _ = "STUB: not implemented"; return 0 }
func (s byRequestedResources) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (s byRequestedResources) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
