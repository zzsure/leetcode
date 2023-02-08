package main

import "fmt"

func main() {
	a := 3
	fmt.Println(climbStairs(a))
}

func climbStairs(n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[0] = 1
		} else if i == 1 {
			dp[1] = 2
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[n-1]
}
