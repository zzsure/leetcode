package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	count := len(nums)
	if count == 0 || count == 1 || count == 2 {
		return count
	}
	lv := nums[0]
	rc := 1
	cnt := 1
	for i := 1; i < count; i++ {
		if nums[i] == lv {
			cnt++
			if cnt <= 2 {
				rc++
			}
		} else {
			cnt = 1
			rc++
		}
		nums[rc-1] = nums[i]
		lv = nums[i]
	}
	return rc
}
