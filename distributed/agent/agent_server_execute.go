package agent

import (
	"io"
	"os/exec"
	"sync"

	"github.com/chrislusf/gleam/pb"
)

func (as *AgentServer) executeCommand(
	stream pb.GleamAgent_ExecuteServer,
	startRequest *pb.ExecutionRequest,
	dir string,
	statChan chan *pb.ExecutionStat,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// start the command

// Note: don't use exec.CommandContext here.
// The executor process will be killed by SIGKILL and all of its child process will be left behind if
// the context passed to exec.CommandContext is canceled.
// Instead, we send a SIGTERM to the executor process when stream.Context is canceled and give
// the executor a chance to reap its children.

// msg.Env = startRequest.Envs

// send instruction set to executor

// wait for finish

func streamOutput(wg *sync.WaitGroup, stream pb.GleamAgent_ExecuteServer, reader io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func streamError(wg *sync.WaitGroup, stream pb.GleamAgent_ExecuteServer, reader io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func streamPulse(wg *sync.WaitGroup,
	stopChan chan bool,
	statChan chan *pb.ExecutionStat,
	stream pb.GleamAgent_ExecuteServer) error {
	_ = "STUB: not implemented"
	return nil
}

func sendExitStats(stream pb.GleamAgent_ExecuteServer, cmd *exec.Cmd) error {
	_ = "STUB: not implemented"
	return nil
}
