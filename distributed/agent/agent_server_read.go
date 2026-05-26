package agent

import (
	"net"
)

func (as *AgentServer) handleReadConnection(conn net.Conn, readerName, channelName string) {
	_ = "STUB: not implemented"
	return
}

// loop for every read

// connection is closed

// println("got problem reading", channelName, offset, err.Error())

// println("reading", channelName, offset, "size:", size)

// connection is closed
