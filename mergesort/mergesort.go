package mergesort

func Merge(A, B []int64, inv int64) ([]int64, int64) {
	C := make([]int64, 0, len(A)+len(B))
	i, j := 0, 0

	for (i < len(A)) && (j < len(B)) {
		if A[i] < B[j] {
			C = append(C, A[i])
			i++
		} else {
			C = append(C, B[j])
			j++
			inv += int64(len(A)) - int64(i)
		}
	}

	C = append(C, A[i:]...)
	C = append(C, B[j:]...)

	return C, inv
}

func MergeSort(A []int64) ([]int64, int64) {
	if len(A) == 1 {
		return A, 0
	}

	m := len(A) / 2
	l, lInv := MergeSort(A[:m])
	r, rInv := MergeSort(A[m:])

	return Merge(l, r, rInv+lInv)
}
