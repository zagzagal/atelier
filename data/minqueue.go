package Data

import (
	"container/heap"
)

//QItem is a Struct defining the objects in the min queue
type QItem struct {
	value    string
	priority int
	index    int
}

//MinPriorityQueue is a Struct defing the min queue
type MinPriorityQueue []*QItem

//Init is a Heap.Init helper function
func (pq *MinPriorityQueue) Init() {
	heap.Init(pq)
}

//HPop is a heap.pop helper function
func (pq *MinPriorityQueue) HPop() *QItem {
	u := heap.Pop(pq).(*QItem)
	return u
}

func (pq MinPriorityQueue) Len() int { return len(pq) }

func (pq MinPriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq MinPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

//Push for heap interface
func (pq *MinPriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*QItem)
	item.index = n
	*pq = append(*pq, item)
}

//Pop for heap interface
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
