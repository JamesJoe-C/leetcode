package main

import (
	"fmt"
	"time"
)

func main() {

	a := ListNode{Val: 1}
	a.Next = &ListNode{Val: 2}
	a.Next.Next = &ListNode{Val: 4}
	printList(&a)

	b := ListNode{Val: 1}
	b.Next = &ListNode{Val: 3}
	b.Next.Next = &ListNode{Val: 4}
	printList(&b)

	fmt.Println(mergeTwoLists(&a, &b))
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
	var return_listNode ListNode
	var l1_next, l2_next ListNode = *l1, *l2
	return_listNode = *l1
	//for l1_next != nil {
	for &l2_next != nil {
		time.Sleep(1 * time.Second)
		if l2_next.Val >= l1_next.Val && (l1_next.Next == nil || l2_next.Val < l1_next.Next.Val) {
			l1_next.Next = l2_next

			l2_next = l2_next.Next
		} else {
			l1_next = l1_next.Next
			continue
		}
		fmt.Println(l1_next, l2_next)
	}

	//}

	return &return_listNode
}

func printList(listnode *ListNode) {
	var ll *ListNode = listnode
	fmt.Println("遍历：")
	for ll != nil {
		fmt.Println(ll)
		ll = ll.Next
	}
}
