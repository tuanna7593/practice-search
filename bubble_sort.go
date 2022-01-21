package main

import "fmt"

func main() {
	arrs := []int{1, 10, 3, -1, 22, 25, 100, -19, 200, -100, -101, -102}
	fmt.Printf("*** Unsort list:%+v\n", arrs)
	bubbleSort(arrs)
	fmt.Printf("*** Bubble Sort:%+v", arrs)
}

func bubbleSort(arrs []int) {
	for i := range arrs {
		for j := i + 1; j < len(arrs); j++ {
			if arrs[i] > arrs[j] {
				arrs[i], arrs[j] = swap(arrs[i], arrs[j])
			}
		}
	}
}

func swap(a, b int) (int, int) {
	return b, a
}
