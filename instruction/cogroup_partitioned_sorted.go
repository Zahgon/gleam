package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetCoGroupPartitionedSorted() != nil {
			return NewCoGroupPartitionedSorted(
				toInts(m.GetCoGroupPartitionedSorted().GetIndexes()),
			)
		}
		return nil
	})
}

type CoGroupPartitionedSorted struct {
	indexes []int
}

func NewCoGroupPartitionedSorted(indexes []int) *CoGroupPartitionedSorted {
	_ = "STUB: not implemented"
	return nil
}

func (b *CoGroupPartitionedSorted) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *CoGroupPartitionedSorted) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *CoGroupPartitionedSorted) SerializeToCommand() *pb.Instruction {
	_ = "STUB: not implemented"
	return nil
}

func (b *CoGroupPartitionedSorted) GetMemoryCostInMB(partitionSize int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func DoCoGroupPartitionedSorted(leftRawChan, rightRawChan io.Reader, writer io.Writer, indexes []int, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

// get first value from both channels
