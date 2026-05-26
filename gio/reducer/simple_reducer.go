package reducer

import (
	"github.com/chrislusf/gleam/gio"
)

var (
	SumInt64   = gio.RegisterReducer(sumInt64)
	SumFloat64 = gio.RegisterReducer(sumFloat64)
)

func sumInt64(x, y interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func sumFloat64(x, y interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }
