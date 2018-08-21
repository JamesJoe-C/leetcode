package main

import (
	"fmt"
)

func main() {
	head := &ListNode{Val: 1, Next: nil}
	head.Next = &ListNode{Val: 1, Next: nil}
	head.Next.Next = &ListNode{Val: 2, Next: nil}
	head.Next.Next.Next = &ListNode{Val: 3, Next: nil}
	head.Next.Next.Next.Next = &ListNode{Val: 3, Next: nil}
	//ListPrint(head)

	hh := deleteDuplicates(head)
	ListPrint(hh)
}

func ListPrint(head *ListNode) {
	hd := head
	for hd != nil {
		fmt.Println(hd.Val)
		hd = hd.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	hd := head
	for hd != nil && hd.Next != nil {
		if hd.Val == hd.Next.Val {
			hd.Next = hd.Next.Next
			continue
		}
		hd = hd.Next
	}
	return head
}
