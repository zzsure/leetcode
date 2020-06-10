package main

import (
	"sort"
)

func main() {
	println(len(combinationSum([]int{2, 3, 6, 7}, 7)))
	println(len(combinationSum([]int{2, 3, 5}, 8)))
}

func combinationSum(candidates []int, target int) [][]int {
	var list [][]int
	var sol []int
	sort.Ints(candidates)
	dfs(candidates, &list, 0, target, sol)
	return list
}

func dfs(candidates []int, list *[][]int, start, target int, sol []int) {
	if target == 0 {
		res := make([]int, len(sol))
		copy(res, sol)
		*list = append(*list, res)
	}
	for i := start; i < len(candidates) && target >= candidates[i]; i++ {
		sol = append(sol, candidates[i])
		dfs(candidates, list, i, target-candidates[i], sol)
		sol = sol[0 : len(sol)-1]
	}
}
