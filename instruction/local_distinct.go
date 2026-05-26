package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetLocalDistinct() != nil {
			return NewLocalDistinct(
				toOrderBys(m.GetLocalDistinct().GetOrderBys()),
			)
		}
		return nil
	})
}

type LocalDistinct struct {
	orderBys []OrderBy
}

func NewLocalDistinct(orderBys []OrderBy) *LocalDistinct { _ = "STUB: not implemented"; return nil }

func (b *LocalDistinct) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *LocalDistinct) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *LocalDistinct) SerializeToCommand() *pb.Instruction { _ = "STUB: not implemented"; return nil }

func (b *LocalDistinct) GetMemoryCostInMB(partitionSize int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func DoLocalDistinct(reader io.Reader, writer io.Writer, orderBys []OrderBy, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

// write the row if key is different
