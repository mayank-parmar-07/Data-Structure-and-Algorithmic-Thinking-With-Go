package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Size int
}

func (ll *LinkedList) InsertFirst10ElementsToListAtBack() {
	for i := 0; i < 10; i++ {
		newNode := ListNode{Val: i, Next: nil}
		if ll.Head == nil {
			ll.Head = &newNode
		} else {
			tempNode := ll.Head
			for tempNode.Next != nil {
				tempNode = tempNode.Next
			}
			tempNode.Next = &newNode
		}
		ll.Size = ll.Size + 1
	}
}

func (ll *LinkedList) Display() {
	tempNode := ll.Head
	for tempNode != nil {
		fmt.Print(tempNode.Val, "-->")
		tempNode = tempNode.Next
	}
	fmt.Println()
}

func (ll *LinkedList) InsertFirst10ElementsToListAtStart() {
	for i := 0; i < 10; i++ {
		newNode := ListNode{Val: i, Next: nil}
		if ll.Head == nil {
			ll.Head = &newNode
		} else {
			newNode.Next = ll.Head
			ll.Head = &newNode
		}
		ll.Size = ll.Size + 1
	}
}

func (ll *LinkedList) InsertAtXPos(n, val int) error {
	newNode := ListNode{Val: val, Next: nil}
	if n == 0 {
		newNode.Next = ll.Head
		ll.Head = &newNode
		return nil
	} else {
		tempNode := ll.Head
		i := 1
		for ; i < n && tempNode != nil; i++ {
			tempNode = tempNode.Next
		}
		if tempNode == nil && i < n {
			return fmt.Errorf("List is of smaller size")
		} else {
			newNode.Next = tempNode.Next
			tempNode.Next = &newNode
			ll.Size += 1

		}
	}
	return nil
}

func (ll *LinkedList) DeleteNode(x int) error {
	if ll.Head == nil {
		return fmt.Errorf("Empty List")
	}
	head := ll.Head
	if head.Val == x {
		ll.Head = ll.Head.Next
		return nil
	}
	for head != nil {
		if head.Next.Val == x {
			break
		}
		head = head.Next
	}
	if head == nil {
		return fmt.Errorf("Element not in the List")
	}
	head.Next = head.Next.Next
	ll.Size -= 1
	return nil
}

func (ll *LinkedList) InsertAtEnd(newNode ListNode) {

	if ll.Head == nil {
		ll.Head = &newNode
		return
	}
	head := ll.Head
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &newNode
	return
}

func (ll *LinkedList) FindLast() *ListNode {
	if ll.Head == nil {
		return nil
	} else {
		head := ll.Head
		for head.Next != nil {
			head = head.Next
		}
		return head
	}
}
