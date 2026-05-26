package file

import (
	"encoding/gob"

	"github.com/chrislusf/gleam/gio"
)

type FileShardInfo struct {
	Config    map[string]string
	FileName  string
	FileType  string
	HasHeader bool
	Fields    []string
}

var (
	registeredMapperReadShard = gio.RegisterMapper(readShard)
)

func init() {
	gob.Register(FileShardInfo{})
}

func readShard(row []interface{}) error { _ = "STUB: not implemented"; return nil }

func (ds *FileShardInfo) ReadSplit() error {
	_ = "STUB: not implemented"

	// println("opening file", ds.FileName)
	return nil
}

func decodeShardInfo(encodedShardInfo []byte) *FileShardInfo { _ = "STUB: not implemented"; return nil }

func encodeShardInfo(shardInfo *FileShardInfo) []byte { _ = "STUB: not implemented"; return nil }
