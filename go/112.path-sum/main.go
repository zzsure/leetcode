package main

import "fmt"

func main() {
	n1 := &TreeNode{
		Val: 5,
	}
	n2 := &TreeNode{
		Val: 4,
	}
	n3 := &TreeNode{
		Val: 8,
	}
	n4 := &TreeNode{
		Val: 11,
	}
	n5 := &TreeNode{
		Val: 13,
	}
	n6 := &TreeNode{
		Val: 4,
	}
	n7 := &TreeNode{
		Val: 7,
	}
	n8 := &TreeNode{
		Val: 2,
	}
	n9 := &TreeNode{
		Val: 1,
	}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n3.Left = n5
	n3.Right = n6
	n4.Left = n7
	n4.Right = n8
	n6.Right = n9
	fmt.Println(hasPathSum(n1, 23))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}
	if hasPathSum(root.Left, sum-root.Val) == true {
		return true
	}
	return hasPathSum(root.Right, sum-root.Val)
}
