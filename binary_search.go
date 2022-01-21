package main

import (
	"fmt"
)

func main() {
	arrs := []int{81, 199}
	searchElement := 199
	l := 0
	r := len(arrs) - 1
	fmt.Printf("**** results:%d\n", binarySearch(arrs, l, r, searchElement))
}

func binarySearch(arrs []int, l, r, searchElement int) int {
	midIndex := l + int((r-l)/2)
	if midIndex >= l {
		if arrs[midIndex] == searchElement {
			return midIndex
		}

		if arrs[midIndex] > searchElement {
			return binarySearch(arrs, l, midIndex-1, searchElement)
		}

		return binarySearch(arrs, midIndex+1, r, searchElement)
	}

	// not found element
	return -1
}
