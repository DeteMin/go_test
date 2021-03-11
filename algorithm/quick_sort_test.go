package algorithm

import (
	"fmt"
	"testing"
)

func TestArr_QuickSort(t *testing.T) {
	var arr Arr = []int{
		5,2,7,1,3,23,67,4,90,12,
	}
	arr.QuickSort(0, len(arr) -1)

	for _, v := range arr {
		fmt.Print(fmt.Sprintf("%v ", v))
	}
}
