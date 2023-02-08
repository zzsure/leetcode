package main

import "fmt"

func main() {
	//[10,5,15,null,null,6,20]
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
	fmt.Println(isValidBST(n1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var temp []int
	return inOrder(root, &temp)
}

func inOrder(root *TreeNode, temp *[]int) bool {
	if root == nil {
		return true
	}
	l := inOrder(root.Left, temp)
	if l == false {
		return false
	}
	if len(*temp) > 0 && root.Val <= (*temp)[len(*temp)-1] {
		return false
	}
	*temp = append(*temp, root.Val)
	return inOrder(root.Right, temp)
}
