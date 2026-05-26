package cassandra

import (
	"github.com/chrislusf/gleam/flow"
)

type CassandraSource struct {
	hosts            string
	Concurrency      int
	ShardCount       int
	LimitInEachShard int
	TimeoutSeconds   int

	prefix string

	selectClause string
	keyspace     string
	table        string
	whereClause  string
}

// Generate generates data shard info,
// partitions them via round robin,
// and reads each shard on each executor
func (s *CassandraSource) Generate(f *flow.Flow) *flow.Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (s *CassandraSource) genShardInfos(f *flow.Flow) *flow.Dataset {
	_ = "STUB: not implemented"
	return nil
}

// find out the partition keys

// println("driver Connected to", s.hosts, "keyspace", s.keyspace)

// divide by token range
