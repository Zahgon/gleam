package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetUnion() != nil {
			return NewUnion(m.GetUnion().GetIsParallel())
		}
		return nil
	})
}

type Union struct {
	isParallel bool
}

func NewUnion(isParallel bool) *Union { _ = "STUB: not implemented"; return nil }

func (b *Union) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *Union) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Union) SerializeToCommand() *pb.Instruction { _ = "STUB: not implemented"; return nil }

func (b *Union) GetMemoryCostInMB(partitionSize int64) int64 { _ = "STUB: not implemented"; return 0 }

func DoUnion(readers []io.Reader, writer io.Writer, isParallel bool,
	stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}
