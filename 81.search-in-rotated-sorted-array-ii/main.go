package main

import "fmt"

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	fmt.Println(search(nums, 0))
	fmt.Println(search(nums, 3))
	//nums = []int{3, 1}
	//fmt.Println(search(nums, 1))
}

func search(nums []int, target int) bool {
	count := len(nums)
	i, j := 0, count-1
	for i <= j {
		k := (i + j) / 2
		if nums[k] == target || nums[i] == target || nums[j] == target {
			return true
		}
		if nums[i] == nums[k] {
			i++
		} else if nums[i] < nums[k] {
			if target > nums[i] && target < nums[k] {
				j = k - 1
			} else {
				i = k + 1
			}
		} else {
			if target > nums[k] && target < nums[j] {
				i = k + 1
			} else {
				j = k - 1
			}
		}
	}
	return false
}
