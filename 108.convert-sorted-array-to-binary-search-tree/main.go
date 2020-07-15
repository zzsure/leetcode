package main

import "fmt"

func main() {
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(root.Val)
	fmt.Println(root.Left.Val)
	fmt.Println(root.Right.Val)
	fmt.Println(root.Left.Left.Val)
	fmt.Println(root.Right.Left.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	count := len(nums)
	if count == 0 {
		return nil
	}
	root := &TreeNode{
		Val: nums[count/2],
	}
	root.Left = sortedArrayToBST(nums[0 : count/2])
	root.Right = sortedArrayToBST(nums[count/2+1 : count])
	return root
}
