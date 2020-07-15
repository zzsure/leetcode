package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return sortedArrayToBST(nums)
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
