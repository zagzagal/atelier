package Data

import (
	"container/heap"
)

type QItem struct {
	value    string
	priority int
	index    int
}

type MinPriorityQueue []*QItem

func (pq MinPriorityQueue) Len() int { return len(pq) }

func (pq MinPriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq MinPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MinPriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*QItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *MinPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq MinPriorityQueue) get(s string) *QItem {
	for _, v := range pq {
		if v.value == s {
			return v
		}
	}
	return nil
}

func (pq *MinPriorityQueue) update(item *QItem, v string, p int) {
	item.value = v
	item.priority = p
	heap.Fix(pq, item.index)
}
