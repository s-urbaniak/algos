package cp

import "math"

type Point struct {
	x, y float64
}

type Pair struct {
	p1, p2 Point
	d      float64
}

func Distance(p1, p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func NewPoints(data ...interface{}) []Point {
	if l := len(data); l%2 != 0 {
		panic("argument does not contain even number of entries")
	}

	P := make([]Point, len(data)/2)
	var x interface{}
	ip := 0
	for i, y := range data {
		if i%2 == 0 {
			x = y
		} else {
			var xf, yf float64

			switch x.(type) {
			case int:
				xf, yf = float64(x.(int)), float64(y.(int))
			case int64:
				xf, yf = float64(x.(int64)), float64(y.(int64))
			case float64:
				xf, yf = x.(float64), y.(float64)
			}

			P[ip] = NewPoint(xf, yf)
			ip++
		}
	}

	return P
}

func Merge(A, B []Point, cmp func(Point, Point) bool) []Point {
	C := make([]Point, 0, len(A)+len(B))
	i, j := 0, 0

	for (i < len(A)) && (j < len(B)) {
		if cmp(A[i], B[j]) {
			C = append(C, A[i])
			i++
		} else {
			C = append(C, B[j])
			j++
		}
	}

	C = append(C, A[i:]...)
	C = append(C, B[j:]...)

	return C
}

var xCmp = func(p1, p2 Point) bool {
	return p1.x < p2.x
}

var yCmp = func(p1, p2 Point) bool {
	return p1.y < p2.y
}

func MergeSort(A ...Point) (Px, Py []Point) {
	if len(A) < 2 {
		return A, A
	}

	m := len(A) / 2
	Lx, Ly := MergeSort(A[:m]...)
	Rx, Ry := MergeSort(A[m:]...)

	Px = Merge(Lx, Rx, xCmp)
	Py = Merge(Ly, Ry, yCmp)

	return Px, Py
}

func ClosestPairBruteForce(p ...Point) Pair {
	smallest := NewPair(p[0], p[1])

	for _, p1 := range p {
		for _, p2 := range p {
			current := NewPair(p1, p2)
			if (p1 != p2) && (current.d < smallest.d) {
				smallest = current
			}
		}
	}

	return smallest
}

func ClosestPair(Px, Py []Point) (best Pair) {
	if len(Px) != len(Py) {
		panic("Px, Py must have the same length")
	}

	if len(Px) < 4 {
		return ClosestPairBruteForce(Px...)
	}

	m := len(Px) / 2
	Qx, Rx := Px[:m], Px[m:]
	Qy, Ry := Py[:m], Py[m:]

	bq := ClosestPair(Qx, Qy)
	br := ClosestPair(Rx, Ry)
	delta := MinPair(bq, br)
	bs := ClosestSplitPair(Px, Py, delta)

	return MinPair(delta, bs)
}

func ClosestSplitPair(Px, Py []Point, delta Pair) Pair {
	xm := Px[(len(Px)/2)-1]
	min, max := xm.x-delta.d, xm.x+delta.d
	Sy := make([]Point, 0, len(Py))

	for _, x := range Px {
		if (x.x >= min) && (x.x <= max) {
			Sy = append(Sy, x)
		}
	}

	best := delta
	for i := 1; i <= len(Sy)-1; i++ {
		jmax := 7
		if len(Sy)-i < 7 {
			jmax = len(Sy) - i
		}

		for j := 1; j <= jmax; j++ {
			pq := NewPair(Sy[i-1], Sy[i+j-1])
			if pq.d < best.d {
				best = pq
			}
		}
	}

	return best
}

func MinPair(pairs ...Pair) Pair {
	min := pairs[0]
	for _, p := range pairs {
		if p.d < min.d {
			min = p
		}
	}
	return min
}

func NewPair(p1, p2 Point) Pair {
	return Pair{p1, p2, Distance(p1, p2)}
}

func NewPairVar(p ...Point) Pair {
	return Pair{p[0], p[1], Distance(p[0], p[1])}
}
