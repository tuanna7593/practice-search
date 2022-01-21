package main

import "fmt"

func main() {
	arrs := []int{1, 3, 20, 4, 9, 200, 3, -1, -10}
	fmt.Printf("*** Unsort Array:%+v\n", arrs)
	insertionSort(arrs)
	fmt.Printf("*** Sorted Array using Insertion Sort:%+v\n", arrs)
}

func insertionSort(arrs []int) {
	for i := 1; i < len(arrs); i++ {
		key := arrs[i]
		j := i - 1
		for ; j >= 0 && key < arrs[j]; j-- {
			arrs[j+1] = arrs[j]
		}
		arrs[j+1] = key
	}
}
