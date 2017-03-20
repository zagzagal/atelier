package Data

import "testing"

func TestContains(t *testing.T) {
	a := adjList{"This": true, "is": true, "A": true, "Test": true}
	t.Logf("%v", a)
	if !a.contains("This") {
		t.Error("TestContains: contains failed")
	}
}

func TestAdd(t *testing.T) {
	a := adjList{"this": true, "is": true, "A": true, "Test": true}
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

func TestEdges(t *testing.T) {
	a := adjList{"this": true, "is": true, "a": true, "test": true}
	t.Logf("TestEdge: %v", a)
	e := a.edges()
	t.Logf("TestEdge: %v", e)
	f := func(s string) bool {
		for _, v := range e {
			if v == s {
				return false
			}
		}
		t.Logf("TestNodes: %v not found", s)
		return true
	}

	if f("this") && f("is") && f("a") && f("test") {
		t.Error("TestEdges: edges didn't match up")
	}
}

func TestALEdges(t *testing.T) {
	a := make(AdjacencyList)
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
	a := make(AdjacencyList)
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
