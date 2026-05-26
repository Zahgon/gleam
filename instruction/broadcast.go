package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

func init() {
	InstructionRunner.Register(func(m *pb.Instruction) Instruction {
		if m.GetBroadcast() != nil {
			return NewBroadcast()
		}
		return nil
	})
}

type Broadcast struct {
}

func NewBroadcast() *Broadcast { _ = "STUB: not implemented"; return nil }

func (b *Broadcast) Name(prefix string) string { _ = "STUB: not implemented"; return "" }

func (b *Broadcast) Function() func(readers []io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Broadcast) SerializeToCommand() *pb.Instruction { _ = "STUB: not implemented"; return nil }

func (b *Broadcast) GetMemoryCostInMB(partitionSize int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func DoBroadcast(reader io.Reader, writers []io.Writer, stats *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}
