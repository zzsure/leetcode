package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix[0][0])
	fmt.Println(matrix[0][1])
	fmt.Println(matrix[0][2])
	fmt.Println(matrix[1][0])
	fmt.Println(matrix[1][1])
	fmt.Println(matrix[1][2])
	fmt.Println(matrix[2][0])
	fmt.Println(matrix[2][1])
	fmt.Println(matrix[2][2])
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i, j := 0, n-1; i < j; i++ {
		for k := 0; k < j-i; k++ {
			matrix[i][n-1-j+k], matrix[i+k][j], matrix[n-1-i][j-k], matrix[n-1-i-k][n-1-j] = matrix[n-1-i-k][n-1-j], matrix[i][n-1-j+k], matrix[i+k][j], matrix[n-1-i][j-k]
		}
		j--
	}
}
