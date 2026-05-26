package flow

// CoGroup joins two datasets by the key,
// Each result row becomes this format:
//
//	(key, []left_rows, []right_rows)
func (d *Dataset) CoGroup(name string, other *Dataset, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

// this should not happen, but just in case

// CoGroupPartitionedSorted joins 2 datasets that are sharded
// by the same key and already locally sorted within each shard.
func (this *Dataset) CoGroupPartitionedSorted(name string, that *Dataset, indexes []int) (ret *Dataset) {
	_ = "STUB: not implemented"
	return nil
}
