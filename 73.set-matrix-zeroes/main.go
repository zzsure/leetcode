package main

import "fmt"

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	//matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	//matrix := [][]int{{1, 0}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rz, cz := false, false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			cz = true
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			rz = true
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	if rz == true {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if cz == true {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
