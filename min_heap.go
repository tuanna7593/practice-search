package main

import "fmt"

type MinHeap struct {
	arr []int
}

func (m *MinHeap) Insert(key int) {
	m.arr = append(m.arr, key)
	m.minHeapifyUp(len(m.arr) - 1)
}

func (m *MinHeap) Extract() int {
	if len(m.arr) == 0 {
		fmt.Println("cannot extract because the length of heap is zero")
		return -1
	}

	extracted := m.arr[0]
	lastIndex := len(m.arr) - 1

	m.arr[0] = m.arr[lastIndex]
	m.arr = m.arr[:lastIndex]

	m.minHeapifyDown(0)

	return extracted
}

// func (m *MinHeap) Extract() int {
// 	if len(m.arr) == 0 {
// 		fmt.Println("cannot extract because the length of heap is zero")
// 		return -1
// 	}

// 	extracted := m.arr[0]
// 	lastIndex := len(m.arr) - 1

// 	m.arr[0] = m.arr[lastIndex]
// 	m.arr = m.arr[:lastIndex]

// 	m.minHeapifyDown(0)

// 	return extracted
// }

func (m *MinHeap) minHeapifyUp(index int) {
	for m.arr[index] < m.arr[parent(index)] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

func (m *MinHeap) minHeapifyDown(index int) {
	lastIndex := len(m.arr) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if m.arr[l] < m.arr[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if m.arr[index] <= m.arr[childToCompare] {
			return
		}

		m.swap(index, childToCompare)
		index = childToCompare
		l, r = left(index), right(index)
	}
}

func (m *MinHeap) swap(i1, i2 int) {
	m.arr[i1], m.arr[i2] = m.arr[i2], m.arr[i1]
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

func main() {
	m := &MinHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for range buildHeap {
		fmt.Println(m.Extract())
		// fmt.Println(m)
	}
}
