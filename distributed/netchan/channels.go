// Package netchan creates network channels. The network channels are managed by
// glow agent.
package netchan

import (
	"context"
	"io"
	"sync"
)

func DialReadChannel(ctx context.Context, wg *sync.WaitGroup, readerName string, address string, channelName string, onDisk bool, outChan io.WriteCloser) error {
	_ = "STUB: not implemented"
	return nil
}

func DialWriteChannel(ctx context.Context, wg *sync.WaitGroup, writerName string, address string, channelName string, onDisk bool, inChan io.Reader, readerCount int) error {
	_ = "STUB: not implemented"
	return nil
}
