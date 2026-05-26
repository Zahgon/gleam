package util

import (
	"context"
	"io"
	"os/exec"
	"sync"

	"github.com/chrislusf/gleam/pb"
)

// all data passing through pipe are all (size, msgpack_encoded) tuples
// The input and output should all be this msgpack format.
// Only the stdin and stdout of Pipe() is line based text.
func Execute(ctx context.Context, executeWaitGroup *sync.WaitGroup, stat *pb.InstructionStat,
	name string, command *exec.Cmd,
	reader io.Reader, writer io.Writer, prevIsPipe, isPipe bool, closeOutput bool,
	errWriter io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// println("step", name, "input is lines->lines")

// println("step", name, "input is msgpack->msgpack")

// println("step", name, "input is msgpack->lines")

// println("step", name, "input is lines->msgpack")

// println(name, "starting...", strings.Join(command.Args, ","))

// fmt.Printf("%s Command is waiting..\n", name)

// defer fmt.Printf("%s Command is finished.\n", name)
