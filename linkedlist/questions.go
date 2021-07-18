package linkedlist

import (
	"fmt"
)

func (sll *LinkedList) CheckIfNilTerminated() string {
	head := sll.Head
	if head == nil {
		return "Empty list"
	}
	if head.Next == nil {
		return "No cycle"
	}
	if head.Next.Next == head {
		return "Cyclic"
	}
	forwardNode := head.Next.Next
	for forwardNode != nil && head != forwardNode {
		if forwardNode.Next == nil {
			forwardNode = forwardNode.Next
			break
		}
		forwardNode = forwardNode.Next.Next
		head = head.Next
	}
	if forwardNode == nil {
		return "No cycle"
	} else {
		return "Cycle"
	}
}

func (ll *LinkedList) ReverseList() (error, *ListNode) {
	if ll.Head == nil {
		return fmt.Errorf("Empty list"), nil
	}
	var prev *ListNode
	curr := ll.Head
	nextNode := curr.Next
	for curr != nil {
		curr.Next = prev
		prev = curr
		curr = nextNode
		if nextNode != nil {
			nextNode = nextNode.Next
		}
	}
	ll.Head = prev
	return nil, ll.Head
}

func ReversePrint(ll *ListNode) {
	if ll == nil {
		return
	} else {
		ReversePrint(ll.Next)
		fmt.Print(ll.Val, "<--")
	}
}

func (ll *LinkedList) CheckEvenOrOddLength() string {
	if ll.Head == nil {
		return "Empty list"
	} else {
		head := ll.Head
		for head != nil && head.Next != nil {
			head = head.Next.Next
		}
		if head == nil {
			return "Even length"
		} else {
			return "Odd length"
		}
	}
}

func (ll *LinkedList) FindMid() (error, int) {
	if ll.Head == nil {
		return fmt.Errorf("Empty list"), 0
	}
	if ll.Head.Next == nil {
		return nil, ll.Head.Val
	}
	nextNode := ll.Head.Next
	head := ll.Head
	for nextNode != nil && nextNode.Next != nil {
		head = head.Next
		nextNode = nextNode.Next.Next
	}
	return nil, head.Val

}

func (cll *CircularList) SplitIntoTwo() (error, *CircularListNode, *CircularListNode) {
	if cll.Front == nil {
		return fmt.Errorf("Empty String"), nil, nil
	} else {
		if cll.Front.next == cll.Front {
			return nil, cll.Front, nil
		} else {
			fast_ptr := cll.Front.next
			slow_ptr := cll.Front
			for fast_ptr != slow_ptr {
				fast_ptr = fast_ptr.next.next
				slow_ptr = slow_ptr.next
			}
			head2 := slow_ptr.next
			slow_ptr.next = nil
			return nil, cll.Front, head2
		}
	}
}

func (ll *LinkedList) PushAllEvensBeforeOdds() *ListNode {
	if ll.Head == nil {
		return nil
	} else {
		head := ll.Head
		next := ll.Head.Next

		for next != nil {
			if head.Val%2 == 0 {
				head = head.Next
				next = next.Next
			} else {
				for next != nil && next.Val%2 != 0 {
					next = next.Next
				}
				if next != nil {
					head.Val, next.Val = next.Val, head.Val
				}
			}
		}
	}
	return ll.Head
}

//List1={a1,a2,...an}withdata,reorderitto{a1,an,....}

func (ll *LinkedList) Reorder() *ListNode {
	if ll.Head == nil {
		return nil
	}
	start := ll.Head
	nextPtr := ll.Head.Next
	for nextPtr != nil && nextPtr.Next != nil {
		start = start.Next
		nextPtr = nextPtr.Next.Next
	}
	midptr := start
	reverseList := midptr.Next
	midptr.Next = nil
	NewLinkedlist := LinkedList{Head: reverseList}
	NewLinkedlist.ReverseList()
	ll.Display()
	NewLinkedlist.Display()
	head := ll.Head
	startReverse := NewLinkedlist.Head
	for head != nil {
		nextPtr = head.Next
		var startReverseNext *ListNode
		if startReverse != nil {
			head.Next = startReverse
			startReverseNext = startReverse.Next
			startReverse.Next = nextPtr
		}
		head = nextPtr
		startReverse = startReverseNext
	}
	return ll.Head
}

func (ll *LinkedList) ReverseinBatches(BatchSize int) {
	if ll.Head == nil {
		return
	}
	var prevNode *ListNode
	curr := ll.Head
	for curr != nil {
		count := BatchSize
		tempNode := curr
		for tempNode != nil && count > 1 {
			tempNode = tempNode.Next
			count -= 1
		}
		if tempNode == nil {
			tempList := LinkedList{Head: curr, Size: BatchSize}
			_, neworderedNode := tempList.ReverseList()
			if prevNode == nil {
				ll.Head = neworderedNode
			} else {
				prevNode.Next = neworderedNode
			}
			break
		} else {
			nextPtr := tempNode.Next
			tempNode.Next = nil
			tempList := LinkedList{Head: curr, Size: BatchSize}
			_, neworderedNode := tempList.ReverseList()

			if prevNode == nil {
				ll.Head = neworderedNode
			} else {
				prevNode.Next = neworderedNode
			}
			prevNode = tempList.FindLast()
			tempList.InsertAtEnd(*nextPtr)
			curr = nextPtr
		}
	}
}
