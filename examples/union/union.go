package main

import (
	"fmt"

	"github.com/chrislusf/gleam/gio"
)

func runEnableParallel() { _ = "STUB: not implemented"; return }

// if isParallel=true
// datasetOut reading from datasetA/datasetB/datasetC competition
// so shards in datasetOut get rows NOT in the order of datasetA/datasetB/datasetC

func runDisableParallel() { _ = "STUB: not implemented"; return }

// if isParallel=false
// datasetOut reading from datasetA/datasetB/datasetC in order,
//     that means it must ALL data in datasetA are readed, then datasetB can be reading
// so shards in datasetOut get rows in the order of datasetA/datasetB/datasetC

func main() {
	gio.Init()

	fmt.Println("================================================================")
	fmt.Println("==== Union isParallel=true ====")
	runEnableParallel()
	fmt.Println("================================================================")
	fmt.Println("==== Union isParallel=false ====")
	runDisableParallel()
	fmt.Println("================================================================")

}
