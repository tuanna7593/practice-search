package main

import "fmt"

func main() {
	arrs := []int{1, 2, 4, 12, 5, 233, 20, 3}
	element := 20
	position := linearSearch(arrs, element)
	if position == -1 {
		fmt.Printf("*** Not found %d in list***\n", element)
	} else {
		fmt.Printf("*** Find %d at position:%d in list***\n", element, position)
	}
}

func linearSearch(arrs []int, element int) (position int) {
	// default value -1 mean not found
	position = -1

	for i := range arrs {
		if arrs[i] == element {
			position = i
			return
		}
	}

	return position
}
