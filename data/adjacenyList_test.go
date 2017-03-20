package Data

import "testing"

func TestContains(t *testing.T) {
	a := adjList{"This", "is", "A", "Test"}
	t.Logf("%v", a)
	if !a.contains("This") {
		t.Error("TestContains: contains failed")
	}
}

func TestAdd(t *testing.T) {
	a := adjList{"this", "is", "A", "Test"}
	t.Logf("TestAdd: Before: %v", a)
	a.add("boo")
	t.Logf("TestAdd: After: %v", a)
	if !a.contains("boo") {
		t.Error("TestAdd: Add Failed")
	}
}

func TestALContains(t *testing.T) {
	a := make(AdjacencyList)
	a.AddEdge("this", "that")
	if !a.Contains("this") {
		t.Error("TestALContains: Contains Failed")
	}
}

func TestGetEdges(t *testing.T) {
	a := make(AdjacencyList)
	a.AddEdge("this", "that")
	a.AddEdge("this", "also")
	e := a.GetEdges("this")
	t.Logf("TestGetEdges: %v", e)
	if e[0] != "that" || e[1] != "also" {
		t.Error("TestGetEdges: edges didn't match up")
	}
}

func TestNodes(t *testing.T) {
	a := make(AdjacencyList)
	a.AddEdge("this", "that")
	a.AddEdge("also", "this")
	e := a.Nodes()
	t.Logf("TestNodes: %v", e)
	if len(e) != 3 {
		t.Error("TestNodes: nodecount doesn't match")
	}

	f := func(a []string, s string) bool {
		for _, v := range a {
			if v == s {
				return false
			}
		}
		t.Logf("TestNodes: %v not found", s)
		return true
	}

	if f(e, "also") && f(e, "this") && f(e, "that") {
		t.Error("TestNodes: nodes didn't match up")
	}
}
