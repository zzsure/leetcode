package main

import "fmt"

func main() {
	nums := []int{2, 5, 9}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
	var idxs [][]int
	count := len(nums)
	for i := 0; i <= count; i++ {
		temp := combine(count, i)
		idxs = append(idxs, temp...)
	}

	res := make([][]int, len(idxs))
	for i, idx := range idxs {
		temp := make([]int, len(idx))
		for j, v := range idx {
			temp[j] = nums[v-1]
		}
		res[i] = temp
	}

	return res
}

func combine(n int, k int) [][]int {
	var res [][]int
	var temp []int
	dfs(&res, &temp, n, k, 1)
	return res
}

func dfs(res *[][]int, temp *[]int, n, k, d int) {
	if len(*temp) == k {
		dst := make([]int, k)
		copy(dst, *temp)
		//fmt.Println(dst)
		*res = append(*res, dst)
		return
	}
	for i := d; i <= n; i++ {
		*temp = append(*temp, i)
		dfs(res, temp, n, k, i+1)
		*temp = (*temp)[0 : len(*temp)-1]
	}
}
