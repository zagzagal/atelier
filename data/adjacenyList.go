package Data

import ()

type adjList map[string]bool

func (a *adjList) add(s string) {
	(*a)[s] = true
}

func (a *adjList) remove(s string) {
	delete(*a, s)
}

func (a adjList) contains(s string) bool {
	_, ok := a[s]
	return ok
}

func (a adjList) edges() []string {
	s := make([]string, len(a))
	i := 0
	for k, _ := range a {
		s[i] = k
		i++
	}
	return s
}

func (a adjList) String() string {
	s := "[ "
	i := 0
	for k, _ := range a {
		s += k
		if i < len(a)-1 {
			s += " -> "
		} else {
			s += " ]"
		}
		i++
	}
	return s
}

type Edge struct {
	head string
	tail string
}

func (e Edge) String() string {
	return "\"" + e.head + "\" -> \"" + e.tail + "\""
}

type AdjacencyList map[string]*adjList

// Contains returns true if an item exists
func (a AdjacencyList) Contains(s string) bool {
	_, ok := a[s]
	return ok
}

func (a *AdjacencyList) AddEdge(s, d string) {
	a.AddNode(s)
	a.AddNode(d)
	(*a)[s].add(d)
}

func (a *AdjacencyList) RemoveEdge(s, d string) {
	if a.Contains(s) {
		(*a)[s].remove(d)
	}
}

func (a *AdjacencyList) AddNode(s string) {
	if !a.Contains(s) {
		(*a)[s] = &adjList{}
	}
}

func (a *AdjacencyList) RemoveNode(s string) {
	delete(*a, s)
	for _, v := range *a {
		v.remove(s)
	}
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

func (a AdjacencyList) Nodes() []string {
	slice := make([]string, len(a))
	i := 0
	for k, _ := range a {
		slice[i] = k
		i++
	}
	return slice
}
