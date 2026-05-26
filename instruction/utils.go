package instruction

import (
	"io"

	"github.com/chrislusf/gleam/pb"
	"github.com/chrislusf/gleam/util"
)

func toInts(indexes []int32) []int { _ = "STUB: not implemented"; return nil }

func toOrderBys(orderBys []*pb.OrderBy) (ret []OrderBy) { _ = "STUB: not implemented"; return nil }

// create a channel to aggregate values of the same key
// automatically close original sorted channel
func newChannelOfValuesWithSameKey(name string, sortedChan io.Reader, indexes []int) chan util.Row {
	_ = "STUB: not implemented"
	return nil
}

// fmt.Printf("%s join read len=%d, row: %s\n", name, len(row), row[0])

func max(x, y int64) int64 { _ = "STUB: not implemented"; return 0 }
