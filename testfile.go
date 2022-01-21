package main

import "fmt"

func main() {
	unsortedArr := []int{1, 3, 5, 10, 2, 4, 9, 8, -1}
	fmt.Println(">>> Unsorted Array: ", unsortedArr)
	sortedArr := quickSort(unsortedArr)
	fmt.Println(">>> Sorted Array: ", sortedArr)
}

func insertionSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		key := arr[i] // value at current index
		prevIndex := i - 1
		for ; prevIndex >= 0; prevIndex-- {
			if arr[prevIndex] <= key {
				break
			}
			arr[prevIndex+1] = arr[prevIndex]
		}
		arr[prevIndex+1] = key
	}

	return arr
}

func bubbleSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if arr[i] > arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	return arr
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	lArr := []int{}
	rArr := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i] <= pivot {
			lArr = append(lArr, arr[i])
			continue
		}

		rArr = append(rArr, arr[i])
	}

	lArr = quickSort(lArr)
	rArr = quickSort(rArr)

	return append(append(lArr, pivot), rArr...)
}
