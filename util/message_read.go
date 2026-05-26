package util

import (
	"io"
	"math"
)

type MessageControl int32

const (
	MessageControlEOF = MessageControl(math.MinInt32)
)

// message contains 3 kinds of data with the formats:
//   the first 4 bytes is int32 flag
//   if flag > 0
//     actual data row bytes with length = flag
//   else if flag == math.MinInt32
//     end of partition: EOF(math.MinInt32)
//   else
//     meta data bytes with length = - flag

// TakeTsv Reads and processes TSV lines.
// If count is less than 0, all lines are processed.
func TakeTsv(reader io.Reader, count int, f func([]string) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// TakeMessage Reads and processes MessagePack encoded messages.
// If count is less than 0, all lines are processed.
func TakeMessage(reader io.Reader, count int, f func([]byte) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ProcessMessage Reads and processes MessagePack encoded messages until EOF
func ProcessMessage(reader io.Reader, f func([]byte) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ReadMessage reads out the []byte for one message
func ReadMessage(reader io.Reader) (m []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
