package scc

type firstPass struct {
	labels []int
	times  []int
}

func newFirstPass(G Graph) firstPass {
	labels := make([]int, len(G.vertex))
	for i, _ := range G.vertex {
		labels[len(labels)-i-1] = i
	}

	return firstPass{labels, make([]int, len(G.vertex))}
}

func (f firstPass) getLabels() []int {
	return f.labels
}

func (f firstPass) getTimes() []int {
	return f.times
}

func (f firstPass) time(v, t int) {
	// insert discovery time of vertex v in reverse order, i.e
	// time(0) = 4 becomes f.times[len(f.times)-4] = 0
	f.times[len(f.times)-t] = v
}

func (f firstPass) neighbour(v int, e Edge) (ok bool, v int) {
	if v == e.u {
		return true, e.v
	}

	return false, -1
}

func (_ firstPass) leader(int, int) {}
