package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	left4 := &TreeNode{
		Val: 4,
		Left: nil,
		Right: nil,
	}
	left := &TreeNode{
		Val: 2,
		Left: left4,
		Right: nil,
	}
	right := &TreeNode{
		Val: 3,
		Left: nil,
		Right: nil,
	}
	root := &TreeNode{
		Val: 1,
		Left: left,
		Right: right,
	}
	//traverse(root)
	res := lcaDeepestLeaves(root)
	traverse(res)
}

func traverse(tree *TreeNode) {
	if tree == nil {
		return
	}
	println(tree.Val)
	traverse(tree.Left)
	traverse(tree.Right)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(tree *TreeNode) *TreeNode {
	if tree == nil {
		return nil
	}
	left := depth(tree.Left)
	right := depth(tree.Right)
	if left == right {
		return tree
	} else if left > right {
		return lcaDeepestLeaves(tree.Left)
	}
	return lcaDeepestLeaves(tree.Right)
}

func depth(tree *TreeNode) int {
	if tree == nil {
		return 0
	}
	left, right := 0, 0
	if tree.Left != nil {
		left = depth(tree.Left)	+ 1
	}
	if tree.Right != nil {
		right = depth(tree.Right) + 1
	}
	return 1 + getMax(left, right)
}

func getMax(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
