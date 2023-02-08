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
	fmt.Println(zigzagLevelOrder(n1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	nodes := []*TreeNode{root}
	dir := 0
	for len(nodes) > 0 {
		var r []int
		if dir == 0 {
			for _, n := range nodes {
				r = append(r, n.Val)
			}
		} else {
			for i := len(nodes) - 1; i >= 0; i-- {
				r = append(r, nodes[i].Val)
			}
		}
		dir = ^dir
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
