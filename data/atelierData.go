// Package alows for the data model for the atelier series
// contains an item library and pathing functions
package Data

import (
	//"errors"
	//"github.com/zagzagal/queue"
	"container/heap"
	//"log"
)

const infinite = 99999
const null = -99999
const theta = 0

// The main datasctruct for the library
type AtelierData struct {
	items map[string]Item
	graph AdjacencyList
}

// Returns a new instatiated instace of the dataset
func NewAtelier() *AtelierData {
	a := new(AtelierData)
	a.items = make(map[string]Item)
	a.graph = make(AdjacencyList)
	return a
}

// Adds an item to the dataset
func (a *AtelierData) AddItem(i Item) {
	//log.Printf("%v", i)
	if i.Name != "" {
		a.items[i.Name] = i
		//log.Printf("%v added: %v [%v]", i.Name, a.items[i.Name], i.equals(a.items[i.Name]))
		a.addNode(i.Name)
		for _, v := range i.Ingredients {
			a.addEdge(v, i.Name)
		}
		for _, v := range i.Types {
			a.addEdge(i.Name, v)
		}
	}
}

func (a AtelierData) dijkstra(s string) (map[string]int, map[string]string) {
	dist := make(map[string]int)
	prev := make(map[string]string)

	nodes := a.Nodes()
	q := make(MinPriorityQueue, len(nodes))
	dist[s] = 0

	for k, v := range nodes {
		if v != s {
			dist[v] = infinite
			prev[v] = ""
		}
		q[k] = &QItem{
			value:    v,
			priority: dist[v],
			index:    k,
		}
	}
	heap.Init(&q)

	for q.Len() > 0 {
		u := heap.Pop(&q).(*QItem).value
		for _, e := range a.graph.getEdges(u) {
			v := e.tail
			alt := dist[u] + 1
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				it := q.get(v)
				q.update(it, v, alt)
			}
		}
	}
	return dist, prev
}

func (a AtelierData) getShortestPath(s, d string) []string {
	_, prev := a.dijkstra(s)

	ans := []string{}
	m, ok := prev[d]
	done := false
	for ok && !done {
		ans = append([]string{d}, ans...)
		if m != s {
			d = m
			m, ok = prev[d]
			if !ok {
				return []string{}
			}
		} else {
			done = true
		}
	}
	return append([]string{m}, ans...)
}

func (a *AtelierData) addNode(s string) {
	//log.Println("Adding Node ", s)
	a.graph.AddNode(s)
}

func (a *AtelierData) addEdge(start, dest string) error {
	a.graph.AddEdge(start, dest)
	return nil
}

func (a *AtelierData) AddPath(start, dest string) error {
	if !a.IsItem(start) {
		a.addNode(start)
	}
	if !a.IsItem(dest) {
		a.addNode(dest)
	}
	return a.addEdge(start, dest)
}

// returns the dot file representation of the dataset
func (a *AtelierData) PrintDot() (s string) {
	s = "digraph test {\n"
	for _, v := range a.graph.Edges() {
		s += "  " + v.String() + ";\n"
	}
	return s + "}"
}

// returns the path from start to dest as a string
func (a *AtelierData) GetPath(start, dest string) ItemPath {
	x := a.getShortestPath(start, dest)
	var ans ItemPath
	ans.Item = x
	return ans
}

func (a *AtelierData) printNodeList() (s string) {
	nodes := a.graph.Nodes()
	for _, v := range nodes {
		s += v + "\n"
	}
	return
}

// Returns a list of all the items (not types) in the dataset
func (a *AtelierData) Items() (s []string) {
	for _, v := range a.items {
		s = append(s, v.Name)
	}
	return
}

// Returns a list of all the items and item types in the dataset
func (a *AtelierData) Nodes() (s []string) {
	for _, v := range a.graph.Nodes() {
		s = append(s, v)
	}
	return
}

// checks to see if the item is in the datase
func (a *AtelierData) IsItem(s string) bool {
	return a.graph.Contains(s)
}

/*func (a *AtelierData) usedIn(s string) string {
	if !a.IsItem(s) {
		return ""
	}
	n, _ := a.findNode(s)
	return a.mapPre(n, func(x graph.Node) string {
		v, ok := a.id2Name[x.ID()]
		if ok {
			if !strings.Contains(v, "(") {
				return v + "\n"
			}
		}
		return ""
	}, 2)
}*/

func (a AtelierData) GetItemData(s string) (i Item) {
	n, ok := a.items[s]
	//log.Printf("[%v] %v", ok, n)
	//log.Printf("Item List: %v", a.Nodes())
	if ok {
		return n
	}
	return Item{}
}

func (a AtelierData) GetRawItemData(s string) string {
	n, _ := a.graph[s]
	//log.Printf("[%v] %v", ok, n)
	return s + " " + n.String()
}
