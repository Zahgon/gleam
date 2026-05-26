package executor

import (
	"context"
	"io"
	"sync"

	"github.com/chrislusf/gleam/pb"
	"github.com/chrislusf/gleam/util"
)

type ExecutorOption struct {
	Dir          string
	AgentAddress string
	HashCode     uint32
}

type Executor struct {
	Option       *ExecutorOption
	instructions *pb.InstructionSet
	stats        []*pb.InstructionStat
	grpcAddress  string
}

func NewExecutor(option *ExecutorOption, instructions *pb.InstructionSet) *Executor {
	_ = "STUB: not implemented"
	return nil
}

func (exe *Executor) ExecuteInstructionSet() error {
	_ = "STUB: not implemented"

	// start a listener for stats
	return nil
}

//TODO pass in the context

// Calling cancel() will stop all the mappers and reducers.

// Wait for all the mappers and reducers to stop.
// If we don't wait here, the executor process may exit before the signal is
// passed to all of its children processes.

func setupReaders(ctx context.Context, wg *sync.WaitGroup, ioErrChan chan error,
	i *pb.Instruction, inPiper *util.Piper, isFirst bool) (readers []io.Reader) {
	_ = "STUB: not implemented"
	return nil
}

// println(i.GetName(), "connecting to", inputLocation.Address(), "to read", inputLocation.GetName())

func setupWriters(ctx context.Context, wg *sync.WaitGroup, ioErrChan chan error,
	i *pb.Instruction, outPiper *util.Piper, isLast bool, readerCount int) (writers []io.Writer) {
	_ = "STUB: not implemented"
	return nil
}

// println(i.GetName(), "connecting to", outputLocation.Address(), "to write", outputLocation.GetName(), "readerCount", readerCount)

func (exe *Executor) executeInstruction(ctx context.Context, wg *sync.WaitGroup,
	ioErrChan, exeErrChan chan error,
	inChan, outChan *util.Piper, prevIsPipe bool,
	is *pb.InstructionSet, i *pb.Instruction,
	isFirst, isLast bool, readerCount int, stat *pb.InstructionStat) {
	_ = "STUB: not implemented"
	return
}

// println(i.GetName(), "running error", err.Error())

//TODO add errChan to scripts also?

// println("starting", i.Name, "inChan", inChan, "outChan", outChan)

// println("args:", i.GetScript().Args[len(i.GetScript().Args)-1])

// fmt.Fprintf(os.Stderr, "starting %d %d: %v\n", i.StepId, i.TaskId, command.Args)
