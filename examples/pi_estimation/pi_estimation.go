package main

import (
	"flag"
	"os"
	"runtime/pprof"

	"github.com/chrislusf/gleam/gio"
)

var (
	master = flag.String("master", "localhost:45326", "master server location")

	monteCarloMapperId = gio.RegisterMapper(monteCarloMapper)
	sumReducerId       = gio.RegisterReducer(sumReducer)

	times  = 1024 * 1024 * 2560
	factor = 1024 * 1024
)

func main() {

	gio.Init()
	flag.Parse()

	f, _ := os.Create("p.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// uncomment this line if you setup the gleam master and agents
	testPureGoGleam("distributed parallel 7", false)
	testPureGoGleam("local mode parallel 7", true)

	testDirectGo()

	testDirectGoConcurrent()

	// this is not fair since many optimization is not applied
	testLocalFlow()
}

func testPureGoGleam(name string, isLocal bool) { _ = "STUB: not implemented"; return }

func testDirectGo() { _ = "STUB: not implemented"; return }

func testDirectGoConcurrent() { _ = "STUB: not implemented"; return }

func testLocalFlow() { _ = "STUB: not implemented"; return }

func monteCarloMapper(row []interface{}) error { _ = "STUB: not implemented"; return nil }

func sumReducer(x, y interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }
