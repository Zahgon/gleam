// Package agent runs on servers with computing resources, and executes
// tasks sent by driver.
package agent

import (
	"net"
	"sync"

	"github.com/chrislusf/gleam/pb"
)

type AgentServerOption struct {
	Master       *string
	Host         *string
	Port         *int32
	Dir          *string
	DataCenter   *string
	Rack         *string
	MaxExecutor  *int32
	MemoryMB     *int64
	CPULevel     *int32
	CleanRestart *bool
}

type AgentServer struct {
	Option                  *AgentServerOption
	Master                  string
	computeResource         *pb.ComputeResource
	allocatedResource       *pb.ComputeResource
	allocatedHasChanges     chan struct{}
	allocatedResourceLock   sync.Mutex
	storageBackend          *LocalDatasetShardsManager
	inMemoryChannels        *LocalDatasetShardsManagerInMemory
	receiveFileResourceLock sync.Mutex
}

func RunAgentServer(option *AgentServerOption) { _ = "STUB: not implemented"; return }

// println("removing old dat file:", name)

// Run starts the heartbeating to master and starts accepting requests.
func (as *AgentServer) serveTcp(listener net.Listener) {
	_ = "STUB: not implemented"

	// Listen for an incoming connection.
	return
}

// Handle connections in a new goroutine.

func (r *AgentServer) handleRequest(conn net.Conn) { _ = "STUB: not implemented"; return }

func (as *AgentServer) handleCommandConnection(conn net.Conn,
	command *pb.ControlMessage) {
	_ = "STUB: not implemented"
	return
}
