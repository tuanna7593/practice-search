package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (ll *LinkedList) Prepend(node *Node) {
	if node == nil {
		fmt.Printf("cannot add the empty node to the list")
		return
	}

	secondNode := ll.Head
	node.Next = secondNode
	ll.Head = node
	ll.Length++
}

func (ll *LinkedList) RemoveValue(value int) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Data == value {
		ll.Head = ll.Head.Next
		ll.Length--
		return
	}

	if ll.Head.Next == nil { // no exists the value in list
		return
	}

	previousItemRemove := ll.Head
	for previousItemRemove.Next.Data != value {
		if previousItemRemove.Next.Next == nil {
			return
		}

		previousItemRemove = previousItemRemove.Next
	}

	previousItemRemove.Next = previousItemRemove.Next.Next
	ll.Length--
}

// insert a data after the value given
// insert at the end list if not found the give value
func (ll *LinkedList) InsertAfterValue(data, afterValue int) {
	newNode := &Node{Data: data}
	if ll.Head == nil {
		ll.Head = newNode
		ll.Length++
		return
	}

	afterItem := ll.Head
	for afterItem.Data != afterValue {
		if afterItem.Next == nil {
			break
		}

		afterItem = afterItem.Next
	}

	if afterItem.Next == nil {
		afterItem.Next = newNode
		ll.Length++
		return
	}

	newNode.Next = afterItem.Next
	afterItem.Next = newNode
	ll.Length++
}

func (ll LinkedList) printList() {
	printItem := ll.Head
	for ll.Length != 0 {
		fmt.Println(printItem.Data, printItem.Next)
		printItem = printItem.Next
		ll.Length--
	}
	fmt.Println()
}

func main() {
	ll := LinkedList{}
	n1 := &Node{Data: 1}
	n2 := &Node{Data: 2}
	n3 := &Node{Data: 3}
	n4 := &Node{Data: 4}
	ll.Prepend(n1)
	ll.Prepend(n2)
	ll.Prepend(n3)
	ll.Prepend(n4)
	ll.printList()

	ll.RemoveValue(5)
	ll.printList()

	ll.InsertAfterValue(5, 4)
	ll.printList()
}
