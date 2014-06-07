package scc

import "github.com/s-urbaniak/algos/graph"

type Edge struct {
	u, v int
}

type Vertex struct {
	adj   []int
	label string
}

type Graph struct {
	vertex []Vertex
	edge   []Edge
}

type updater interface {
	// time is being invoked when vertex v was discovered at time t
	time(v, t int)

	// leader is being invoked when a leader vertex l is being
	// determined for vertex v
	leader(v, l int)

	// neighbour must return the index of the neighbour of vertex v else false
	neighbour(u int, e Edge) (v int, ok bool)

	// returns the labels (vertex indices) over which the SCC algorithm iterates
	getLabels() []int
}

func NewGraph(vertex []Vertex, edge []Edge) *Graph {
	return &Graph{vertex, edge}
}

func (g Graph) Edges() chan graph.Edge {
	ch := make(chan graph.Edge)
	go func() {
		for _, e := range g.edge {
			ch <- e
		}
		close(ch)
	}()
	return ch
}

func (g Graph) GetVertices(e graph.Edge) (graph.Vertex, graph.Vertex) {
	ee, ok := e.(Edge)
	if !ok {
		panic("not an scc edge")
	}

	return g.vertex[ee.u], g.vertex[ee.v]
}

func NewVertex(label string) *Vertex {
	return &Vertex{make([]int, 0), label}
}

func (v *Vertex) AddEdgeIndex(e int) {
	v.adj = append(v.adj, e)
}

func (v *Vertex) SetLabel(label string) {
	v.label = label
}

func (v Vertex) GetLabel() string {
	return v.label
}

func NewEdge(u, v int) *Edge {
	return &Edge{u, v}
}

func (e Edge) GetU() int {
	return e.u
}

func (e Edge) GetV() int {
	return e.v
}

func (g Graph) dfsLoop(u updater) {
	time := 0
	explored := make([]bool, len(g.vertex))

	for _, i := range u.getLabels() {
		if !explored[i] {
			time = g.dfs(i, i, time, explored, u)
		}
	}
}

func (g Graph) dfs(i, s, time int, explored []bool, u updater) int {
	explored[i] = true

	u.leader(i, s)
	for _, e := range g.vertex[i].adj {
		if j, ok := u.neighbour(i, g.edge[e]); ok && !explored[j] {
			time = g.dfs(j, s, time, explored, u)
		}
	}

	time++
	u.time(i, time)
	return time
}

func (g Graph) Scc() []int {
	fp := newFirstPass(g)
	g.dfsLoop(fp)

	sp := newSecondPass(g, fp.getTimes())
	g.dfsLoop(sp)

	return sp.TopFiveScc()
}
