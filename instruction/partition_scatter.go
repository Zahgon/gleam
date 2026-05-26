package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetScatterPartitions() != nil {
			return NewScatterPartitions(
				toInts(m.GetScatterPartitions().GetIndexes()),
			)
		}
		return nil
	})
}

type ScatterPartitions struct {
	indexes []int
}

func NewScatterPartitions(indexes []int) *ScatterPartitions { _ = "STUB: not implemented"; return nil }

func (b *ScatterPartitions) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *ScatterPartitions) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *ScatterPartitions) SerializeToCommand() *pb.Instruction {
	_ = "STUB: not implemented"
	return nil
}

func (b *ScatterPartitions) GetMemoryCostInMB(partitionSize int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func DoScatterPartitions(reader io.Reader, writers []io.Writer, indexes []int, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}
