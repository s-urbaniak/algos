package mergesort

type MergeResult struct {
	Inv uint64
	A   []int
}

func Merge(A, B []int, Inv uint64) MergeResult {
	C := make([]int, 0, len(A)+len(B))
	i, j := 0, 0

	for (i < len(A)) && (j < len(B)) {
		if A[i] < B[j] {
			C = append(C, A[i])
			i++
		} else {
			C = append(C, B[j])
			j++
			Inv += uint64(len(A)) - uint64(i)
		}
	}

	C = append(C, A[i:]...)
	C = append(C, B[j:]...)

	r := MergeResult{Inv, C}
	return r
}

func MergeSort(A []int) MergeResult {
	if len(A) == 1 {
		return MergeResult{0, A}
	} else {
		m := len(A) / 2
		l := MergeSort(A[:m])
		r := MergeSort(A[m:])
		return Merge(l.A, r.A, r.Inv+l.Inv)
	}
}
