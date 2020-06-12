package main

import (
	"sort"
)

func main() {
	//[1, 7],
	//[1, 2, 5],
	//[2, 6],
	//[1, 1, 6]
	// [1, 1, 2, 5, 6, 7, 10]
	println(len(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)))
	//[1,2,2],
	//[5]
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
		//fmt.Println(*temp)
		for _, r := range *ret {
			if len(r) == len(*temp) {
				i := 0
				for ; i < len(r); i++ {
					if r[i] != (*temp)[i] {
						break
					}
				}
				if i == len(r) {
					return
				}
			}
		}
		copy(sol, *temp)
		//fmt.Println(*temp)
		*ret = append(*ret, sol)
	}
	for i := start; i < len(*candidates) && (*candidates)[i] <= target; i++ {
		*temp = append(*temp, (*candidates)[i])
		dfs(candidates, ret, temp, i+1, target-(*candidates)[i])
		*temp = (*temp)[0 : len(*temp)-1]
	}
}
