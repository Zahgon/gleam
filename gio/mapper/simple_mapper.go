package mapper

import (
	"github.com/chrislusf/gleam/gio"
)

var (
	Tokenize  = gio.RegisterMapper(tokenize)
	AppendOne = gio.RegisterMapper(addOne)
)

func tokenize(row []interface{}) error { _ = "STUB: not implemented"; return nil }

func addOne(row []interface{}) error { _ = "STUB: not implemented"; return nil }
