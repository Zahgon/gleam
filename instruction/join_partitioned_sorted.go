package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetJoinPartitionedSorted() != nil {
			return NewJoinPartitionedSorted(
				m.GetJoinPartitionedSorted().GetIsLeftOuterJoin(),
				m.GetJoinPartitionedSorted().GetIsRightOuterJoin(),
				toInts(m.GetJoinPartitionedSorted().GetIndexes()),
			)
		}
		return nil
	})
}

type JoinPartitionedSorted struct {
	isLeftOuterJoin  bool
	isRightOuterJoin bool
	indexes          []int
}

func NewJoinPartitionedSorted(isLeftOuterJoin bool, isRightOuterJoin bool, indexes []int) *JoinPartitionedSorted {
	_ = "STUB: not implemented"
	return nil
}

func (b *JoinPartitionedSorted) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *JoinPartitionedSorted) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *JoinPartitionedSorted) SerializeToCommand() *pb.Instruction {
	_ = "STUB: not implemented"
	return nil
}

func (b *JoinPartitionedSorted) GetMemoryCostInMB(partitionSize int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func DoJoinPartitionedSorted(leftRawChan, rightRawChan io.Reader, writer io.Writer, indexes []int,
	isLeftOuterJoin, isRightOuterJoin bool, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

// get first value from both channels

// left and right cartician join

func addNils(target []interface{}, nilCount int) []interface{} {
	_ = "STUB: not implemented"
	return nil
}
