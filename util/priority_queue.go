package util

import (
	"sync"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value interface{} // The value of the item; arbitrary.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index    int // The index of the item in the heap.
	sourceId int // payload containing the value's source id
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
	lessFunc func(a, b interface{}) bool
	lock     sync.RWMutex
	items    []*Item
}

func NewPriorityQueue(lessFunc func(a, b interface{}) bool) *PriorityQueue {
	_ = "STUB: not implemented"
	return nil
}

func (pq *PriorityQueue) Enqueue(x interface{}, sourceId int) { _ = "STUB: not implemented"; return }

func (pq *PriorityQueue) Dequeue() (interface{}, int) { _ = "STUB: not implemented"; return nil, 0 }

func (pq *PriorityQueue) Top() interface{} { _ = "STUB: not implemented"; return nil }

func (pq *PriorityQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (pq *PriorityQueue) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (pq *PriorityQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (pq *PriorityQueue) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (pq *PriorityQueue) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// for safety
