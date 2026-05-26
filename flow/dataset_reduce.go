package flow

import (
	"github.com/chrislusf/gleam/gio"
)

// ReduceByKey runs the reducer registered to the reducerId,
// combining rows with the same key fields into one row
func (d *Dataset) ReduceByKey(name string, reducerId gio.ReducerId) (ret *Dataset) {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) ReduceBy(name string, reducerId gio.ReducerId, keyFields *SortOption) (ret *Dataset) {
	_ = "STUB: not implemented"
	return nil
}

// Reduce runs the reducer registered to the reducerId,
// combining all rows into one row
func (d *Dataset) Reduce(name string, reducerId gio.ReducerId) (ret *Dataset) {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) LocalReduceBy(name string, reducerId gio.ReducerId, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

// add key indexes for reducer command line option

// combine all rows directly
