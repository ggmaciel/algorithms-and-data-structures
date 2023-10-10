package main

import "fmt"

func main() {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(binarySearch(haystack, 0))  // false
	fmt.Println(binarySearch(haystack, 1))  // true
	fmt.Println(binarySearch(haystack, 3))  // true
	fmt.Println(binarySearch(haystack, 55)) // false
	fmt.Println(binarySearch(haystack, 13)) // false
	fmt.Println(binarySearch(haystack, 5))  // true
	fmt.Println(binarySearch(haystack, 6))  // true
	fmt.Println(binarySearch(haystack, 7))  // true
	fmt.Println(binarySearch(haystack, 8))  // true
}

func binarySearch(haystack []int, needle int) bool {
	lo := 0
	hi := len(haystack) - 1

	for lo < hi {
		middle := lo + (hi-lo)/2
		value := haystack[middle]

		if value == needle {
			return true
		} else if value > needle {
			hi = middle
		} else {
			lo = middle + 1
		}
	}

	return false
}
