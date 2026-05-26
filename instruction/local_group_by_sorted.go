package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetLocalGroupBySorted() != nil {
			return NewLocalGroupBySorted(
				toInts(m.GetLocalGroupBySorted().GetIndexes()),
			)
		}
		return nil
	})
}

type LocalGroupBySorted struct {
	indexes []int
}

func NewLocalGroupBySorted(indexes []int) *LocalGroupBySorted {
	_ = "STUB: not implemented"
	return nil
}

func (b *LocalGroupBySorted) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *LocalGroupBySorted) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *LocalGroupBySorted) SerializeToCommand() *pb.Instruction {
	_ = "STUB: not implemented"
	return nil
}

func (b *LocalGroupBySorted) GetMemoryCostInMB(partitionSize int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func DoLocalGroupBySorted(reader io.Reader, writer io.Writer,
	indexes []int, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

// write prev row if key is different
