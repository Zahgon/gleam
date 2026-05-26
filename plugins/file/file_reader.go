package file

import (
	"github.com/chrislusf/gleam/filesystem"
	"github.com/chrislusf/gleam/util"
)

type FileReader interface {
	Read() (row *util.Row, err error)
	ReadHeader() (fieldNames []string, err error)
}

func Csv(fileOrPattern string, partitionCount int) *FileSource {
	_ = "STUB: not implemented"
	return nil
}

func Txt(fileOrPattern string, partitionCount int) *FileSource {
	_ = "STUB: not implemented"
	return nil
}

func Tsv(fileOrPattern string, partitionCount int) *FileSource {
	_ = "STUB: not implemented"
	return nil
}

func Orc(fileOrPattern string, partitionCount int) *FileSource {
	_ = "STUB: not implemented"
	return nil
}

func Parquet(fileOrPattern string, partitionCount int) *FileSource {
	_ = "STUB: not implemented"
	return nil
}

func Zip(fileOrPattern string, partitionCount int) *FileSource {
	_ = "STUB: not implemented"
	return nil
}

func (ds *FileShardInfo) NewReader(vf filesystem.VirtualFile) (FileReader, error) {
	_ = "STUB: not implemented"
	// These formats require seeking, so they cannot be
	// sequentially read by a compress/* reader.
	return *new(FileReader), nil
}
