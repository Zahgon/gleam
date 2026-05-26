package util

import (
	"io"
)

// BufWrites ensures all writers are bufio.Writer
// For any bufio.Writer created here, flush it before returning.
func BufWrites(rawWriters []io.Writer, function func([]io.Writer)) {
	_ = "STUB: not implemented"
	return
}
