package Data

import "testing"

func contains(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

func containsAll(a []string, s ...string) bool {
	for _, v := range s {
		if !contains(a, v) {
			return false
		}
	}
	return true
}

func TestNewAtelier(t *testing.T) {
	a := NewAtelier()
	if a == nil {
		t.Error("Can't create NewAtelier")
	}
}

func getAtelier() *AtelierData {
	a := NewAtelier()
	i := Item{"Test",
		[]string{"test1", "test2"},
		[]string{"(type1)", "(type2)"}}
	a.AddItem(i)
	return a
}

func TestAddItem(t *testing.T) {
	a := NewAtelier()
	i := Item{"Test", []string{"test1", "test2"}, []string{"type1", "type2"}}
	a.AddItem(i)
	t.Logf("TestAddItem: %v", a.Nodes())
	if !containsAll(a.Nodes(), "Test",
		"test1", "test2", "type1", "type2") {
		t.Error("AddItem: something is missing")
	}
}

func TestAddNode(t *testing.T) {
	a := getAtelier()
	a.addNode("newTest")
	if !a.IsItem("newTest") {
		t.Error("AddNode not working")
	}
}

func TestAddEdge(t *testing.T) {
	a := getAtelier()
	a.addEdge("test1", "test2")
	if !a.graph["test1"].contains("test2") {
		t.Error("addEdge not working")
	}
}

func TestIsItem(t *testing.T) {
	a := getAtelier()
	if !a.IsItem("Test") {
		t.Error("IsItem: not found")
	}
}

func TestDijkstra(t *testing.T) {
	a := NewAtelier()
	items := []Item{
		Item{"one",
			[]string{"five", "two"},
			[]string{"five", "two"},
		},
		Item{"five",
			[]string{"four", "two", "one"},
			[]string{"four", "two", "one"},
		},
		Item{"two",
			[]string{"three", "five", "one"},
			[]string{"three", "five", "one"},
		},
		Item{"three",
			[]string{"four", "two"},
			[]string{"four", "two"},
		},
		Item{"four",
			[]string{"six", "three", "five"},
			[]string{"six", "three", "five"},
		},
		Item{"six",
			[]string{"four"},
			[]string{"four"},
		},
	}
	for _, v := range items {
		a.AddItem(v)
	}
	dist, prev := a.dijkstra("one")
	t.Logf("dist: %v", dist)
	t.Logf("prev: %v", prev)
	if !(dist["two"] == 1 && dist["four"] == 2 && dist["three"] == 2 &&
		dist["one"] == 0 && dist["six"] == 3 && dist["five"] == 1) {
		t.Error("bla")
	}

}
