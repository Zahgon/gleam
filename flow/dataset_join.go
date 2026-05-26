package flow

// Join joins two datasets by the key.
func (d *Dataset) Join(name string, other *Dataset, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

// TODO use actual key fields instead of Field(1)

func (d *Dataset) JoinByKey(name string, other *Dataset) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) LeftOuterJoin(name string, other *Dataset, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) LeftOuterJoinByKey(name string, other *Dataset) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) RightOuterJoin(name string, other *Dataset, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) RightOuterJoinByKey(name string, other *Dataset) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) DoJoin(name string, other *Dataset, leftOuter, rightOuter bool, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

// JoinPartitionedSorted Join multiple datasets that are sharded by the same key, and locally sorted within the shard
func (this *Dataset) JoinPartitionedSorted(name string, that *Dataset, sortOption *SortOption,
	isLeftOuterJoin, isRightOuterJoin bool) *Dataset {
	_ = "STUB: not implemented"
	return nil
}
