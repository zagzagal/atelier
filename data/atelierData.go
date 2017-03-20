// Package alows for the data model for the atelier series
// contains an item library and pathing functions
package AtelierData

/*import (
	"errors"
	"github.com/zagzagal/queue"
	//"fmt"
	"strings"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
)

const infinite = 99999
const null = -99999
const theta = 0

// The main datasctruct for the library
type AtelierData struct {
	id2Name map[int]string
	name2Id map[string]int

	graph *cayley.Handle
}

// Returns a new instatiated instace of the dataset
func NewAtelier() *AtelierData {
	a := new(AtelierData)
	a.id2Name = map[int]string{}
	a.name2Id = map[string]int{}
	a.graph, _ = cayley.NewMemoryGraph()
	return a
}

// Adds an item to the dataset
func (a *AtelierData) AddItem(i Item) {
	if i.Name != "" {
		a.addNode(i.Name)
		for _, v := range i.Ingredients {
			if !a.IsItem(v) {
				a.addNode(v)
			}
			a.addEdge(i.Name, v)
		}
		for _, v := range i.Types {
			if !a.IsItem(v) {
				a.addNode(v)
			}
			a.addEdge(v, i.Name)
		}
	}
}

func (a *AtelierData) floydWarshallAlg() {
	n := len(a.id2Name)
	dist := make([][]int, n)
	next := make([][]int, n)
	for k := 0; k < n; k++ {
		dist[k] = make([]int, n)
		next[k] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[k][j] = infinite
			next[k][j] = null
		}
}
	for _, v := range a.graph.Nodes() {
		dist[v.ID()-1][v.ID()-1] = theta
	}
	for _, v := range a.graph.Edges() {
		dist[v.From().ID()-1][v.To().ID()-1] = 1
		next[v.From().ID()-1][v.To().ID()-1] = v.To().ID() - 1
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][j] > (dist[i][k] + dist[k][j]) {
					dist[i][j] = dist[i][k] + dist[k][j]
					next[i][j] = next[i][k]
				}
			}
		}
	}
	a.floydDist = dist
	a.floydNext = next
}

func (a *AtelierData) addNode(s string) {
	n := a.graph.NewNode()
	a.graph.AddNode(n)
	a.id2Name[n.ID()] = s
	a.name2Id[s] = n.ID()
}

func (a *AtelierData) addEdge(start, dest string) error {
	s, err := a.getNode(start)
	if err != nil {
		return errors.New("node " + start + " does not exist")
	}
	d, err := a.getNode(dest)
	if err != nil {
		return errors.New("node " + dest + " does not exist")
	}
	a.graph.AddDirectedEdge(&concrete.Edge{s, d}, 2)
	a.floydWarshallAlg()
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

func (a *AtelierData) getNode(s string) (graph.Node, error) {
	id, err := a.name2Id[s]
	if err != true {
		return nil, errors.New("node does not exist")
	}
	nodes := a.graph.NodeList()
	for _, v := range nodes {
		if v.ID() == id {
			return v, nil
		}
	}
	return nil, nil
}

// returns the dot file representation of the dataset
func (a *AtelierData) PrintDot() (s string) {
	s = "digraph test {\n"
	for _, v := range a.graph.EdgeList() {
		s += "  " + a.printEdgeForDot(v) + ";\n"
	}
	return s + "}"
}

func (a *AtelierData) printEdgeForDot(e graph.Edge) string {
	return "\"" + a.id2Name[e.Head().ID()] + "\"" + " -> " + "\"" +
		a.id2Name[e.Tail().ID()] + "\""
}

// returns the path from start to dest as a string
func (a *AtelierData) GetPath(start, dest string) (ItemPath, error) {
	x, err := a.findPath(start, dest)
	if err == nil {
		var ans ItemPath
		ans.Item = x
		return ans, err
	}
	return ItemPath{Item: nil}, err
}

func (a *AtelierData) findFloydPath(start, dest string) []string {
	u, _ := a.getNode(start)
	v, _ := a.getNode(dest)
	path := queue.NewQueue()
	var ans []string

	if a.floydNext[u.ID()-1][v.ID()-1] == null {
		return nil
	}
	path.Push(a.id2Name[u.ID()])
	for u.ID() != v.ID() {
		u, _ = a.findNodeID(a.floydNext[u.ID()-1][v.ID()-1] + 1)
		path.Push(a.id2Name[u.ID()])
	}
	for !path.IsEmpty() {
		n, _ := path.Pop()
		ans = append(ans, n.(string))
	}
	return ans
}

func (a *AtelierData) findPath(start, dest string) ([]string, error) {
	s, err := a.getNode(start)
	if err != nil {
		return nil, errors.New("node " + start + " does not exist")
	}
	d, err := a.getNode(dest)
	if err != nil {
		return nil, errors.New("node " + dest + " does not exist")
	}
	path, _ := search.BreadthFirstSearch(s, d, a.graph)
	if path == nil {
		return nil, errors.New("there is no path")
	}

	var ans []string
	q := queue.NewQueue()
	for _, v := range path {
		q.Push(a.id2Name[v.ID()])
	}
	for !q.IsEmpty() {
		if q.Size() != 1 {
			n, _ := q.Pop()
			ans = append(ans, n.(string))
		} else {
			n, _ := q.Pop()
			ans = append(ans, n.(string))
		}
	}
	return ans, nil
}

func (a *AtelierData) printNodeList() (s string) {
	nodes := a.graph.NodeList()
	for _, v := range nodes {
		s += a.id2Name[v.ID()] + "\n"
	}
	return
}

// Returns a list of all the items (not types) in the dataset
func (a *AtelierData) GetItemList() (s []string) {
	nodes := a.graph.NodeList()
	for _, v := range nodes {
		if !strings.Contains(a.id2Name[v.ID()], "(") {
			s = append(s, a.id2Name[v.ID()])
		}
	}
	return
}

// Returns a list of all the items and item types in the dataset
func (a *AtelierData) GetAllItems() (s []string) {
	for _, v := range a.graph.NodeList() {
		s = append(s, a.id2Name[v.ID()])
	}
	return
}

// checks to see if the item is in the dataset
func (a *AtelierData) IsItem(s string) bool {
	_, v := a.name2Id[s]
	return v
}

func (a *AtelierData) usedIn(s string) string {
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
}

func (a *AtelierData) mapPre(n graph.Node, f func(graph.Node) string, depth int) string {
	if depth >= 1 {
		s := f(n)
		for _, v := range a.graph.Predecessors(n) {
			s += a.mapPre(v, f, depth-1)
		}
		return s
	}
	return f(n)
}

func (a *AtelierData) findNode(s string) (graph.Node, bool) {
	v, ok := a.name2Id[s]
	if ok {
		for _, n := range a.graph.NodeList() {
			if n.ID() == v {
				return n, true
			}
		}
	}
	return nil, false
}

func (a *AtelierData) findNodeID(i int) (graph.Node, bool) {
	_, ok := a.id2Name[i]
	if ok {
		for _, n := range a.graph.NodeList() {
			if n.ID() == i {
				return n, true
			}
		}
	}
	return nil, false
}

func (a *AtelierData) GetItemData(s string) (i Item) {
	n, ok := a.findNode(s)
	if ok {
		i.Name = s
		for _, r := range a.graph.Predecessors(n) {
			if strings.Contains(a.id2Name[r.ID()], "(") {
				i.Types = append(i.Types, a.id2Name[r.ID()])
			}
		}
		for _, r := range a.graph.Successors(n) {
			i.Ingredients = append(i.Ingredients, a.id2Name[r.ID()])
		}
		return
	}
	return
}*/
