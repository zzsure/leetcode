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
	n1.Left = nil
	n1.Right = n2
	n2.Left = n3
	res := inorderTraversal(n1)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorder(root, &res)
	return res
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}
