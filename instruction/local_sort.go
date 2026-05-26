package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
	"github.com/chrislusf/gleam/util"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetLocalSort() != nil {
			return NewLocalSort(
				toOrderBys(m.GetLocalSort().GetOrderBys()),
				int(m.GetMemoryInMB()),
			)
		}
		return nil
	})
}

type LocalSort struct {
	orderBys   []OrderBy
	memoryInMB int
}

func NewLocalSort(orderBys []OrderBy, memoryInMB int) *LocalSort {
	_ = "STUB: not implemented"
	return nil
}

func (b *LocalSort) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *LocalSort) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *LocalSort) SerializeToCommand() *pb.Instruction { _ = "STUB: not implemented"; return nil }

func (b *LocalSort) GetMemoryCostInMB(partitionSize int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func DoLocalSort(reader io.Reader, writer io.Writer, orderBys []OrderBy, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

// println("sorted key", kv.(pair).keys[0].(string))

func getIndexesFromOrderBys(orderBys []OrderBy) (indexes []int) {
	_ = "STUB: not implemented"
	return nil
}

func lessThan(orderBys []OrderBy, x, y *util.Row) bool { _ = "STUB: not implemented"; return false }

func getIndexes(storedValues []int) (indexes []int32) { _ = "STUB: not implemented"; return nil }

func getOrderBys(storedValues []OrderBy) (orderBys []*pb.OrderBy) {
	_ = "STUB: not implemented"
	return nil
}
