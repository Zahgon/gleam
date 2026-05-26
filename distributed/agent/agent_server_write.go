package agent

import (
	"io"
)

func (as *AgentServer) handleLocalWriteConnection(reader io.Reader, writerName, channelName string, readerCount int) {
	_ = "STUB: not implemented"
	return
}

// println("agent recv eof:", string(message.Bytes()))

// println("agent recv:", string(message.Bytes()))
