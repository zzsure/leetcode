package main

import (
	"fmt"
)

func main() {
	a := []int{9, 3, 15, 20, 7}
	b := []int{9, 15, 7, 20, 3}
	root := buildTree(a, b)
	fmt.Println(root.Val)
	fmt.Println(root.Left.Val)
	fmt.Println(root.Right.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// left->cur->right
// left->right->cur

func buildTree(inorder []int, postorder []int) *TreeNode {
	inCount := len(inorder)
	postCount := len(inorder)
	if inCount == 0 || postCount == 0 || inCount != postCount {
		return nil
	}
	root := &TreeNode{
		Val: postorder[postCount-1],
	}
	inIdx := -1
	for i, v := range inorder {
		if v == root.Val {
			inIdx = i
		}
	}
	if inIdx == -1 {
		return nil
	}
	root.Left = buildTree(inorder[0:inIdx], postorder[0:inIdx])
	root.Right = buildTree(inorder[inIdx+1:inCount], postorder[inIdx:postCount-1])
	return root
}
