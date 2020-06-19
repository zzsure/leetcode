package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m := len(matrix)
	n := len(matrix[0])
	count := m * n
	order := make([]int, count)
	left, right, top, bottom := 0, n-1, 0, m-1
	idx := 0
	for left <= right && top <= bottom {
		for k := left; k <= right; k++ {
			order[idx] = matrix[top][k]
			idx++
		}
		for k := top + 1; k <= bottom; k++ {
			order[idx] = matrix[k][right]
			idx++
		}
		if left < right && top < bottom {
			for k := right - 1; k > left; k-- {
				order[idx] = matrix[bottom][k]
				idx++
			}
			for k := bottom; k > top; k-- {
				order[idx] = matrix[k][left]
				idx++
			}
		}

		left++
		right--
		top++
		bottom--
	}
	return order
}
