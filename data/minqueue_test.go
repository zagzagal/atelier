package Data

import (
	"container/heap"
	"testing"
)

func items() map[string]int {
	return map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
}

func getTQueue() MinPriorityQueue {
	it := items()
	i := 0
	pq := make(MinPriorityQueue, len(it))

	for k, v := range it {
		pq[i] = &QItem{
			value:    k,
			priority: v,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	return pq
}

func TestingLen(t *testing.T) {
	pq := getTQueue()
	if pq.Len() != 3 {
		t.Error("Len: something is wrong")
	}
}

func TestningLess(t *testing.T) {
	pq := getTQueue()
	if pq.Less(2, 1) {
		t.Error("Less: something is wrong")
	}
}
