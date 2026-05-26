package util

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

func ListFiles(dir string, pattern string) (fileNames []string) {
	_ = "STUB: not implemented"
	return nil
}

func Range(from, to int) func(io.Writer, *pb.InstructionStat) error {
	_ = "STUB: not implemented"
	return nil
}
