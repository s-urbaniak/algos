package main

import "fmt"

const rows = 8

func cp(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func fromTo(x, y int) [][]int {
	res := make([][]int, rows)
	for i := x; i <= y; i++ {
		res[i-x] = []int{i}
	}
	return res
}

func check(i, j, m, n int) bool {
	return (j == n) || (i+j == m+n) || (i-j == m-n)
}

func safe(p []int, n int) bool {
	safe := true
	m := len(p) + 1
	for i, j := range p {
		safe = safe && !check(i+1, j, m, n)
	}
	return safe
}

func queens(m int) [][]int {
	if m == 1 {
		return fromTo(1, rows)
	}

	ps := queens(m - 1)
	res := make([][]int, 0)
	for _, p := range ps {
		for n := 1; n <= rows; n++ {
			if safe(p, n) {
				safep := append(cp(p), n)
				res = append(res, safep)
			}
		}
	}
	return res
}

func main() {
	fmt.Println("8 queens problem using the 'lists of successes' method\n", queens(8))
}
