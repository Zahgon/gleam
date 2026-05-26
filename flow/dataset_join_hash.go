package flow

// HashJoin joins two datasets by putting the smaller dataset in memory on all
// executors and streams through the bigger dataset.
func (bigger *Dataset) HashJoin(name string, smaller *Dataset, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (this *Dataset) LocalHashAndJoinWith(name string, that *Dataset, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

// Broadcast replicates itself to all shards.
func (d *Dataset) Broadcast(name string, shardCount int) *Dataset {
	_ = "STUB: not implemented"
	return nil
}
