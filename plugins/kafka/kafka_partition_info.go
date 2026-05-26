package kafka

import (
	"encoding/gob"

	"github.com/chrislusf/gleam/gio"
)

type KafkaPartitionInfo struct {
	Brokers        []string
	Topic          string
	Group          string
	TimeoutSeconds int
	PartitionId    int32
}

var (
	MapperReadShard = gio.RegisterMapper(readShard)
)

func init() {
	gob.Register(KafkaPartitionInfo{})
}

func readShard(row []interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *KafkaPartitionInfo) ReadSplit() error {
	_ = "STUB: not implemented"

	// println("brokers:", s.Brokers)
	return nil
}

func decodeShardInfo(encodedShardInfo []byte) *KafkaPartitionInfo {
	_ = "STUB: not implemented"
	return nil
}

func encodeShardInfo(shardInfo *KafkaPartitionInfo) []byte { _ = "STUB: not implemented"; return nil }
