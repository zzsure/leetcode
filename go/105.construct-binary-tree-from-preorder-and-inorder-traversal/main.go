package main

func main() {
	preorder := []int{1, 2, 3}
	inorder := []int{3, 2, 1}
	output(inorder)
	tree := buildTree(preorder, inorder)
	inorderTree(tree)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func output(s []int) {
	for _, v := range s {
		println(v)
	}
}

func inorderTree(tree *TreeNode) {
	if nil == tree {
		return
	}
	inorderTree(tree.Left)
	println(tree.Val)
	inorderTree(tree.Right)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	val := preorder[0]
	tree := &TreeNode{
		Val: val,
		Left: nil,
		Right: nil,
	}
	inIdx := searchIdx(inorder, val)

	tree.Left = buildTree(preorder[1:inIdx+1], inorder[0:inIdx])
	tree.Right = buildTree(preorder[inIdx+1:len(preorder)], inorder[inIdx+1:len(inorder)])

	return tree
}

func searchIdx(arr []int, num int) int {
	idx := -1
	for i, v := range arr {
		if num == v {
			idx = i
			break
		}
	}
	return idx
}