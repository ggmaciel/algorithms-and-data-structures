package main

import "fmt"

func quickSort(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)
	quickSort(arr, lo, pivotIdx-1)
	quickSort(arr, pivotIdx+1, hi)
}

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			tmp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}

	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx
}

func main() {
	arr := []int{9, 3, 7, 4, 58, 300, 42}
	fmt.Println("Before sorting: ", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("After sorting: ", arr)
}
