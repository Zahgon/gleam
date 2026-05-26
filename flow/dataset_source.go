package flow

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

type Sourcer interface {
	Generate(*Flow) *Dataset
}

// Read accepts a function to read data into the flow, creating a new dataset.
// This allows custom complicated pre-built logic for new data sources.
func (fc *Flow) Read(s Sourcer) (ret *Dataset) { _ = "STUB: not implemented"; return nil }

// Listen receives textual inputs via a socket.
// Multiple parameters are separated via tab.
func (fc *Flow) Listen(network, address string) (ret *Dataset) {
	_ = "STUB: not implemented"
	return nil
}

// Source produces data feeding into the flow.
// Function f writes to this writer.
// The written bytes should be MsgPack encoded []byte.
// Use util.EncodeRow(...) to encode the data before sending to this channel
func (fc *Flow) Source(name string, f func(io.Writer, *pb.InstructionStat) error) (ret *Dataset) {
	_ = "STUB: not implemented"
	return nil
}

// Channel accepts a channel to feed into the flow.
func (fc *Flow) Channel(ch chan interface{}) (ret *Dataset) { _ = "STUB: not implemented"; return nil }

// Bytes begins a flow with an [][]byte
func (fc *Flow) Bytes(slice [][]byte) (ret *Dataset) { _ = "STUB: not implemented"; return nil }

// println("sent []byte of size:", len(data), string(data))

// Strings begins a flow with an []string
func (fc *Flow) Strings(lines []string) (ret *Dataset) { _ = "STUB: not implemented"; return nil }

// Ints begins a flow with an []int
func (fc *Flow) Ints(numbers []int) (ret *Dataset) { _ = "STUB: not implemented"; return nil }

// Slices begins a flow with an [][]interface{}
func (fc *Flow) Slices(slices [][]interface{}) (ret *Dataset) {
	_ = "STUB: not implemented"
	return nil
}
