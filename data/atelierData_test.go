package Data

import "testing"

func contains(t *testing.T, a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	t.Logf("Contains: %v not in %v", s, a)
	return false
}

func containsAll(t *testing.T, a []string, s ...string) bool {
	for _, v := range s {
		if !contains(t, a, v) {
			return false
		}
	}
	return true
}

func checkAdd(t *testing.T, a AtelierData, i Item) {
	adj, ok := a.graph[i.Name]
	t.Logf("CheckAdd: [ Item: %v ] [ Graph: %v ]", i, adj)
	if !ok {
		t.Errorf("Item %v is not in graph", i.Name)
	}
	for k, _ := range adj {
		if !contains(t, i.Types, k) {
			t.Errorf("Item %v, type %v is not in graph [%v]",
				i.Name, k, adj)
		}
	}
	for _, v := range i.Ingredients {
		if !contains(t, a.graph[v].edges(), i.Name) {
			t.Errorf("Itme %v, type %v ins not in graph [%v]",
				i.Name, i.Name, v)
		}
	}
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
	a.addEdge("test1", "test3")
	a.addEdge("test1", "test4")
	b := a.graph["test1"]
	n, ok := a.graph["test2"]
	if !ok {
		t.Error("addEdge: not adding nodes right")
	}
	if contains(t, n.edges(), "test1") {
		t.Error("addEdge: reversing edges")
	}
	if !containsAll(t, b.edges(), "test2", "test3", "test4") {
		t.Error("addEdge not working")
	}
}

func TestIsItem(t *testing.T) {
	a := getAtelier()
	if !a.IsItem("Test") {
		t.Error("IsItem: not found")
	}
}

func TestAddItem(t *testing.T) {
	a := NewAtelier()
	i := Item{"Test", []string{"test1", "test2"}, []string{"type1", "type2"}}
	a.AddItem(i)
	t.Logf("TestAddItem: %v", a.Nodes())
	if !containsAll(t, a.Nodes(), "Test",
		"test1", "test2", "type1", "type2") {
		t.Error("AddItem: something is missing")
	}
	checkAdd(t, *a, i)
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
