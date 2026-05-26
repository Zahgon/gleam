package gio

import (
	"sync"

	"github.com/chrislusf/gleam/pb"
)

func (runner *gleamRunner) statusHeartbeat(wg *sync.WaitGroup, finishedChan chan bool) {
	_ = "STUB: not implemented"
	return
}

func (runner *gleamRunner) reportStatus() { _ = "STUB: not implemented"; return }

// defer stream.CloseSend()

func withClient(server string, fn func(client pb.GleamExecutorClient) error) error {
	_ = "STUB: not implemented"
	return nil
}
