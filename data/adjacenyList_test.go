package Data

import "testing"

func TestConstains(t *testing.T) {
	var a AdjacencyList
	a = AdjacencyList{
		"Test":  subList{"this": true, "is": true, "A": true, "Test": true},
		"Test2": subList{"this": true, "is": true, "A": true, "Test": true},
		"Test3": subList{"this": true, "is": true, "A": true, "Test": true}}
	if !a.Contains("Test") || !a.Contains("Test2") || !a.Contains("Test3") {
		t.Error("TestContains: contains failed")
	}
}

func TestALEdges(t *testing.T) {
	a := newAdjList()
	if a == nil {
		t.Error("NewAdjList failed")
	}
	a.AddEdge("this", "that")
	a.AddEdge("this", "also")
	e := a.getEdges("this")
	t.Logf("TestGetEdges: %v", e)
	f := func(s string) bool {
		for _, v := range e {
			if v.tail == s {
				return false
			}
		}
		t.Logf("TestNodes: %v not found", s)
		return true
	}

	if f("that") || f("also") {
		t.Error("TestGetEdges: edges didn't match up")
	}
}

func TestNodes(t *testing.T) {
	a := newAdjList()
	a.AddEdge("this", "that")
	a.AddEdge("also", "this")
	e := a.Nodes()
	t.Logf("TestNodes: %v", e)
	if len(e) != 3 {
		t.Error("TestNodes: nodecount doesn't match")
	}

	f := func(s string) bool {
		for _, v := range e {
			if v == s {
				return false
			}
		}
		t.Logf("TestNodes: %v not found", s)
		return true
	}

	if f("also") && f("this") && f("that") {
		t.Error("TestNodes: nodes didn't match up")
	}
}
