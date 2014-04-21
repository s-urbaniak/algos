package cp

import (
	"math"
	"testing"
)

func TestPair(t *testing.T) {
	p := NewPair(NewPoint(0, 0), NewPoint(1, 1))
	if p.d == 0 {
		t.Error("p.d is 0")
	}
}

func TestClosestPairBruteForce(t *testing.T) {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 3)
	p3 := NewPoint(5, 5)
	p4 := NewPoint(6, 6)

	if s := ClosestPairBruteForce(p1, p2, p3, p4); s != NewPair(p3, p4) {
		t.Errorf("wrong closest pair %v\n", s)
	}

	if s := ClosestPairBruteForce(p1, p2, p3); s != NewPair(p2, p3) {
		t.Errorf("wrong closest pair %v\n", s)
	}

	if s := ClosestPairBruteForce(p1, p2); s != NewPair(p1, p2) {
		t.Errorf("wrong closest pair %v\n", s)
	}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("error expected, got none")
		}
	}()
	ClosestPairBruteForce(p1)
}

func TestClosestPair(t *testing.T) {
	P := NewPoints(0, 0, 4, 4, 6, 6, 9, 9, 14, 14, 19, 19, 50, 50, 50.5, 50.5, 90, 90)

	e := NewPoints(50, 50, 50.5, 50.5)
	if s := ClosestPair(MergeSort(P...)); s != NewPair(e[0], e[1]) {
		t.Errorf("wrong closest pair %v\n", s)
	}
}

func TestNew(t *testing.T) {
	assertDistance(t, NewPoint(0, 0), NewPoint(1, 1), math.Sqrt2)
	assertDistance(t, NewPoint(0, 0), NewPoint(-1, -1), math.Sqrt2)
	assertDistance(t, NewPoint(-1, -1), NewPoint(0, 0), math.Sqrt2)
	assertDistance(t, NewPoint(-1, -1), NewPoint(1, 1), 2*math.Sqrt2)
	assertDistance(t, NewPoint(0, 0), NewPoint(0, 0), 0)
}

func assertDistance(t *testing.T, p1, p2 Point, expected float64) {
	d := Distance(p1, p2)
	if expected != d {
		t.Errorf("unexpected distance %v\n", d)
	}
}

func TestMergeSort(t *testing.T) {
	p1 := NewPoint(0, 9)
	p2 := NewPoint(1, 6)
	p3 := NewPoint(4, 8)
	p4 := NewPoint(2, 7)
	p5 := NewPoint(3, 5)

	Px, Py := MergeSort(p1, p2, p3, p4, p5)
	assertPoints(t, Px, []Point{p1, p2, p4, p5, p3})
	assertPoints(t, Py, []Point{p5, p2, p4, p3, p1})

	Px, Py = MergeSort(p1)
	assertPoints(t, Px, []Point{p1})
	assertPoints(t, Py, []Point{p1})
}

func assertPoints(t *testing.T, A, B []Point) {
	if len(A) != len(B) {
		t.Errorf("len(A) = %v is not equal len(B) = %v\n", len(A), len(B))
	}

	for i, p := range A {
		if B[i] != p {
			t.Errorf("%v != %v at index %v", B[i], p, i)
		}
	}
}

func TestNewPoints(t *testing.T) {
	p := NewPoints()
	assertPoints(t, p, make([]Point, 0))

	p = NewPoints(0, 0)
	assertPoints(t, p, []Point{NewPoint(0, 0)})

	p = NewPoints(0, 0, 1, 1)
	assertPoints(t, p, []Point{NewPoint(0, 0), NewPoint(1, 1)})
}

func TestMinPair(t *testing.T) {
	p1 := NewPairVar(NewPoints(0, 0, 3, 3)...)
	p2 := NewPairVar(NewPoints(0, 0, 2, 2)...)
	p3 := NewPairVar(NewPoints(0, 0, 1, 1)...)

	if p := MinPair(p1, p2, p3); p != p3 {
		t.Errorf("invalid min pair %v\n", p)
	}
}
