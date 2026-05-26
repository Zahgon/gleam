package kafka

import (
	"github.com/chrislusf/gleam/flow"
)

type KafkaSource struct {
	Brokers        []string
	Group          string
	Topic          string
	TimeoutSeconds int

	prefix string
}

// Generate generates data shard info,
// partitions them via round robin,
// and reads each shard on each executor
func (s *KafkaSource) Generate(f *flow.Flow) *flow.Dataset { _ = "STUB: not implemented"; return nil }

func (s *KafkaSource) fetchPartitionIds() ([]int32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the partition ids for a topic

func (s *KafkaSource) genShardInfos(f *flow.Flow, partitionIds []int32) *flow.Dataset {
	_ = "STUB: not implemented"
	return nil
}
