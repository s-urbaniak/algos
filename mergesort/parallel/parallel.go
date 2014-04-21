package parallel

import (
	"runtime"

	"github.com/s-urbaniak/algos/mergesort"
)

type mergeResult struct {
	A   []int64
	Inv int64
}

func ParallelMergeSort(A []int64) ([]int64, int64) {
	np := runtime.NumCPU()
	lenA := len(A)

	if lenA < np {
		np = lenA
	}

	slice_size := lenA / np
	rem := lenA % np

	presult := make([]mergeResult, np)
	ready := make(chan bool)
	for i := 0; i < np; i++ {
		x := i * slice_size
		y := (i+1)*slice_size + rem

		go func(i int) {
			A, inv := mergesort.MergeSort(A[x:y])
			presult[i] = mergeResult{A, inv}
			ready <- true
		}(i)
	}

	for i := 0; i < np; i++ {
		<-ready
	}

	merged := mergeResult{make([]int64, 0), 0}
	for i := range presult {
		A, inv := mergesort.Merge(merged.A, presult[i].A, merged.Inv+presult[i].Inv)
		merged = mergeResult{A, inv}
	}

	return merged.A, merged.Inv
}
