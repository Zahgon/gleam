package parquet

import (
	"github.com/chrislusf/gleam/filesystem"
	"github.com/chrislusf/gleam/util"
	. "github.com/xitongsys/parquet-go/reader"
	. "github.com/xitongsys/parquet-go/source"
	. "github.com/xitongsys/parquet-go/types"
)

type PqFile struct {
	FileName string
	VF       filesystem.VirtualFile
}

func (self *PqFile) Create(name string) (ParquetFile, error) {
	_ = "STUB: not implemented"
	return *new(ParquetFile), nil
}

func (self *PqFile) Open(name string) (ParquetFile, error) {
	_ = "STUB: not implemented"
	return *new(ParquetFile), nil
}

func (self *PqFile) Seek(offset int64, pos int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (self *PqFile) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (self *PqFile) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (self *PqFile) Close() error { _ = "STUB: not implemented"; return nil }

type ParquetFileReader struct {
	pqReader *ParquetReader
	NumRows  int
	Cursor   int
}

func New(reader filesystem.VirtualFile, fileName string) *ParquetFileReader {
	_ = "STUB: not implemented"
	return nil
}

func (self *ParquetFileReader) ReadHeader() (fieldNames []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (self *ParquetFileReader) Read() (row *util.Row, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
