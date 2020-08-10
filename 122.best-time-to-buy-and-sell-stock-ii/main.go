package main

import "fmt"

func main() {
	//arr := []int{7, 1, 5, 3, 6, 4}
	//arr := []int{1, 2, 3, 4, 5}
	arr := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(arr))
}

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
