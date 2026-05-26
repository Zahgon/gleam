package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetLocalHashAndJoinWith() != nil {
			return NewLocalHashAndJoinWith(
				toInts(m.GetLocalHashAndJoinWith().GetIndexes()),
			)
		}
		return nil
	})
}

type LocalHashAndJoinWith struct {
	indexes []int
}

func NewLocalHashAndJoinWith(indexes []int) *LocalHashAndJoinWith {
	_ = "STUB: not implemented"
	return nil
}

func (b *LocalHashAndJoinWith) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *LocalHashAndJoinWith) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *LocalHashAndJoinWith) SerializeToCommand() *pb.Instruction {
	_ = "STUB: not implemented"
	return nil
}

func (b *LocalHashAndJoinWith) GetMemoryCostInMB(partitionSize int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func DoLocalHashAndJoinWith(leftReader, rightReader io.Reader, writer io.Writer, indexes []int, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

// write the row if key is different

// write the row if key is different
