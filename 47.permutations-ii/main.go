package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(len(permuteUnique([]int{1, 1, 2})))
}

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)

	var temp []int
	var res [][]int
	used := make([]bool, len(nums))
	dfs(nums, &used, 0, &temp, &res)
	return res
}

func dfs(nums []int, used *[]bool, depth int, temp *[]int, res *[][]int) {
	if depth == len(nums) {
		r := make([]int, len(*temp))
		copy(r, *temp)
		*res = append(*res, r)
		return
	}

	for i := 0; i < len(nums); i++ {
		if (*used)[i] == true {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] {
			continue
		}
		*temp = append(*temp, nums[i])
		(*used)[i] = true
		dfs(nums, used, depth+1, temp, res)
		*temp = (*temp)[0 : len(*temp)-1]
		(*used)[i] = false
	}
}
