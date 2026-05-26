package flow

type FlowHintOption func(c *FlowConfig)

type FlowConfig struct {
	OnDisk bool
}

// Hint adds hints to the flow.
func (d *Flow) Hint(options ...FlowHintOption) { _ = "STUB: not implemented"; return }

// GetTotalSize returns the total size in MB for the dataset.
// This is based on the given hint.
func (d *Dataset) GetTotalSize() int64 { _ = "STUB: not implemented"; return 0 }

// GetPartitionSize returns the size in MB for each partition of
// the dataset. This is based on the hinted total size divided by
// the number of partitions.
func (d *Dataset) GetPartitionSize() int64 { _ = "STUB: not implemented"; return 0 }

// GetIsOnDiskIO returns true if the dataset is persisted
// to disk in distributed mode.
func (d *Dataset) GetIsOnDiskIO() bool { _ = "STUB: not implemented"; return false }
