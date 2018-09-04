package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return l1
	}
	plus := l1.Val + l2.Val
	if plus < 10 {
		l1.Val = plus
	} else {
		if l1.Next == nil {
			l1.Next = &ListNode{}
		}
		l1.Val = plus % 10
		l1.Next.Val += plus / 10
	}
	if l1.Next != nil && l2.Next == nil {
		addTwoNumbers(l1.Next, &ListNode{})
	} else if l1.Next == nil && l2.Next != nil {
		l1.Next = &ListNode{}
		addTwoNumbers(l1.Next, l2.Next)
	} else {
		addTwoNumbers(l1.Next, l2.Next)
	}
	return l1
}

func main() {
	fmt.Println(addTwoNumbers(
		&ListNode{
			Val: 5,
		},
		&ListNode{
			Val: 5,
		},
	))
}
