package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetPipeAsArgs() != nil {
			return NewPipeAsArgs(m.GetPipeAsArgs().GetCode())
		}
		return nil
	})
}

type PipeAsArgs struct {
	code string
}

func NewPipeAsArgs(code string) *PipeAsArgs { _ = "STUB: not implemented"; return nil }

func (b *PipeAsArgs) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *PipeAsArgs) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *PipeAsArgs) SerializeToCommand() *pb.Instruction { _ = "STUB: not implemented"; return nil }

func (b *PipeAsArgs) GetMemoryCostInMB(partitionSize int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func DoPipeAsArgs(reader io.Reader, writer io.Writer, code string, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

// feed parts as input to the code

// println("pipeAsArgs command:", actualCode)

// write output to writer

//wg.Wait()
