package algorithm

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T){
	var a = []int{13, 8, 10, 2, 88, 3, 19, 33, 101, 15}
	var b = []int{2, 3, 8, 10, 13, 15, 19, 33, 88, 101}
	QuickSort(a, 0, 9)
	if reflect.DeepEqual(a, b) {
		t.Log("QuickSort Algorithm is running succesfully!")
	} else {
		t.Error("QuickSort Algorithm runs into error")
	}
}