package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetSelect() != nil {
			return NewSelect(
				toInts(m.GetSelect().GetKeyIndexes()),
				toInts(m.GetSelect().GetValueIndexes()),
			)
		}
		return nil
	})
}

type Select struct {
	keyIndexes   []int
	valueIndexes []int
}

func NewSelect(keyIndexes, valueIndexes []int) *Select { _ = "STUB: not implemented"; return nil }

func (b *Select) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *Select) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Select) SerializeToCommand() *pb.Instruction { _ = "STUB: not implemented"; return nil }

func (b *Select) GetMemoryCostInMB(partitionSize int64) int64 {
	_ = "STUB: not implemented"

	// DoSelect projects the fields
	return 0
}

func DoSelect(reader io.Reader, writer io.Writer, keyIndexes, valueIndexes []int, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func joinInts(a []int, sep string) string { _ = "STUB: not implemented"; return "" }
