package orc

import (
	"github.com/chrislusf/gleam/util"
	"github.com/scritchley/orc"
)

type OrcFileReader struct {
	reader     *orc.Reader
	cursor     *orc.Cursor
	fieldNames []string
}

// TODO predicate pushdown
func New(reader orc.SizedReaderAt) (*OrcFileReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *OrcFileReader) Select(fields []string) *OrcFileReader {
	_ = "STUB: not implemented"
	return nil
}

func (r *OrcFileReader) ReadHeader() (fieldNames []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *OrcFileReader) Read() (row *util.Row, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Iterate over each row in the stripe.
