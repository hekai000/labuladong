package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node[T any] struct {
	val  T
	next *Node[T]
	prev *Node[T]
}

type DoublyListNode struct {
	Val        int
	Prev, Next *DoublyListNode
}

func NewNode[T any](prev *Node[T], element T, next *Node[T]) *Node[T] {
	return &Node[T]{
		val:  element,
		next: next,
		prev: prev,
	}
}

func createLinkedList(arr []int) *ListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return head
}

func traverseLinkedList() {
	head := createLinkedList([]int{1, 2, 3, 4, 5})

	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}

func insertFront() {
	head := createLinkedList([]int{1, 2, 3, 4, 5})

	newHead := &ListNode{Val: 0}

	newHead.Next = head
	head = newHead
}

func insertBack() {
	head := createLinkedList([]int{1, 2, 3, 4, 5})

	p := head
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &ListNode{Val: 6}
	fmt.Println(p)
}

func insertMiddle() {
	head := createLinkedList([]int{1, 2, 3, 4, 5})
	p := head
	for i := 0; i < 2; i++ {
		p = p.Next
	}
	newNode := &ListNode{Val: 66}
	newNode.Next = p.Next

	p.Next = newNode
}

func deleteNodeMiddle() {
	head := createLinkedList([]int{1, 2, 3, 4, 5})
	p := head

	for i := 0; i < 2; i++ {
		p = p.Next
	}

	p.Next = p.Next.Next
}

func deleteNodeBack() {
	head := createLinkedList([]int{1, 2, 3, 4, 5})
	p := head
	for p.Next.Next != nil {
		p = p.Next
	}
	p.Next = nil
}

func deleteNodeFront() {
	head := createLinkedList([]int{1, 2, 3, 4, 5})
	head = head.Next
}

func NewDoublyListNode(x int) *DoublyListNode {
	return &DoublyListNode{Val: x}
}

func CreateDoublyLinkedList(arr []int) *DoublyListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	head := &DoublyListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		newNode := NewDoublyListNode(arr[i])
		cur.Next = newNode
		newNode.Prev = cur
		cur = cur.Next
	}
	return head
}

func TraverseDoublyLinkedList() {
	head := CreateDoublyLinkedList([]int{1, 2, 3, 4, 5})
	var tail *DoublyListNode

	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
		tail = p
	}

	for p := tail; p != nil; p = p.Prev {
		fmt.Println(p.Val)
	}

}

func InsertFrontDoublyLinkedList() {
	head := CreateDoublyLinkedList([]int{1, 2, 3, 4, 5})
	newHead := NewDoublyListNode(0)
	newHead.Next = head
	head.Prev = newHead
	head = newHead
}

func InsertBackDoublyLinkedList() {
	head := CreateDoublyLinkedList([]int{1, 2, 3, 4, 5})
	newNode := NewDoublyListNode(0)
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = newNode
	newNode.Prev = tail
	tail = newNode

}

func InsertMiddleDoublyLinkedList() {
	head := CreateDoublyLinkedList([]int{1, 2, 3, 4, 5})
	newNode := NewDoublyListNode(0)
	p := head

	for i := 0; i < 2; i++ {
		p = p.Next
	}

	newNode.Next = p.Next
	newNode.Prev = p

	p.Next.Prev = newNode
	p.Next = newNode

}

func deleteMiddleNodeDoublyLinkedList() {
	head := CreateDoublyLinkedList([]int{1, 2, 3, 4, 5})

	p := head
	for i := 0; i < 2; i++ {
		p = p.Next
	}
	toDelete := p.Next

	p.Next = toDelete.Next
	if toDelete.Next != nil {
		toDelete.Next.Prev = p
	}

	toDelete.Next = nil
	toDelete.Prev = nil

}

func deleteHeadDoublyLinkedList() {
	head := CreateDoublyLinkedList([]int{1, 2, 3, 4, 5})
	toDelete := head
	head = head.Next
	head.Prev = nil

	toDelete.Next = nil
}

func deleteTailDoublyLinkedList() {
	head := CreateDoublyLinkedList([]int{1, 2, 3, 4, 5})

	p := head
	for p.Next != nil {
		p = p.Next
	}

	p.Prev.Next = nil
	p.Prev = nil
}
