package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	res := fourSum(nums, 0)
	for _, r := range res {
		for _, n := range r {
			println(n)
		}
	}
	nums = []int{-3, -2, -1, 0, 0, 1, 2, 3}
	res = fourSum(nums, 0)
	for _, r := range res {
		for _, n := range r {
			println(n)
		}
	}
}

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	length := len(nums)
	if length < 4 {
		return res
	}
	sort.Ints(nums)

	resMap := make(map[string]bool)
	numMap := make(map[int]int)
	for idx, n := range nums {
		numMap[n] = idx
	}

	for i := 0; i < length-3; i++ {
		for j := i + 1; j < length-2; j++ {
			for k := j + 1; k < length-1; k++ {
				negSum := target - nums[i] - nums[j] - nums[k]
				if negSum < nums[k] {
					break
				}
				if _, ok := numMap[negSum]; ok {
					if numMap[negSum] > k {
						key := numKey(nums[i], nums[j], nums[k], negSum)
						if _, ok := resMap[key]; !ok {
							res = append(res, []int{nums[i], nums[j], nums[k], negSum})
							resMap[key] = true
						}
					}
				}
			}
		}
	}
	return res
}

func numKey(i, j, k, h int) string {
	return fmt.Sprintf("%d_%d_%d_d", i, j, k, h)
}
