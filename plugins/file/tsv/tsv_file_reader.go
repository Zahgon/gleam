package tsv

import (
	"bufio"
	"io"

	"github.com/chrislusf/gleam/util"
)

type TsvFileReader struct {
	scanner *bufio.Scanner
}

func New(reader io.Reader) *TsvFileReader { _ = "STUB: not implemented"; return nil }

func (r *TsvFileReader) ReadHeader() (fieldNames []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TsvFileReader) Read() (row *util.Row, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TsvFileReader) readOneLine() (values []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
