package main

import "fmt"

func main() {
	arrs := []int{5, 1, 2}
	fmt.Printf("Sorted Array:%+v", selectionSort(arrs))
}

func selectionSort(arrs []int) []int {
	for i := range arrs {
		minIndex := i
		for j := range arrs[i+1:] {
			targetIndex := j + i + 1
			if arrs[targetIndex] < arrs[minIndex] {
				minIndex = targetIndex
			}
		}
		// swap
		temp := arrs[i]
		arrs[i] = arrs[minIndex]
		arrs[minIndex] = temp
	}

	return arrs
}
