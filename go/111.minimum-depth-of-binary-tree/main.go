package main

import "fmt"

func main() {
	n1 := &TreeNode{
		Val: 3,
	}
	n2 := &TreeNode{
		Val: 9,
	}
	n3 := &TreeNode{
		Val: 20,
	}
	n4 := &TreeNode{
		Val: 15,
	}
	n5 := &TreeNode{
		Val: 7,
	}
	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	fmt.Println(minDepth(n1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	minDepth := 0
	dfs(root, 0, &minDepth)
	return minDepth
}

func dfs(root *TreeNode, depth int, minDepth *int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if *minDepth == 0 || depth+1 < *minDepth {
			*minDepth = depth + 1
		}
	}
	dfs(root.Left, depth+1, minDepth)
	dfs(root.Right, depth+1, minDepth)
}
