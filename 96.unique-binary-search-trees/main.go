package main

import "fmt"

func main() {
	fmt.Println(numTrees(19))
}

func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		num := 0
		for j := 1; j <= i; j++ {
			num += dp[j-1] * dp[i-j]
		}
		dp[i] = num
	}
	return dp[n]
}
