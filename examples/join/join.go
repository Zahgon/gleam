package main

import (
	"flag"

	"github.com/chrislusf/gleam/gio"
)

var (
	isDistributed = flag.Bool("distributed", false, "run in distributed mode")
)

func main() {

	gio.Init()

	join1()

	hashjoin()

}

func join1() { _ = "STUB: not implemented"; return }

func hashjoin() { _ = "STUB: not implemented"; return }
