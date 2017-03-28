package Data

import ()

// Edge defines an edge struct
type Edge struct {
	head string
	tail string
}

func (e Edge) String() string {
	return "\"" + e.head + "\" -> \"" + e.tail + "\""
}

//AdjacencyList defines a simple graph stucture
type AdjacencyList map[string]subList

func newAdjList() *AdjacencyList {
	var a AdjacencyList
	a = make(AdjacencyList)
	return &a
}

// Contains returns true if an item exists
func (a AdjacencyList) Contains(s string) bool {
	for k := range a {
		if k == s {
			return true
		}
	}
	return false
}

//AddEdge adds a relationship between start and destination nodes
func (a *AdjacencyList) AddEdge(s, d string) {
	if !a.Contains(s) {
		a.AddNode(s)
	}
	if !a.Contains(d) {
		a.AddNode(d)
	}
	(*a)[s] = (*a)[s].add(d)
}

//RemoveEdge removes a relationship between start and destination nodes
func (a *AdjacencyList) RemoveEdge(s, d string) {
	if a.Contains(s) {
		(*a)[s] = (*a)[s].remove(d)
	}
}

//AddNode adds a node s to the graph
func (a *AdjacencyList) AddNode(s string) {
	if !a.Contains(s) {
		(*a)[s] = newSubList()
	}
}

//RemoveNode removes a node s from the graph and all of its relationships
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

//Edges returns an array of all of Edges
func (a AdjacencyList) Edges() []Edge {
	ans := make([]Edge, len(a))
	for k := range a {
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

//Nodes returns an array of all of the nodes
func (a AdjacencyList) Nodes() []string {
	slice := make([]string, len(a))
	i := 0
	for k := range a {
		slice[i] = k
		i++
	}
	return slice
}
