package util

import (
	"io"

	"github.com/chrislusf/gleam/pb"
)

// TsvPrintf reads TSV lines from reader,
// and formats according to a format specifier and writes to writer.
func TsvPrintf(writer io.Writer, reader io.Reader, format string) error {
	_ = "STUB: not implemented"
	return nil
}

// Fprintf reads MessagePack encoded messages from reader,
// and formats according to a format specifier and writes to writer.
func Fprintf(writer io.Writer, reader io.Reader, format string) error {
	_ = "STUB: not implemented"
	return nil
}

// fmt.Printf("chan input encoded: %s\n", string(encodedBytes))

// PrintDelimited Reads and formats MessagePack encoded messages
// with delimiter and lineSeparator.
func PrintDelimited(stat *pb.InstructionStat, reader io.Reader, writer io.Writer, delimiter string, lineSperator string) error {
	_ = "STUB: not implemented"
	return nil
}

// fmt.Printf("chan input encoded: %s\n", string(encodedBytes))

// fmt.Printf("> len=%d row:%s\n", len(decodedObjects), decodedObjects[0])

func fprintRow(writer io.Writer, base int, delimiter string, decodedObjects ...interface{}) (written int, err error) {
	_ = "STUB: not implemented"
	// fmt.Printf("chan input decoded: %v\n", decodedObjects)
	return 0, nil
}

// only string or []byte is allowed in piping. numbers or other types need to be converted to string
