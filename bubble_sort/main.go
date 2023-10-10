package main

import "fmt"

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90, 1, 8, 75}

	fmt.Println("Before sorting:", arr)
	fmt.Println("After sorting:", bubbleSort(arr))
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}
