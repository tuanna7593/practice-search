package main

import (
	"fmt"
)

type MaxHeap struct {
	arr []int
}

func (m *MaxHeap) Insert(key int) {
	m.arr = append(m.arr, key)
	m.maxHeapifyUp(len(m.arr) - 1)
}

func (m *MaxHeap) Extract() int {
	if len(m.arr) == 0 {
		fmt.Println("cannot extract because the length of heap is 0")
		return -1
	}

	extracted := m.arr[0]
	lastIndex := len(m.arr) - 1

	m.arr[0] = m.arr[lastIndex]
	m.arr = m.arr[:lastIndex]

	m.maxHeapifyDown(0)

	return extracted
}

func (m *MaxHeap) Sort(method string) {
	if method == "ASC" {
		m.sortASC()
		return
	}
	m.sortDESC()
}

func (m MaxHeap) sortASC() {
	sortedArr := make([]int, len(m.arr))
	l := len(m.arr) - 1
	for i := l; i > -1; i-- {
		extracted := m.Extract()
		sortedArr[i] = extracted
	}
	for i := range sortedArr {
		fmt.Println(sortedArr[i])
	}
}

func (m MaxHeap) sortDESC() {
	sortedArr := make([]int, len(m.arr))
	l := len(m.arr)
	for i := 0; i < l; i++ {
		extracted := m.Extract()
		sortedArr[i] = extracted
	}
	for i := range sortedArr {
		fmt.Println(sortedArr[i])
	}
}

func (m *MaxHeap) maxHeapifyUp(index int) {
	for m.arr[parent(index)] < m.arr[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

func (m *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(m.arr) - 1
	childToCompare := index
	l := left(index)
	r := right(index)

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if m.arr[l] > m.arr[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if m.arr[index] >= m.arr[childToCompare] {
			return
		}

		m.swap(index, childToCompare)
		index = childToCompare
		l, r = left(index), right(index)
	}
}

func (m *MaxHeap) swap(i, j int) {
	m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
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
	m := &MaxHeap{}
	fmt.Println(m)

	buildHealp := []int{10, 20, 30, 5, 7}
	for _, v := range buildHealp {
		m.Insert(v)
		fmt.Println(m)
	}

	// for range buildHealp {
	// 	m.Extract()
	// 	fmt.Println(m)
	// }
	m.Sort("DESC")
}
