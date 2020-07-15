package main

import "fmt"

func main() {
	// [1,2,2,3,3,null,null,4,4]
	n1 := &TreeNode{
		Val: 1,
	}
	n2 := &TreeNode{
		Val: 2,
	}
	n3 := &TreeNode{
		Val: 2,
	}
	n4 := &TreeNode{
		Val: 3,
	}
	n5 := &TreeNode{
		Val: 3,
	}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n4
	n3.Right = n5
	fmt.Println(isBalanced(n1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := depth(root.Left)
	right := depth(root.Right)
	if left == right || left == right+1 || left == right-1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left)
	right := depth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
