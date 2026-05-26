package main

import (
	"flag"

	"github.com/chrislusf/gleam/gio"
)

var (
	size          = flag.Int("size", 0, "0 for small, 1 for 1GB, 2 for 10GB")
	isDistributed = flag.Bool("distributed", false, "distributed mode or not")
	isInMemory    = flag.Bool("inMemory", true, "distributed mode but only through memory")
	isProfiling   = flag.Bool("isProfiling", false, "profiling the flow")

	splitter = gio.RegisterMapper(splitLine)
)

func main() {
	flag.Parse()
	gio.Init()

	bigFile := *size

	fileName := "/Users/chris/Desktop/record_10K_input.txt"
	partition := 2
	size := int64(10)
	if bigFile == 1 {
		fileName = "/Users/chris/Desktop/record_1GB_input.txt"
		partition = 4
		size = 1024
	}
	if bigFile == 2 {
		fileName = "/Users/chris/Desktop/record_10GB_input.txt"
		partition = 40
		size = 10240
	}

	gleamSortDistributed(fileName, size, partition, *isDistributed, *isInMemory)

}

func linuxSortDistributed(fileName string, partition int) { _ = "STUB: not implemented"; return }

func linuxSortStandalone(fileName string, partition int) { _ = "STUB: not implemented"; return }

func gleamSortDistributed(fileName string, size int64, partition int, isDistributed, isInMemory bool) {
	_ = "STUB: not implemented"
	return
}

func splitLine(row []interface{}) error { _ = "STUB: not implemented"; return nil }
