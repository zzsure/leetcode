package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(grayCode(0))
}

func grayCode(n int) []int {
	res := make([]int, int(math.Pow(2, float64(n))))
	res[0] = 0
	for i := 1; i <= n; i++ {
		start := math.Pow(2, float64(i-1))
		end := math.Pow(2, float64(i))
		for j := start; j < end; j++ {
			res[int(j)] = res[int(end-j-1)] | (1 << (i - 1))
		}
	}
	return res
}
