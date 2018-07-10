package algorithm

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var a = []int{34, 45, 3, 6, 76, 34, 46, 809, 92, 8}
	var b = []int{3, 6, 8, 34, 34, 45, 46, 76, 92, 809}
	MergeSort(a, 0, 9)
	if reflect.DeepEqual(a, b) {
		t.Log("QuickSort Algorithm is running succesfully!")
	} else {
		t.Error("QuickSort Algorithm runs into error")
	}
}