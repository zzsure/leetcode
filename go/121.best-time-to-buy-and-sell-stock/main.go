package main

func main() {
	//prices := []int{7, 1, 5, 3, 6, 4}
	prices := []int{7, 6, 4, 3, 1}
	println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	min := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] - min > profit {
			profit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return profit
}