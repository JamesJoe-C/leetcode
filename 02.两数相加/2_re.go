package main

//////////////////////////////////抄的，重点复习//////////
import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	current := head
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		fmt.Println(sum)
		carry = sum / 10
		current.Next = new(ListNode)
		current.Next.Val = sum % 10
		current = current.Next
	}
	return head.Next
}

// func main() {
// 	l111 := ListNode{
// 		Val:  3,
// 		Next: nil,
// 	}
// 	l11 := ListNode{
// 		Val:  4,
// 		Next: &l111,
// 	}

// 	l1 := &ListNode{
// 		Val:  2,
// 		Next: &l11,
// 	}

// 	l222 := &ListNode{
// 		Val:  4,
// 		Next: nil,
// 	}
// 	l22 := &ListNode{
// 		Val:  6,
// 		Next: l222,
// 	}
// 	l2 := &ListNode{
// 		Val:  5,
// 		Next: l22,
// 	}
// 	addTwoNumbers(l1, l2)

// 	//twoSum([]int{3,3}, 6)
// 	//fmt.Printf("reverse:%d", reverse(120))
// }
