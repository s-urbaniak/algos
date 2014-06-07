package dijkstra

import (
	"container/heap"
	"errors"
)

const maxInt = ^uint(0)

type Edge struct {
	u, v int
	len  uint
}

type Vertex struct {
	adj    []int
	label  string
	key    uint
	inHeap bool
	pos    int
}

type Graph struct {
	vertex []Vertex
	edge   []Edge
	heap   VertexHeap
}

type HeapEntry int

type VertexHeap struct {
	heap   []HeapEntry
	vertex []Vertex
	dist   []uint
}

func NewVertex(label string, edges ...int) Vertex {
	return Vertex{edges, label, 0, false, -1}
}

func NewEdge(u, v int, len uint) Edge {
	return Edge{u, v, len}
}

func NewGraph(V []Vertex, s int, E []Edge) (*Graph, error) {
	if len(V) <= s {
		return nil, errors.New("unknown source vertex")
	}

	// create heap, add all vertices
	H := make([]HeapEntry, len(V))
	for i, _ := range H {
		V[i].key = maxInt
		V[i].inHeap = true
		H[i] = HeapEntry(i)
	}

	// calculate keys for adjacent vertices of source vertex s
	for _, i := range V[s].adj {
		if E[i].u == s {
			v := E[i].v
			if E[i].len < V[v].key {
				V[v].key = E[i].len
			}
		}
	}

	// remove source vertex s
	H = append(H[:s], H[s+1:]...)
	V[s].inHeap = false

	A := make([]uint, len(V))
	h := VertexHeap{H, V, A}
	return &Graph{V, E, h}, nil
}

func min(x, y uint) uint {
	if x < y {
		return x
	}
	return y
}

func (g Graph) DoShortestPath() {
	heap.Init(&g.heap)
	for g.heap.Len() > 0 {
		wStar := heap.Pop(&g.heap).(HeapEntry)
		g.heap.dist[wStar] = g.vertex[wStar].key

		for _, i := range g.vertex[wStar].adj {
			u := g.edge[i].u
			v := g.edge[i].v

			if (u == int(wStar)) && (g.vertex[v].inHeap) {
				g.vertex[v].key = min(
					g.vertex[v].key,
					g.vertex[wStar].key+g.edge[i].len)
				heap.Fix(&g.heap, g.vertex[v].pos)
			}
		}
	}
}

func (g Graph) GetDist() []uint {
	return g.heap.dist
}

func (h VertexHeap) Len() int {
	return len(h.heap)
}

func (h VertexHeap) Less(i, j int) bool {
	return h.vertex[h.heap[i]].key < h.vertex[h.heap[j]].key
}

func (h VertexHeap) Swap(i, j int) {
	vi := &h.vertex[h.heap[i]]
	vj := &h.vertex[h.heap[j]]
	vi.pos = j
	vj.pos = i
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *VertexHeap) Pop() interface{} {
	old := h.heap
	n := len(old)
	x := old[n-1]

	h.vertex[x].inHeap = false
	h.heap = old[0 : n-1]

	for i, entry := range h.heap {
		h.vertex[entry].pos = i
	}

	return x
}

func (h *VertexHeap) Push(x interface{}) {
	he := x.(HeapEntry)
	h.vertex[he].pos = len(h.heap)
	h.heap = append(h.heap, he)
}
