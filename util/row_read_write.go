package util

import (
	"io"
)

// WriteTo encode and write a row of data to the writer
func (row Row) WriteTo(writer io.Writer) (err error) { _ = "STUB: not implemented"; return nil }

// ReadRow read and decode one row of data
func ReadRow(reader io.Reader) (row *Row, err error) { _ = "STUB: not implemented"; return nil, nil }

// EncodeRow encode one row of data to a blob
func encodeRow(row Row) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// EncodeKeys encode keys to a blob, for comparing or sorting
		nil
}

func EncodeKeys(anyObject ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DecodeRow decodes one row of data from a blob
func DecodeRow(encodedBytes []byte) (*Row, error) { _ = "STUB: not implemented"; return nil, nil }

// ProcessRow Reads and processes rows until EOF
func ProcessRow(reader io.Reader, indexes []int, f func(*Row) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// read the row
