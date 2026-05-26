package util

import (
	"io"
)

type Piper struct {
	Reader  *io.PipeReader
	Writer  *io.PipeWriter
	Counter int64
	Error   error
}

func NewPiper() *Piper { _ = "STUB: not implemented"; return nil }
