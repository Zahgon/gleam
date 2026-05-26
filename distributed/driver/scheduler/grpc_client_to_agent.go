package scheduler

import (
	"context"

	"github.com/chrislusf/gleam/distributed/resource"
	"github.com/chrislusf/gleam/pb"
)

func sendRelatedFile(ctx context.Context, client pb.GleamAgentClient, flowHashCode uint32, relatedFile resource.FileResource) error {
	_ = "STUB: not implemented"
	return nil
}

// receive ack

func sendExecutionRequest(ctx context.Context,
	_ *pb.FlowExecutionStatus_TaskGroup,
	executionStatus *pb.FlowExecutionStatus_TaskGroup_Execution,
	server string, request *pb.ExecutionRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// stream.CloseSend()

// log.Printf("%s %v>  UserTime: %2.2fs SystemTime: %2.2fs\n", server, request.InstructionSet.Name, response.GetSystemTime(), response.GetUserTime())

// merge existing stats with incoming stats
func mergeStats(a, b []*pb.InstructionStat) (ret []*pb.InstructionStat) {
	_ = "STUB: not implemented"
	return nil
}

func sendDeleteRequest(server string, request *pb.DeleteDatasetShardRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func SendCleanupRequest(server string, request *pb.CleanupRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func withClient(server string, fn func(client pb.GleamAgentClient) error) error {
	_ = "STUB: not implemented"
	return nil
}
