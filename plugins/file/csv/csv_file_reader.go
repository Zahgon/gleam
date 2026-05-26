package csv

import (
	"io"

	"github.com/chrislusf/gleam/util"
)

type CsvFileReader struct {
	csvReader *Reader
}

func New(reader io.Reader) *CsvFileReader { _ = "STUB: not implemented"; return nil }

func (r *CsvFileReader) ReadHeader() (fieldNames []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *CsvFileReader) Read() (row *util.Row, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
