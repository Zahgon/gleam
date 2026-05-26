package flow

import (
	"github.com/chrislusf/gleam/gio"
)

// Mapper runs the mapper registered to the mapperId.
// This is used to execute pure Go code.
func (d *Dataset) Map(name string, mapperId gio.MapperId) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func add1ShardTo1Step(d *Dataset) (ret *Dataset, step *Step) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Select selects multiple fields into the next dataset. The index starts from 1.
// The first one is the key
func (d *Dataset) Select(name string, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

// Select selects multiple fields into the next dataset. The index starts from 1.
func (d *Dataset) SelectKV(name string, keys, values *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

// LocalLimit take the local first n rows and skip all other rows.
func (d *Dataset) LocalLimit(name string, n int, offset int) *Dataset {
	_ = "STUB: not implemented"
	return nil
}
