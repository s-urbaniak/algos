package mst

import (
	"container/heap"
	"errors"
)

const inf = int(^uint(0) >> 1)

type Vertex struct {
	label string
	adj   []int // adjacent edges
	key   int   // the key value in the heap
	hIdx  int   // the position in the heap
	gIdx  int   // the position in the graph
}

type VertexHeap []*Vertex

type Edge struct {
	u, v, cost int
}

type Graph struct {
	e    []Edge
	v    []Vertex
	x    []bool
	xLen int
	heap VertexHeap
}

func (h VertexHeap) Len() int {
	return len(h)
}

func (h VertexHeap) Less(i, j int) bool {
	return h[i].key < h[j].key
}

func (h VertexHeap) Swap(i, j int) {
	h[i].hIdx, h[j].hIdx = j, i
	h[i], h[j] = h[j], h[i]
}

func (h *VertexHeap) Pop() interface{} {
	old := h
	n := len(*old)
	x := (*old)[n-1]
	x.hIdx = -1
	*h = (*old)[0 : n-1]

	for i, entry := range *h {
		entry.hIdx = i
	}

	return x
}

func (h *VertexHeap) Push(x interface{}) {
	he := x.(*Vertex)
	he.hIdx = len(*h)
	*h = append(*h, he)
}

func NewGraph(s int, E []Edge, V []Vertex) (*Graph, error) {
	if len(V) <= s {
		return nil, errors.New("invalid source vertex index")
	}

	// create h, add all vertices
	h := VertexHeap(make([]*Vertex, len(V)))
	for i, _ := range h {
		V[i].key = inf
		V[i].hIdx = i
		V[i].gIdx = i
		h[i] = &V[i]
	}

	// initialize keys for adjacent vertices of source vertex s
	for _, e := range V[s].adj {
		if E[e].IsSourceVertex(s) {
			v := &V[E[e].v]
			v.key = E[e].cost
		}
	}

	heap.Init(&h)
	// remove source vertex s from h
	heap.Remove(&h, V[s].HeapIndex())

	G := &Graph{E, V, make([]bool, len(V)), 0, h}
	G.AddToX(s)

	return G, nil
}

func (g *Graph) AddToX(v int) {
	g.x[v] = true
	g.xLen += 1
}

func (g Graph) NotInX(v int) bool {
	return !g.x[v]
}

func (g Graph) XequalsV() bool {
	return g.xLen == len(g.v)
}

func (g Graph) EdgesLen() int {
	return len(g.e)
}

func (g Graph) VertexLen() int {
	return len(g.v)
}

func (g Graph) MSTCost() int {
	c := 0

	fixW := func(e Edge, wIdx int) {
		if g.NotInX(wIdx) {
			w := &g.v[wIdx]
			w.key = min(w.key, e.cost)
			heap.Fix(&g.heap, w.hIdx)
		}
	}

	for !g.XequalsV() {
		vMin := heap.Pop(&g.heap).(*Vertex)
		c += vMin.key
		g.AddToX(vMin.gIdx)

		for _, i := range vMin.adj {
			e := g.e[i]
			fixW(e, e.u)
			fixW(e, e.v)
		}
	}

	return c
}

func NewVertex(label string, edges ...int) Vertex {
	return Vertex{label, edges, inf, -1, -1}
}

func (v *Vertex) AddAdjEdge(e int) {
	v.adj = append(v.adj, e)
}

func (v *Vertex) ClearHeapValues() {
	v.hIdx = -1
	v.key = inf
}

func (v Vertex) HeapIndex() int {
	return v.hIdx
}

func NewEdge(u, v, cost int) *Edge {
	return &Edge{u, v, cost}	
}

func (e Edge) MaxVertex() int {
	if e.u > e.v {
		return e.u
	}

	return e.v
}

func (e Edge) IsSourceVertex(u int) bool {
	return e.u == u
}

func (e Edge) GetU() int {
	return e.u
}

func (e Edge) GetV() int {
	return e.v
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
