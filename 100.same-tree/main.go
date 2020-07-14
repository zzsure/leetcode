package main

import "fmt"

func main() {
	n1 := &TreeNode{
		Val: 1,
	}
	n2 := &TreeNode{
		Val: 2,
	}
	n3 := &TreeNode{
		Val: 3,
	}
	n1.Right = n2
	n2.Left = n3

	n4 := &TreeNode{
		Val: 1,
	}
	n5 := &TreeNode{
		Val: 2,
	}
	n6 := &TreeNode{
		Val: 3,
	}
	n4.Right = n5
	n5.Left = n6
	fmt.Println(isSameTree(n1, n4))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if isSameTree(p.Left, q.Left) == false {
		return false
	}
	return isSameTree(p.Right, q.Right)
}
