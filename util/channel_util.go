package util

import (
	"io"
	"sync"

	"github.com/chrislusf/gleam/pb"
)

/*
On the wire Message format, pipe:
  32 bits byte length
  []byte encoded in msgpack format

Channel Message format:
  []byte
    consecutive sections of []byte, each section is an object encoded in msgpack format

	This is not actually an array object,
	but just a consecutive list of encoded bytes for each object,
	because msgpack can sequentially decode the objects

When used by Shell scripts:
  from input channel:
    decode the msgpack-encoded []byte into strings that's tab and '\n' separated
	and feed into the shell script
  to output channel:
    encode the tab and '\n' separated lines into msgpack-format []byte
	and feed into the output channel

When used by Lua scripts:
  from input channel:
    decode the msgpack-encoded []byte into array of objects
	and pass these objects as function parameters
  to output channel:
    encode returned objects as an array of objects, into msgpack encoded []byte
	and feed into the output channel

Output Message format:
  decoded objects

Lua scripts need to decode the input and encode the output in msgpack format.
Go code also need to decode the input to "see" the data, e.g. Sort(),
and encode the output, e.g. Source().

Shell scripts via Pipe should see clear data, so the
*/

const (
	BUFFER_SIZE = 1024 * 512
)

// setup asynchronously to merge multiple channels into one channel
func CopyMultipleReaders(readers []io.Reader, writer io.Writer) (inCounter int64, outCounter int64, e error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func ReaderToChannel(wg *sync.WaitGroup, name string, readCloser io.ReadCloser, writer io.WriteCloser, closeOutput bool, errorOutput io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// println("reader", name, "copied", n, "bytes.")

func ChannelToWriter(wg *sync.WaitGroup, name string, reader io.Reader, writer io.WriteCloser, errorOutput io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// println("writer", name, "moved", n, "bytes.")

func LineReaderToChannel(wg *sync.WaitGroup, stat *pb.InstructionStat, name string, reader io.Reader, ch io.WriteCloser, closeOutput bool, errorOutput io.Writer) {
	_ = "STUB: not implemented"
	return
}

// fmt.Printf("%s>line input: %s\n", name, scanner.Text())

// seems the program could have ended when reading the output.

func ConvertLineReaderToRowReader(lineReader io.Reader, name string, errorOutput io.Writer) (rowReader io.Reader) {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

// fmt.Fprintf(errorOutput, "> %s\n", m)

func ChannelToLineWriter(wg *sync.WaitGroup, stat *pb.InstructionStat, name string, reader io.Reader, writer io.WriteCloser, errorOutput io.Writer) {
	_ = "STUB: not implemented"
	return
}

func copyBuffer(dst io.Writer, src io.Reader, written *int64) (err error) {
	_ = "STUB: not implemented"

	// this is a third buffer, additional to dst buffer and src buffer
	return nil
}
