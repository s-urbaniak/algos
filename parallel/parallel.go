package parallel

import (
	"runtime"

	"github.com/s-urbaniak/mergesort"
)

func ParallelMergeSort(A []uint64) mergesort.MergeResult {
	np := runtime.NumCPU()
	lenA := len(A)

	if lenA < np {
		np = lenA
	}

	slice_size := lenA / np
	rem := lenA % np

	presult := make([]mergesort.MergeResult, np)
	ready := make(chan bool)
	for i := 0; i < np; i++ {
		x := i * slice_size
		y := (i+1)*slice_size + rem

		go func(i int) {
			presult[i] = mergesort.MergeSort(A[x:y])
			ready <- true
		}(i)
	}

	for i := 0; i < np; i++ {
		<-ready
	}

	merged := mergesort.MergeResult{0, make([]uint64, 0)}
	for i := range presult {
		merged = mergesort.Merge(merged.A, presult[i].A, merged.Inv+presult[i].Inv)
	}

	return merged
}
