package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetLocalLimit() != nil {
			return NewLocalLimit(
				int(m.GetLocalLimit().GetN()),
				int(m.GetLocalLimit().GetOffset()),
			)
		}
		return nil
	})
}

type LocalLimit struct {
	n      int
	offset int
}

func NewLocalLimit(n int, offset int) *LocalLimit { _ = "STUB: not implemented"; return nil }

func (b *LocalLimit) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *LocalLimit) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *LocalLimit) SerializeToCommand() *pb.Instruction { _ = "STUB: not implemented"; return nil }

func (b *LocalLimit) GetMemoryCostInMB(partitionSize int64) int64 {
	_ = "STUB: not implemented"

	// DoLocalLimit streamingly get the n items starting from offset
	return 0
}

func DoLocalLimit(reader io.Reader, writer io.Writer, n int, offset int, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}
