package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	//nums := []int{3, 2, 1, 0, 4}
	//nums := []int{2, 0}
	//nums := []int{0}
	//nums := []int{0, 2, 3}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	reach := 0
	count := len(nums)
	for i := 0; i < count-1; i++ {
		if reach >= i && nums[i]+i > reach {
			reach = nums[i] + i
		}
		if reach >= count-1 {
			return true
		}
	}
	return false
}
