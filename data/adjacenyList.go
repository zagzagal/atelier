package Data

import ()

type Edge struct {
	head string
	tail string
}

func (e Edge) String() string {
	return "\"" + e.head + "\" -> \"" + e.tail + "\""
}

type AdjacencyList map[string]subList

func newAdjList() *AdjacencyList {
	var a AdjacencyList
	a = make(AdjacencyList)
	return &a
}

// Contains returns true if an item exists
func (a AdjacencyList) Contains(s string) bool {
	for k, _ := range a {
		if k == s {
			return true
		}
	}
	return false
}

func (a *AdjacencyList) AddEdge(s, d string) {
	if !a.Contains(s) {
		a.AddNode(s)
	}
	if !a.Contains(d) {
		a.AddNode(d)
	}
	(*a)[s] = (*a)[s].add(d)
}

func (a *AdjacencyList) RemoveEdge(s, d string) {
	if a.Contains(s) {
		(*a)[s] = (*a)[s].remove(d)
	}
}

func (a *AdjacencyList) AddNode(s string) {
	if !a.Contains(s) {
		(*a)[s] = newSubList()
	}
}

func (a *AdjacencyList) RemoveNode(s string) {
	for _, v := range *a {
		if v.contains(s) {
			v.remove(s)
		}
	}
	delete(*a, s)
}

func (a AdjacencyList) getEdges(s string) []Edge {
	if a.Contains(s) {
		tail := a[s].edges()
		edges := make([]Edge, len(tail))
		for k, v := range tail {
			edges[k] = Edge{s, v}
		}
		return edges
	}
	return []Edge{}
}

func (a AdjacencyList) Edges() []Edge {
	ans := make([]Edge, len(a))
	for k, _ := range a {
		ans = append(ans, a.getEdges(k)...)
	}
	return ans
}

func (a AdjacencyList) String() string {
	var s string
	s += "[ "
	for k, v := range a {
		s += "{ " + k + " : " + v.String() + " }"
	}
	s += " ] "
	return s
}

func (a AdjacencyList) Nodes() []string {
	slice := make([]string, len(a))
	i := 0
	for k, _ := range a {
		slice[i] = k
		i++
	}
	return slice
}
