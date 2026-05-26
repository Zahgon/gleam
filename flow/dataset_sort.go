package flow

import (
	"github.com/chrislusf/gleam/instruction"
)

type pair struct {
	keys []interface{}
	data []byte
}

// Distinct sort on specific fields and pick the unique ones.
// Required Memory: about same size as each partition.
// example usage: Distinct(Field(1,2)) means
// distinct on field 1 and 2.
// TODO: optimize for low cardinality case.
func (d *Dataset) Distinct(name string, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

// Sort sort on specific fields, default to the first field.
// Required Memory: about same size as each partition.
// example usage: Sort(Field(1,2)) means
// sorting on field 1 and 2.
func (d *Dataset) Sort(name string, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) SortByKey(name string) *Dataset { _ = "STUB: not implemented"; return nil }

// Top streams through total n items, picking reverse ordered k items with O(n*log(k)) complexity.
// Required Memory: about same size as n items in memory
func (d *Dataset) Top(name string, k int, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) LocalDistinct(name string, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) LocalSort(name string, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dataset) LocalTop(name string, n int, sortOption *SortOption) *Dataset {
	_ = "STUB: not implemented"
	return nil
}

func isOrderByEquals(a []instruction.OrderBy, b []instruction.OrderBy) bool {
	_ = "STUB: not implemented"
	return false
}

func getReverseOrderBy(a []instruction.OrderBy) (reversed []instruction.OrderBy) {
	_ = "STUB: not implemented"
	return nil
}
