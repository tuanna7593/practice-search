package main

import "fmt"

type MaxHeap struct {
	arrs []int
}

func (m *MaxHeap) Insert(data int) {
	m.arrs = append(m.arrs, data)
	m.maxHeapifyUp(len(m.arrs) - 1)
}

func (m *MaxHeap) Extract() {
	if len(m.arrs) == 0 {
		fmt.Println("Cannot extract an empty heap")
		return
	}
	extracted := m.arrs[0]
	fmt.Printf("Extracted:%d\n", extracted)

	lastIndex := len(m.arrs) - 1
	m.arrs[0] = m.arrs[lastIndex]
	m.arrs = m.arrs[:lastIndex]

}

func (m *MaxHeap) maxHeapifyUp(index int) {
	for m.arrs[parent(index)] < m.arrs[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return (2 * index) + 1
}

func right(index int) int {
	return (2 * index) + 2
}

func (m *MaxHeap) swap(a, b int) {
	m.arrs[a], m.arrs[b] = m.arrs[b], m.arrs[a]
}

func main() {
	m := &MaxHeap{}
	initArrs := []int{23, 5, 3, 10, 1, 30, 12, 4}
	for i := range initArrs {
		m.Insert(initArrs[i])
	}

	fmt.Println(m)
}
