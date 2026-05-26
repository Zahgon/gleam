package cassandra

import (
	"encoding/gob"

	"github.com/chrislusf/gleam/gio"
)

type CassandraShardInfo struct {
	Hosts                 string
	StartToken, StopToken string
	PartitionKeys         []string
	TimeoutSeconds        int

	Select   string
	Keyspace string
	Table    string
	Where    string
	Limit    int
}

var (
	MapperReadShard = gio.RegisterMapper(readShard)
)

func init() {
	gob.Register(CassandraShardInfo{})
}

func readShard(row []interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *CassandraShardInfo) ReadSplit() error {
	_ = "STUB: not implemented"

	// println("hosts:", s.Hosts)
	return nil
}

// println("cql:", cql)

func decodeShardInfo(encodedShardInfo []byte) *CassandraShardInfo {
	_ = "STUB: not implemented"
	return nil
}

func encodeShardInfo(shardInfo *CassandraShardInfo) []byte { _ = "STUB: not implemented"; return nil }
