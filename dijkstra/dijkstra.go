package dijkstra

import (
	"container/heap"
	"errors"
)

const inf = ^uint(0)

type Edge struct {
	u, v int
	len  uint
}

type Vertex struct {
	vPos  int
	adj   []int
	label string
	dist  uint

	// heap invariants
	key  uint
	hPos int
}

type Graph struct {
	vertex []Vertex
	edge   []Edge
	heap   VertexHeap
}

type VertexHeap struct {
	heap []*Vertex
}

func NewVertex(i int, label string, edges ...int) Vertex {
	return Vertex{i, edges, label, inf, 0, -1}
}

func NewEdge(u, v int, len uint) Edge {
	return Edge{u, v, len}
}

func min(x, y uint) uint {
	if x < y {
		return x
	}
	return y
}

func NewGraph(V []Vertex, s int, E []Edge) (*Graph, error) {
	if len(V) <= s {
		return nil, errors.New("unknown source vertex")
	}

	// create heap, add all vertices
	H := make([]*Vertex, len(V))
	for i, _ := range H {
		// initialize heap invariants
		V[i].hPos = i
		V[i].key = inf
		V[i].dist = inf
		H[i] = &V[i]
	}

	// calculate keys for adjacent vertices of source vertex s
	for _, i := range V[s].adj {
		if E[i].u == s {
			v := E[i].v
			V[v].key = min(V[v].key, E[i].len)
		}
	}

	// remove source vertex s
	V[s].hPos = -1
	V[s].dist = 0
	H = append(H[:s], H[s+1:]...)

	return &Graph{V, E, VertexHeap{H}}, nil
}

func (g Graph) DoShortestPath() {
	heap.Init(&g.heap)
	for g.heap.Len() > 0 {
		w := heap.Pop(&g.heap).(*Vertex)
		w.dist = w.key

		for _, i := range w.adj {
			edge := g.edge[i]
			u := &g.vertex[edge.u]
			v := &g.vertex[edge.v]

			if (u.vPos == w.vPos) && (v.hPos >= 0) {
				v.key = min(v.key, w.key+edge.len)
				heap.Fix(&g.heap, v.hPos)
			}
		}
	}
}

func (g Graph) GetDist() []uint {
	dist := make([]uint, len(g.vertex))
	for i, v := range g.vertex {
		dist[i] = v.dist
	}
	return dist
}

func (h VertexHeap) Len() int {
	return len(h.heap)
}

func (h VertexHeap) Less(i, j int) bool {
	return h.heap[i].key < h.heap[j].key
}

func (h VertexHeap) Swap(i, j int) {
	h.heap[i].hPos, h.heap[j].hPos = j, i
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *VertexHeap) Pop() interface{} {
	old := h.heap
	n := len(old)
	x := old[n-1]
	x.hPos = -1
	h.heap = old[0 : n-1]

	for i, entry := range h.heap {
		entry.hPos = i
	}

	return x
}

func (h *VertexHeap) Push(x interface{}) {
	he := x.(*Vertex)
	he.hPos = len(h.heap)
	h.heap = append(h.heap, he)
}
