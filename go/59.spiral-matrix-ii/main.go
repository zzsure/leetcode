package main

import "fmt"

func main() {
	n := 4
	fmt.Println(generateMatrix(n))
}

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	for k := 1; k <= n*n && left < right && top < bottom; {
		for j := left; j < right; j++ {
			res[top][j] = k
			k++
		}
		for j := top; j < bottom; j++ {
			res[j][right] = k
			k++
		}
		for j := right; j > left; j-- {
			res[bottom][j] = k
			k++
		}
		for j := bottom; j > top; j-- {
			res[j][left] = k
			k++
		}
		left++
		right--
		top++
		bottom--
	}
	if n%2 != 0 {
		res[n/2][n/2] = n * n
	}
	return res
}
