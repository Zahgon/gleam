package executor

import (
	"sync"

	"github.com/chrislusf/gleam/pb"
)

func (exe *Executor) statusHeartbeat(wg *sync.WaitGroup, finishedChan chan bool) {
	_ = "STUB: not implemented"
	return
}

func (exe *Executor) reportStatus() { _ = "STUB: not implemented"; return }

// defer stream.CloseSend()

func withClient(server string, fn func(client pb.GleamAgentClient) error) error {
	_ = "STUB: not implemented"
	return nil
}
