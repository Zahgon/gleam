package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetMergeTo() != nil {
			return NewMergeTo()
		}
		return nil
	})
}

type MergeTo struct{}

func NewMergeTo() *MergeTo { _ = "STUB: not implemented"; return nil }

func (b *MergeTo) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *MergeTo) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *MergeTo) SerializeToCommand() *pb.Instruction { _ = "STUB: not implemented"; return nil }

func (b *MergeTo) GetMemoryCostInMB(partitionSize int64) int64 { _ = "STUB: not implemented"; return 0 }

func DoMergeTo(readers []io.Reader, writer io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	// enqueue one item to the pq from each channel
	return nil
}
