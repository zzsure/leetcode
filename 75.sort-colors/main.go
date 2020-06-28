package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	//nums := []int{1, 0}
	//nums := []int{2, 1, 2}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	count := len(nums)
	i, j := 0, count-1
	idx := 0
	for idx <= j {
		if nums[idx] == 0 {
			nums[i], nums[idx] = nums[idx], nums[i]
			i++
			idx++
		} else if nums[idx] == 2 {
			nums[j], nums[idx] = nums[idx], nums[j]
			j--
		} else {
			idx++
		}
	}
}
