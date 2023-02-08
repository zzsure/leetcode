package main

import "fmt"

func main() {
	trees := generateTrees(0)
	fmt.Println(len(trees))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	var res []*TreeNode
	res = preTrees(1, n)
	return res
}

func preTrees(i, j int) []*TreeNode {
	if i > j {
		return []*TreeNode{nil}
	}
	if i == j {
		return []*TreeNode{{Val: i}}
	}
	var res []*TreeNode
	for k := i; k <= j; k++ {
		lefts := preTrees(i, k-1)
		rights := preTrees(k+1, j)
		for _, left := range lefts {
			for _, right := range rights {
				res = append(res, &TreeNode{Val: k, Left: left, Right: right})
			}
		}
	}
	return res
}
