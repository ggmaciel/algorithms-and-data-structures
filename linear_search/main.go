package main

import "fmt"

func main() {
	fmt.Println(linearSearch([]int{1, 2, 3, 4, 5}, 6))
	fmt.Println(linearSearch([]int{1, 2, 3, 4, 5}, 3))
}

func linearSearch(haystack []int, needle int) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}
