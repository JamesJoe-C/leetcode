package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	p := TreeNode{Val: 111, Left: &TreeNode{Val: 222}, Right: &TreeNode{Val: 2122}}
	q := TreeNode{Val: 111, Left: &TreeNode{Val: 222}, Right: &TreeNode{Val: 2122}}
	temp := isSameTree(&p, &q)
	fmt.Println(temp)
}

/**
 * Definition for a binary tree node.
 * type TreeNtemp
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {

	temp := true

	if p == nil || q == nil {
		return false
	}
	if p == nil && q == nil {
		return true
	}

	// return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

	if p.Val != q.Val {
		return false
	}

	if p.Left != nil && q.Left != nil {
		temp = isSameTree(p.Left, q.Left)
	} else if p.Left == nil && q.Left == nil {
		//此种情况代表相等，不做处理
	} else {

		return false
	}

	if !temp {
		return temp
	}

	if p.Right != nil && q.Right != nil {
		temp = isSameTree(p.Right, q.Right)
	} else if p.Right == nil && q.Right == nil {
		//此种情况代表相等，不做处理
	} else {

		return false
	}

	if !temp {
		return temp
	}

	return temp
}
