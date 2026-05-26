// Package market is a market system to match tasks with resources.
//
// Continuous Double Auction Protocol
// (CDA) is to allocate the best possible resource to an arriving
// task and to prioritise tasks according to their price
// bid. When a Task Query object arrives at the market the
// protocol searches all available resource offers and returns
// the first occurrence of the ’best’ match, i.e. the
// cheapest or the fastest resource which satisfies the task’s
// constraints. Whenever a resource becomes available and
// there are several tasks waiting, the one with the highest
// price bid is processed first.

// this implmentation only support one supplier

package market

import (
	//"fmt"
	"sync"
)

type Object interface{}

type Requirement interface{}

type Demand struct {
	Requirement Requirement
	Bid         float64
	ReturnChan  chan Supply
}

type Supply struct {
	Object Object
}

type Market struct {
	Demands    []Demand
	Supplies   []Supply
	Lock       sync.Mutex
	ScoreFn    func(Requirement, float64, Object) float64
	FetchFn    func([]Demand)
	hasDemands *sync.Cond
}

func NewMarket() *Market { _ = "STUB: not implemented"; return nil }

func (m *Market) SetScoreFunction(scorer func(Requirement, float64, Object) float64) *Market {
	_ = "STUB: not implemented"
	return nil
}

func (m *Market) SetFetchFunction(fn func([]Demand)) *Market { _ = "STUB: not implemented"; return nil }

// retChan should be a buffered channel
func (m *Market) AddDemand(r Requirement, bid float64, retChan chan Supply) {
	_ = "STUB: not implemented"
	return
}

func (m *Market) FetcherLoop() {
	_ = "STUB: not implemented"

	// println("FetcherLoop Lock:", len(m.Demands))
	return
}

// println("FetcherLoop wait:", len(m.Demands))

// println("FetcherLoop UnLock:", len(m.Demands))

// println("fetching current demands:", len(m.Demands))

// println("fetching finished demands:", len(m.Demands))

func (m *Market) ReturnSupply(s Supply) { _ = "STUB: not implemented"; return }

func (m *Market) AddSupply(supply Supply) { _ = "STUB: not implemented"; return }

func (m *Market) pickBestSupplyFor(r Requirement) (ret Supply, matched bool) {
	_ = "STUB: not implemented"
	return *new(Supply), false
}

func (m *Market) pickBestDemandFor(supply Supply) (ret Demand, matched bool) {
	_ = "STUB: not implemented"
	return *new(Demand), false
}

// fmt.Printf("matched demand: %+v\n", ret)
