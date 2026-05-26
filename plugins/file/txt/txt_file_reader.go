package txt

import (
	"bufio"
	"io"

	"github.com/chrislusf/gleam/util"
)

type TxtFileReader struct {
	scanner *bufio.Scanner
}

func New(reader io.Reader) *TxtFileReader { _ = "STUB: not implemented"; return nil }

func (r *TxtFileReader) ReadHeader() (fieldNames []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TxtFileReader) Read() (row *util.Row, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
