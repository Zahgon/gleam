package util

import (
	"io"
)

func WriteEOFMessage(writer io.Writer) (err error) { _ = "STUB: not implemented"; return nil }

func WriteMessage(writer io.Writer, m []byte) (err error) { _ = "STUB: not implemented"; return nil }

type BufferedMessageWriter struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}

func NewBufferedMessageWriter(w io.Writer, size int) *BufferedMessageWriter {
	_ = "STUB: not implemented"
	return nil
}

func (b *BufferedMessageWriter) Available() int { _ = "STUB: not implemented"; return 0 }
func (b *BufferedMessageWriter) Buffered() int  { _ = "STUB: not implemented"; return 0 }

func (b *BufferedMessageWriter) WriteMessage(m []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Large write, empty buffer.
// Write directly from p to avoid copy.

func (b *BufferedMessageWriter) Flush() error { _ = "STUB: not implemented"; return nil }

func (b *BufferedMessageWriter) flush() error { _ = "STUB: not implemented"; return nil }
