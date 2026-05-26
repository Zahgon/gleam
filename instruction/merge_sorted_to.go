package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
	"github.com/chrislusf/gleam/util"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetMergeSortedTo() != nil {
			return NewMergeSortedTo(
				toOrderBys(m.GetMergeSortedTo().GetOrderBys()),
			)
		}
		return nil
	})
}

type MergeSortedTo struct {
	orderBys []OrderBy
}

func NewMergeSortedTo(orderBys []OrderBy) *MergeSortedTo { _ = "STUB: not implemented"; return nil }

func (b *MergeSortedTo) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *MergeSortedTo) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *MergeSortedTo) SerializeToCommand() *pb.Instruction { _ = "STUB: not implemented"; return nil }

func (b *MergeSortedTo) GetMemoryCostInMB(partitionSize int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

type rowWithOriginalData struct {
	row       *util.Row
	originalK []interface{}
	originalV []interface{}
}

func newRowWithOriginalData(row *util.Row) *rowWithOriginalData {
	_ = "STUB: not implemented"
	return nil
}

func newMinQueueOfRowsWithOriginalData(orderBys []OrderBy) *util.PriorityQueue {
	_ = "STUB: not implemented"
	return nil
}

func DoMergeSortedTo(readers []io.Reader, writer io.Writer, orderBys []OrderBy, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

// enqueue one item to the pq from each channel
