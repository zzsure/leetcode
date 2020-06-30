package main

import "fmt"

func main() {
	n, k := 4, 9
	fmt.Println(getPermutation(n, k))
}

func getPermutation(n int, k int) string {
	used := make([]bool, n)
	var temp []byte
	return dfs(&temp, used, 0, n, k)
}

func dfs(temp *[]byte, used []bool, depth, n, k int) string {
	if depth == n {
		return string(*temp)
	}
	cur := factorial(n - depth - 1)
	for i := 0; i < n; i++ {
		if used[i] == true {
			continue
		}
		if cur < k {
			k = k - cur
			continue
		}
		*temp = append(*temp, '0'+byte(i+1))
		used[i] = true
		return dfs(temp, used, depth+1, n, k)
	}
	return ""
}

func factorial(n int) int {
	res := 1
	for n > 0 {
		res *= n
		n--
	}
	return res
}
