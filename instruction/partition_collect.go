package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetCollectPartitions() != nil {
			return NewCollectPartitions()
		}
		return nil
	})
}

type CollectPartitions struct {
}

func NewCollectPartitions() *CollectPartitions { _ = "STUB: not implemented"; return nil }

func (b *CollectPartitions) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *CollectPartitions) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *CollectPartitions) SerializeToCommand() *pb.Instruction {
	_ = "STUB: not implemented"
	return nil
}

func (b *CollectPartitions) GetMemoryCostInMB(partitionSize int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func DoCollectPartitions(readers []io.Reader, writer io.Writer, stats *pb.InstructionStat) (err error) {
	_ = "STUB: not implemented"
	return nil
}
