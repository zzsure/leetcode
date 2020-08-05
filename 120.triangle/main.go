package main

import (
	"fmt"
)

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
	count := len(triangle)
	if count == 0 {
		return 0
	}
	dp := make([]int, len(triangle[count-1]))
	for i := 0; i < count; i++ {
		pre := dp[0]
		for j := 0; j < i+1; j++ {
			temp := dp[j]
			if j == 0 {
				dp[j] = dp[0] + triangle[i][j]
			} else if pre < dp[j] || j == i {
				dp[j] = pre + triangle[i][j]
			} else {
				dp[j] = dp[j] + triangle[i][j]
			}
			pre = temp
		}
		fmt.Println(dp)
	}
	min := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] < min {
			min = dp[i]
		}
	}
	return min
}
