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
	n1.Left = n2
	n1.Right = n3
	fmt.Println(levelOrderBottom(n1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		r := make([]int, len(nodes))
		for i, v := range nodes {
			r[i] = v.Val
		}
		res = append(res, r)
		var temp []*TreeNode
		for _, node := range nodes {
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		nodes = temp
	}
	for i, j := 0, len(res)-1; i < len(res)/2; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
