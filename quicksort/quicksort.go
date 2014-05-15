package quicksort

// Strategy is function which returns a pivot index
type Strategy func([]int64) int

// Partition partitions a given []int64 slice between the indices x, y around the
// pivot index p
func Partition(A []int64, p int) int {
	// swap x <-> p
	A[0], A[p] = A[p], A[0]

	pv := A[0]
	i := 0

	for j := range A {
		if A[j] < pv {
			// swap i <-> j
			i += 1
			A[i], A[j] = A[j], A[i]
		}
	}

	// swap x <-> i
	A[0], A[i] = A[i], A[0]

	return i
}

// Quicksort sorts a given []int64 slice using a given pivot Strategy and returns
// the total amount of comparisons
func Quicksort(A []int64, s Strategy) int {
	if len(A) < 2 {
		return 0
	}

	p := Partition(A, s(A))

	cl := 0
	if p > 0 {
		cl = Quicksort(A[:p], s)
	}

	cr := 0
	if p < len(A) {
		cr = Quicksort(A[p+1:], s)
	}

	// the total number of comparisons is composed of the following ones:
	// 1. around the pivot in the current recursion (= len(A)-1)
	// 2. all recursions left to the pivot
	// 3. all recursions right to the pivot
	cmp := len(A) - 1 + cl + cr
	return cmp
}
