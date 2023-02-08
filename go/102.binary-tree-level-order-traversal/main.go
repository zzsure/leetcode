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
		Val: 2,
	}
	n1.Left = n2
	n1.Right = n3
	fmt.Println(levelOrder(n1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		var r []int
		for _, n := range nodes {
			r = append(r, n.Val)
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
	return res
}
