package mincut

import "math/rand"

// An edge is represented as tuple of adjacent vertices. The entries U, V point to vertex entries in Graph.V. The following variant must be maintained: u < v (thus also u != v).
type Edge struct {
	u, v    int
	deleted bool
}

// A vertex is represented as a sorted list of adjacent edges. Each entry points to an edge entry in Graph.E.
type Vertex struct {
	adj   []int
	label string
}

type Graph struct {
	// The vertices of the graph
	vertex []Vertex

	// The edges of the graph
	edge []Edge

	vertexCount int
	edgeCount   int
}

func NewEdge(u, v int) Edge {
	if u > v {
		return Edge{v, u, false}
	}

	return Edge{u, v, false}
}

// Swap swaps a source vertex index with the target vertex index and returns a new
// Edge. If the source vertex index is not present in the edge this method will
// panic.
func (e Edge) Swap(source, target int) Edge {
	if e.v == source {
		return NewEdge(target, e.u)
	}

	if e.u == source {
		return NewEdge(target, e.v)
	}

	panic("source vertex not present in edge")
}

func (e Edge) IsDeleted() bool {
	return e.deleted
}

func (e *Edge) Delete() {
	e.deleted = true
}

func NewVertex(label string) Vertex {
	return Vertex{make([]int, 0), label}
}

// AddEdges adds edge indices to the vertex's list of adjacent edges maintaining
// the invariant of a sorted list and returns the newly constructed vertex.
func (v Vertex) AddEdges(e ...int) Vertex {
	i, j := 0, 0
	w := make([]int, 0, len(v.adj)+len(e))

	for (i < len(v.adj)) && (j < len(e)) {
		if v.adj[i] < e[j] {
			w = append(w, v.adj[i])
			i++
		} else {
			w = append(w, e[j])
			j++
		}
	}

	w = append(w, v.adj[i:]...)
	w = append(w, e[j:]...)

	return Vertex{w, v.label}
}

func NewGraph(vertex []Vertex, edge []Edge) *Graph {
	return &Graph{vertex, edge, len(vertex), len(edge)}
}

func (g Graph) VertexCount() int {
	return g.vertexCount
}

func (g Graph) EdgeCount() int {
	return g.edgeCount
}

func (g Graph) GetEdge(e int) Edge {
	return g.edge[e]
}

// MinCut returns the minimun cut of the graph using Karger's algorithm.
func (g Graph) MinCut(iterations int) int {
	mins := make(chan (int))

	for i := 0; i < iterations; i++ {
		edgeCopy := make([]Edge, len(g.edge))
		vertexCopy := make([]Vertex, len(g.vertex))
		copy(edgeCopy, g.edge)
		copy(vertexCopy, g.vertex)

		go func(mins chan int, vertex []Vertex, edge []Edge) {
			gcopy := NewGraph(vertex, edge)

			for gcopy.VertexCount() > 2 {
				gcopy.Contract(gcopy.RandomEdge())
			}

			mins <- gcopy.EdgeCount()
		}(mins, vertexCopy, edgeCopy)
	}

	min := g.EdgeCount()
	for i := 0; i < iterations; i++ {
		m := <-mins
		if m < min {
			min = m
		}
	}

	return min
}

// Random returns a random edge index.
func (g *Graph) RandomEdge() int {
	return rand.Intn(g.EdgeCount())
}

// EdgeIndex returns the edge index of the given edge number e, where e <
// g.EdgeCount(). If e >= g.EdgeCount() this method will panic.
// Note: this method runs in O(n)!
func (g Graph) EdgeIndex(e int) int {
	i, ei := 0, -1

	for ei < e {
		if !g.edge[i].IsDeleted() {
			ei++
		}
		i++
	}

	return i - 1
}

// Contract contracts the graph at edge index e and modifies it accordingly.
//
// The given edge e points to vertices u, v. The vertices u, v maintain a sorted
// list of edge indices which are being merged into one new sorted list w of edge
// indices. The edge list of vertex u is replaced with the merged edge list w. The
// edge list of vertex v (and v itself) is being deleted.
//
// During the merge operation there are three cases:
// 1. An adjacent edge of vertex u is being merged.
// 2. An adjacent edge of vertex v is being merged.
// 3. A duplicate adjacent edge to u, v is being detected.
//
// For case 1. the adjacent edge of u is appended as is to the new edge list w.
// For case 2. the adjacent edge of v is added to the new edge list w but also
// modified such that it now points to u instead of v.
// For case 3. the adjacent duplicate edge of u, v is being deleted.
func (g *Graph) Contract(e int) {
	i, j := 0, 0
	e = g.EdgeIndex(e)

	// ui, vi are te
	ui, vi := g.edge[e].u, g.edge[e].v

	// edge indices from vertices u, v
	u, v := g.vertex[ui], g.vertex[vi]

	// w includes the merged edge indices
	w := make([]int, 0, len(u.adj)+len(v.adj))

	// merge indices from u and v
	for (i < len(u.adj)) && (j < len(v.adj)) {
		if u.adj[i] < v.adj[j] {
			w = append(w, u.adj[i])
			i++
		} else if u.adj[i] > v.adj[j] {
			ev := v.adj[j]
			g.edge[ev] = g.edge[ev].Swap(vi, ui)
			w = append(w, ev)
			j++
		} else {
			// self-loop, u and v point to the same edge
			loop := u.adj[i]
			g.edge[loop].Delete()
			g.edgeCount--
			i++
			j++
		}
	}

	// append remaining edges from u
	w = append(w, u.adj[i:]...)

	// append remaining edges from v
	for _, ev := range v.adj[j:] {
		g.edge[ev] = g.edge[ev].Swap(vi, ui)
		w = append(w, ev)
	}

	// u becomes the contracted vertex
	g.vertex[ui] = Vertex{w, u.label}
	g.vertex[vi] = Vertex{nil, v.label}
	g.vertexCount--
}
