package mergesort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergesort(*testing.T) {
	result := MergeSort([]uint64{7, 2, 6})

	b := []uint64{2, 6, 7}
	if !(reflect.DeepEqual(result.A, b)) {
		fmt.Println("oh noes")
	}
}

func TestMergesortSingle(*testing.T) {
	result := MergeSort([]uint64{7})

	b := []uint64{7}
	if !(reflect.DeepEqual(result.A, b)) {
		fmt.Println("oh noes")
	}
}

func TestMergesortEmpty(*testing.T) {
	result := MergeSort([]uint64{})

	b := []uint64{}
	if !(reflect.DeepEqual(result.A, b)) {
		fmt.Println("oh noes")
	}
}
