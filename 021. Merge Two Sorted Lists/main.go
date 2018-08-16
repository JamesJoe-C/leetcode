package main

import (
	"fmt"
)

func main() {

	a := &ListNode{Val: 1}
	a.Next = &ListNode{Val: 2}
	a.Next.Next = &ListNode{Val: 4}
	//printList(a)

	b := &ListNode{Val: 1}
	b.Next = &ListNode{Val: 3}
	b.Next.Next = &ListNode{Val: 4}
	//printList(b)

	fmt.Println(mergeTwoLists(a, b))
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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var return_listNode, node *ListNode

	if l1.Val > l2.Val {
		return_listNode = l2
		node = l2
		l2 = l2.Next
	} else {
		return_listNode = l1
		node = l1
		l1 = l1.Next
	}

	for l2 != nil && l1 != nil {
		if l1.Val > l2.Val {
			node.Next = l2
			l2 = l2.Next
		} else {
			node.Next = l1
			l1 = l1.Next
		}
		node = node.Next
	}

	if l1 != nil {
		node.Next = l1
	}

	if l2 != nil {
		node.Next = l2
	}

	return return_listNode
}

func printList(listnode *ListNode) {
	var ll *ListNode = listnode
	fmt.Println("遍历：")
	for ll != nil {
		fmt.Println(ll)
		ll = ll.Next
	}
}
