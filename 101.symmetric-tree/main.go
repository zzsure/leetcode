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
	n4 := &TreeNode{
		Val: 3,
	}
	n5 := &TreeNode{
		Val: 3,
	}
	n1.Left = n2
	n1.Right = n3
	n2.Right = n4
	n5.Right = n5
	// [1,2,2,null,3,null,3]   false
	fmt.Println(isSymmetric(n1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		var temp []*TreeNode
		for _, node := range nodes {
			if node != nil {
				temp = append(temp, node.Left)
				temp = append(temp, node.Right)
			}
		}
		nodes = temp
		if len(nodes)%2 != 0 {
			return false
		}
		for i, j := 0, len(nodes)-1; i < len(nodes)/2; i, j = i+1, j-1 {
			if nodes[i] == nil && nodes[j] != nil {
				return false
			}
			if nodes[i] != nil && nodes[j] == nil {
				return false
			}
			if nodes[i] == nil && nodes[j] == nil {
				continue
			}
			if nodes[i].Val != nodes[j].Val {
				return false
			}
		}
	}
	return true
}
