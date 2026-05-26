package flow

type DasetsetHint func(d *Dataset)

// Hint adds options for previous dataset.
func (d *Dataset) Hint(options ...DasetsetHint) *Dataset { _ = "STUB: not implemented"; return nil }

// TotalSize hints the total size in MB for all the partitions.
// This is usually used when sorting is needed.
func TotalSize(n int64) DasetsetHint { _ = "STUB: not implemented"; return *new(DasetsetHint) }

// PartitionSize hints the partition size in MB.
// This is usually used when sorting is needed.
func PartitionSize(n int64) DasetsetHint { _ = "STUB: not implemented"; return *new(DasetsetHint) }

// OnDisk ensure the intermediate dataset are persisted to disk.
// This allows executors to run not in parallel if executors are limited.
func (d *Dataset) OnDisk(fn func(*Dataset) *Dataset) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

/*

// Datacenter hints the previous dataset output location
func Datacenter(dc string) DasetsetOption {
	return func(c *DasetsetConfig) {
		c.Datacenter = dc
	}
}

// Rack hints the previous dataset output location
func Rack(rack string) DasetsetOption {
	return func(c *DasetsetConfig) {
		c.Rack = rack
	}
}

*/
