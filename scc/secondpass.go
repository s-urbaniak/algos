package scc

import "sort"

type secondPass struct {
	labels  []int
	leaders []int
}

func newSecondPass(G Graph, labels []int) secondPass {
	return secondPass{labels, make([]int, len(G.vertex))}
}

func (s secondPass) getLabels() []int {
	return s.labels
}

func (_ *secondPass) time(int, int) {}

func (s *secondPass) neighbour(u int, e Edge) (v int, ok bool) {
	if u == e.v {
		return e.u, true
	}

	return -1, false
}

func (sp *secondPass) leader(_, s int) {
	sp.leaders[s] = sp.leaders[s] + 1
}

func (sp secondPass) TopFiveScc() []int {
	sort.Sort(sort.Reverse(sort.IntSlice(sp.leaders)))

	head := 5
	if len(sp.leaders) < head+1 {
		head = len(sp.leaders)
	}

	return sp.leaders[:head]
}
