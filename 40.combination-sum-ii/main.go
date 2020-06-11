package main

import (
	"fmt"
	"sort"
)

func main() {
	println(len(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)))
	println(len(combinationSum2([]int{2, 5, 2, 1, 2}, 5)))
}

func combinationSum2(candidates []int, target int) [][]int {
	var ret [][]int
	var temp []int
	sort.Ints(candidates)
	dfs(&candidates, &ret, &temp, 0, target)
	return ret
}

func dfs(candidates *[]int, ret *[][]int, temp *[]int, start, target int) {
	if target == 0 {
		sol := make([]int, len(*temp))
		copy(sol, *temp)
		for _, r := range *ret {
			if len(r) == len(*temp) {
				for i := 0; i < len(r); i++ {
					if r[i] == (*temp)[i] {
						return
					}
				}
			}
		}
		*ret = append(*ret, sol)
		fmt.Println(*temp)
	}
	for i := start; i < len(*candidates) && (*candidates)[i] <= target; i++ {
		*temp = append(*temp, (*candidates)[i])
		dfs(candidates, ret, temp, i+1, target-(*candidates)[i])
		*temp = (*temp)[0 : len(*temp)-1]
	}
}
