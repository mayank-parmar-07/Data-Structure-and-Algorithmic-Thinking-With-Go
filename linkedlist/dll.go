package linkedlist

import "fmt"

type DoubleListNode struct {
	Val  int
	Next *DoubleListNode
	Prev *DoubleListNode
}

type DLL struct {
	Head *DoubleListNode
	Size int
}

func (dll *DLL) Display() {
	headNode := dll.Head
	for headNode != nil {
		fmt.Print(headNode.Val, "<-->")
		headNode = headNode.Next
	}
	fmt.Println()
}

func (dll *DLL) InsertAtBeginning(x int) {
	newNode := DoubleListNode{Val: x, Prev: nil, Next: nil}
	if dll.Head == nil {
		dll.Head = &newNode
	} else {
		newNode.Next = dll.Head
		dll.Head.Prev = &newNode
		dll.Head = &newNode
	}
	dll.Size += 1
}

func (dll *DLL) InsertAtEnd(x int) {
	newNode := DoubleListNode{Val: x, Prev: nil, Next: nil}
	if dll.Head == nil {
		dll.Head = &newNode
	} else {
		headNode := dll.Head
		for headNode.Next != nil {
			headNode = headNode.Next
		}
		headNode.Next = &newNode
		newNode.Prev = headNode
	}
	dll.Size += 1
}

func (dll *DLL) DeleteNode(x int) error {
	if dll.Head.Val == x {
		dll.Head = dll.Head.Next
		dll.Size -= 1
		return nil
	} else {
		headNode := dll.Head
		for headNode.Next != nil {
			if headNode.Next.Val == x {
				headNode.Next = headNode.Next.Next
				if headNode.Next != nil {
					headNode.Next.Prev = headNode
				}
				dll.Size -= 1
				return nil
			}
			headNode = headNode.Next
		}
		return fmt.Errorf("Element doesn't exists in the list")
	}
}
