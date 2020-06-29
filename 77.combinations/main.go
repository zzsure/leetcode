package main

import "fmt"

func main() {
	n, k := 4, 2
	fmt.Println(combine(n, k))
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
