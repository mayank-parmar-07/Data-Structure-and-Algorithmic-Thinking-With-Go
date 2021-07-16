package linkedlist

import "fmt"

type CircularListNode struct {
	val  int
	next *CircularListNode
}

type CircularList struct {
	Front *CircularListNode
	Rear  *CircularListNode
	Size  int
}

func (cll *CircularList) Display() {
	if cll.Front != nil {
		fmt.Print(cll.Front.val, "->")
		second := cll.Front.next
		for second != cll.Front {
			fmt.Print(second.val, "->")
			second = second.next
		}
		fmt.Println()
	}
}

func (cll *CircularList) InsertAtEnd(x int) {
	newNode := CircularListNode{val: x, next: nil}
	if cll.Front == nil {
		cll.Front = &newNode
		cll.Rear = &newNode
		cll.Rear.next = cll.Front
	} else {
		newNode.next = cll.Front
		cll.Rear.next = &newNode
		cll.Rear = &newNode
	}
	cll.Size += 1
}

func (cll *CircularList) Delete(x int) error {
	if x == cll.Front.val {
		cll.Rear.next = cll.Front.next
		cll.Front = cll.Rear.next
		return nil
	} else {
		second := cll.Front.next
		for second != cll.Front {
			if second.next.val == x {
				if second.next == cll.Rear {
					second.next = second.next.next
					cll.Rear = second
				} else {
					second.next = second.next.next
				}
				cll.Size -= 1
				return nil
			}
			second = second.next
		}
		if second == cll.Front {
			return fmt.Errorf("Element not in CLL")
		} else {
			return nil
		}
	}
}

func (cll *CircularList) InsertAtXPos(x, pos int) error {
	newNode := CircularListNode{val: x, next: nil}
	if pos == 0 {
		newNode.next = cll.Front
		cll.Rear.next = &newNode
		cll.Front = &newNode
		cll.Size += 1
		return nil
	} else {
		prev := cll.Front
		curr := cll.Front.next
		for pos > 1 && curr != cll.Front {
			curr = curr.next
			prev = prev.next
			pos -= 1
		}
		if curr != cll.Front {
			newNode.next = prev.next
			prev.next = &newNode
			cll.Size += 1
			return nil
		} else if curr == cll.Front && pos == 1 {
			newNode.next = cll.Front
			cll.Rear.next = &newNode
			cll.Rear = &newNode
			cll.Size += 1
			return nil
		} else {
			return fmt.Errorf("Index too big")
		}
	}
}
