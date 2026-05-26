package flow

import (
	"github.com/chrislusf/gleam/instruction"
)

type SortOption struct {
	orderByList []instruction.OrderBy
}

func Field(indexes ...int) *SortOption { _ = "STUB: not implemented"; return nil }

func OrderBy(index int, ascending bool) *SortOption { _ = "STUB: not implemented"; return nil }

// OrderBy chains a list of sorting order by
func (o *SortOption) By(index int, ascending bool) *SortOption {
	_ = "STUB: not implemented"
	return nil
}

// return a list of indexes
func (o *SortOption) Indexes() []int { _ = "STUB: not implemented"; return nil }

func (o *SortOption) String() string { _ = "STUB: not implemented"; return "" }
