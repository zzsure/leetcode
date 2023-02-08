package main

import "fmt"

func main() {
	n1 := &TreeNode{
		Val: 5,
	}
	n2 := &TreeNode{
		Val: 4,
	}
	n3 := &TreeNode{
		Val: 8,
	}
	n4 := &TreeNode{
		Val: 11,
	}
	n5 := &TreeNode{
		Val: 13,
	}
	n6 := &TreeNode{
		Val: 4,
	}
	n7 := &TreeNode{
		Val: 7,
	}
	n8 := &TreeNode{
		Val: 2,
	}
	n9 := &TreeNode{
		Val: 5,
	}
	n10 := &TreeNode{
		Val: 1,
	}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n3.Left = n5
	n3.Right = n6
	n4.Left = n7
	n4.Right = n8
	n6.Left = n9
	n6.Right = n10
	fmt.Println(pathSum(n1, 22))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	var temp []int
	preTree(root, sum, &temp, &res)
	return res
}

func preTree(root *TreeNode, sum int, temp *[]int, res *[][]int) {
	if root == nil {
		return
	}
	*temp = append(*temp, root.Val)
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			r := make([]int, len(*temp))
			copy(r, *temp)
			*res = append(*res, r)
		}
		*temp = (*temp)[0 : len(*temp)-1]
		return
	}
	preTree(root.Left, sum-root.Val, temp, res)
	preTree(root.Right, sum-root.Val, temp, res)
	*temp = (*temp)[0 : len(*temp)-1]
}
