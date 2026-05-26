package flow

func (d *Dataset) RoundRobin(name string, n int) *Dataset { _ = "STUB: not implemented"; return nil }

// hash data or by data key, return a new dataset
// This is divided into 2 steps:
// 1. Each record is sharded to a local shard
// 2. The destination shard will collect its child shards and merge into one
func (d *Dataset) Partition(name string, shard int, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) PartitionByKey(name string, shard int) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) partition_scatter(name string, shardCount int, indexes []int) (ret *Dataset) {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) partition_collect(name string, shardCount int, indexes []int) (ret *Dataset) {
	_ = "STUB: not implemented"
	return nil
}

func intArrayEquals(a []int, b []int) bool { _ = "STUB: not implemented"; return false }
