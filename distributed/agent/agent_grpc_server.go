package agent

import (
	"net"

	"context"

	"github.com/chrislusf/gleam/pb"
)

func (as *AgentServer) serveGrpc(listener net.Listener) { _ = "STUB: not implemented"; return }

func (as *AgentServer) SendFileResource(stream pb.GleamAgent_SendFileResourceServer) error {
	_ = "STUB: not implemented"
	return nil
}

// ack

// Cleanup remove all files related to a particular flow
func (as *AgentServer) Cleanup(ctx context.Context, cleanupRequest *pb.CleanupRequest) (*pb.CleanupResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Execute executes a request and stream stdout and stderr back
func (as *AgentServer) Execute(request *pb.ExecutionRequest, stream pb.GleamAgent_ExecuteServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Collect stat from "gleam execute" process
func (as *AgentServer) CollectExecutionStatistics(stream pb.GleamAgent_CollectExecutionStatisticsServer) error {
	_ = "STUB: not implemented"
	return nil
}

// fmt.Printf("received stats: %+v\n", stats)

// Delete deletes a particular dataset shard
func (as *AgentServer) Delete(ctx context.Context, deleteRequest *pb.DeleteDatasetShardRequest) (*pb.DeleteDatasetShardResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (as *AgentServer) plusAllocated(allocated pb.ComputeResource) {
	_ = "STUB: not implemented"
	return
}

func (as *AgentServer) minusAllocated(allocated pb.ComputeResource) {
	_ = "STUB: not implemented"
	return
}
