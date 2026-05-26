package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
	"github.com/chrislusf/gleam/util"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetLocalTop() != nil {
			return NewLocalTop(
				int(m.GetLocalTop().GetN()),
				toOrderBys(m.GetLocalTop().GetOrderBys()),
			)
		}
		return nil
	})
}

type LocalTop struct {
	n        int
	orderBys []OrderBy
}

func NewLocalTop(n int, orderBys []OrderBy) *LocalTop { _ = "STUB: not implemented"; return nil }

func (b *LocalTop) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *LocalTop) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *LocalTop) SerializeToCommand() *pb.Instruction { _ = "STUB: not implemented"; return nil }

func (b *LocalTop) GetMemoryCostInMB(partitionSize int64) int64 {
	_ = "STUB: not implemented"

	// DoLocalTop streamingly compare and get the top n items
	return 0
}

func DoLocalTop(reader io.Reader, writer io.Writer, n int, orderBys []OrderBy, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

// read data out of the priority queue

func newMinQueueOfPairs(orderBys []OrderBy) *util.PriorityQueue {
	_ = "STUB: not implemented"
	return nil
}
