package main

import (
	"fmt"
	"learning/algos"
	"learning/linkedlist"
)

func main() {
	algos.TowersOfHanoi(3)
	results := make([]string, 0)
	algos.Generate_str(&results, 4, 0, "")
	fmt.Println(results)
	// matrix, M, N := algos.Input_two_dimensional_arr()
	// fmt.Println(matrix)
	// visited := make([][]bool, M)
	// regions := 0
	// for i := 0; i < M; i++ {
	// 	for j := 0; j < N; j++ {
	// 		if !visited[i][j] && matrix[i][j] == 1 {
	// 			algos.Check_connected_components(matrix, &visited, i, j, M, N)
	// 			regions += 1
	// 		}
	// 	}
	// }
	// fmt.Println("Region count is ", regions)

	// Initiate singly linked list
	list := linkedlist.LinkedList{Head: nil, Size: 0}
	headNode := linkedlist.ListNode{Val: 10, Next: nil}
	list.Head = &headNode
	list.InsertFirst10ElementsToListAtBack()
	list.Display()
	list.InsertFirst10ElementsToListAtStart()
	fmt.Println("")
	list.Display()
	fmt.Println("\n size is", list.Size)
	err := list.InsertAtXPos(1, 51)
	if err != nil {
		fmt.Println(err.Error())
	}
	list.Display()
	fmt.Println("")
	err = list.InsertAtXPos(0, 51)
	if err != nil {
		fmt.Println(err.Error())
	}
	list.Display()
	err = list.DeleteNode(51)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("\nAfter deletion")
	list.Display()
	fmt.Println()
	err = list.DeleteNode(0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("\nAfter deletion")
	list.Display()
	fmt.Println()
	err = list.InsertAtXPos(20, 1151)
	if err != nil {
		fmt.Println(err.Error())
	}
	list.Display()
	fmt.Println()
	err = list.DeleteNode(1151)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("\nAfter deletion")
	list.Display()
	fmt.Println()

	// Instantiate a Double linked List

	dll := linkedlist.DLL{Head: nil, Size: 0}
	for i := 0; i < 10; i++ {
		dll.InsertAtBeginning(i)
	}
	dll.Display()
	for i := 0; i < 10; i++ {
		dll.InsertAtEnd(i)
	}
	dll.Display()

	err = dll.DeleteNode(9)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("After deleting first element")
		dll.Display()
	}
	err = dll.DeleteNode(0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("After deleting middle element")
		dll.Display()
	}
	dll.InsertAtEnd(99)
	dll.Display()
	err = dll.DeleteNode(99)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("After deleting last element")
		dll.Display()
	}
	err = dll.DeleteNode(-12)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("After deleting last element")
		dll.Display()
	}

	// Working with circular linked list

	cll := linkedlist.CircularList{Front: nil, Rear: nil, Size: 0}
	for i := 1; i < 10; i++ {
		cll.InsertAtEnd(i)
	}
	cll.Display()

	//Delete starting node

	err = cll.Delete(1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		cll.Display()
	}

	//Delete Last node

	cll.InsertAtEnd(19)
	cll.Display()
	err = cll.Delete(19)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		cll.Display()
	}
	// Delete at mid

	err = cll.Delete(5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		cll.Display()
	}

	err = cll.Delete(19)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		cll.Display()
	}

	//Insertion to cll

	err = cll.InsertAtXPos(19, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		cll.Display()
	}
	err = cll.InsertAtXPos(78, 5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		cll.Display()
	}
	err = cll.InsertAtXPos(191, 9)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		cll.Display()
	}
	err = cll.InsertAtXPos(1345, 191991)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		cll.Display()
	}

	// Check for cycle
	fmt.Println(list.CheckIfNilTerminated())
	newList := linkedlist.LinkedList{Head: nil, Size: 0}
	fmt.Println(newList.CheckIfNilTerminated())
	node1 := linkedlist.ListNode{Val: 10, Next: nil}
	newList.Head = &node1
	node2 := linkedlist.ListNode{Val: 10, Next: nil}
	node3 := linkedlist.ListNode{Val: 10, Next: nil}
	node4 := linkedlist.ListNode{Val: 10, Next: nil}
	node5 := linkedlist.ListNode{Val: 10, Next: newList.Head}
	newList.Head.Next = &node2
	newList.Head.Next.Next = &node3
	newList.Head.Next.Next.Next = &node4
	newList.Head.Next.Next.Next.Next = &node5

	fmt.Println(newList.CheckIfNilTerminated())

	//Reverse a list
	list.Display()
	err = list.InsertAtXPos(12, 2345)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		list.Display()
	}
	err = list.ReverseList()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		list.Display()
	}

	linkedlist.ReversePrint(list.Head)
	fmt.Println("\n", list.CheckEvenOrOddLength())
	list.DeleteNode(9)
	list.Display()
	fmt.Println("\n", list.CheckEvenOrOddLength())

	err, ans := list.FindMid()
	fmt.Println("mid is ", ans)

	// Solve Jospheus
	visited := map[int]bool{}
	for i := 0; i < 100; i++ {
		visited[i] = false
	}
	algos.Josepheus(visited, 100, 2, 0, 100)
}
