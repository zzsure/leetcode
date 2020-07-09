package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	count := len(nums)

	var res [][]int
	for i := 0; i <= count; i++ {
		temp := combine(nums, count, i)
		res = append(res, temp...)
	}

	return res
}

func combine(nums []int, n, k int) [][]int {
	var res [][]int
	var temp []int

	dfs(nums, &res, &temp, n, k, 0)

	return res
}

func dfs(nums []int, res *[][]int, temp *[]int, n, k, d int) {
	if len(*temp) == k {
		dst := make([]int, k)
		copy(dst, *temp)
		*res = append(*res, dst)
		return
	}
	for i := d; i < n; i++ {
		if i > d && nums[i] == nums[i-1] {
			continue
		}
		*temp = append(*temp, nums[i])
		dfs(nums, res, temp, n, k, i+1)
		*temp = (*temp)[0 : len(*temp)-1]
	}
}
