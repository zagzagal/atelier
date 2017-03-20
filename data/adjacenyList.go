package AtelierData

import ()

type adjList []string

func (a *adjList) add(s string) {
	*a = append(*a, s)
}

func (a adjList) contains(s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

func (a adjList) String() string {
	s := "[ "
	for k, v := range a {
		s += v
		if k < len(a)-1 {
			s += " -> "
		} else {
			s += " ]"
		}
	}
	return s
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

func (a *AdjacencyList) AddNode(s string) {
	if !a.Contains(s) {
		(*a)[s] = &adjList{}
	}
}

func (a AdjacencyList) GetEdges(s string) []string {
	if a.Contains(s) {
		slice := make([]string, len(*a[s]))
		for k, v := range *a[s] {
			slice[k] = v
		}
		return slice
	}
	return []string{}
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
