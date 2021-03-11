package algorithm

import (
	"math/rand"
)

type Arr []int

//交换位置
func (arr Arr) swap(i, j int){
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

//将数据划分边界、获取边界下标
func partition(arr Arr, l, r int)[]int{
	less := l -1
	more := r
	for l < more {
		if arr[l] < arr[r] {
			less++
			arr.swap(less, l)
			l++
		}else if arr[l] > arr[r] {
			more--
			arr.swap(more, l)
		}else {
			l++
		}
	}
	arr.swap(more, r)
	less++
	return []int{less, more}
}


func (arr Arr) QuickSort(l, r int){
	if l < r {
		arr.swap(l + rand.Intn(r-l), r)
		p := partition(arr, l, r)
		arr.QuickSort(l, p[0] - 1)
		arr.QuickSort(p[1]+1, r)
	}
}