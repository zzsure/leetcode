package main

import "fmt"

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	fmt.Println(searchMatrix(matrix, 13))
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i := 0
	j := m - 1
	row := 0
	for i <= j {
		k := (i + j) / 2
		if target == matrix[k][0] {
			return true
		}
		if k == m-1 {
			row = k
			break
		}
		if k < m-1 && target > matrix[k][0] && target < matrix[k+1][0] {
			row = k
			break
		}
		if target < matrix[k][0] {
			j = k - 1
		} else {
			i = k + 1
		}
	}

	n := len(matrix[0])
	i = 0
	j = n - 1
	for i <= j {
		k := (i + j) / 2
		if target == matrix[row][k] {
			return true
		}
		if target < matrix[row][k] {
			j = k - 1
		} else {
			i = k + 1
		}
	}
	return false
}
