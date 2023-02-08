package main

import "encoding/json"

func main() {
	//matrix := [][]int{}
	matrix := [][]int{{1, 1}, {1, 1}, {1, 1}}
	//matrix := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
	res := pacificAtlantic(matrix)
	j, _ := json.Marshal(res)
	println(string(j))
}

func pacificAtlantic(matrix [][]int) [][]int {
	dpt := make([][]int, 0)
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return dpt
	}
	pdpt := make([][]int, len(matrix))
	adpt := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		pdpt[i] = make([]int, len(matrix[0]))
		adpt[i] = make([]int, len(matrix[0]))
	}
	for j := 0; j < len(matrix[0]); j++ {
		search(matrix, 0, j, -1, pdpt)
	}
	for i := 1; i < len(matrix); i++ {
		search(matrix, i, 0, -1, pdpt)
	}
	for j := 0; j < len(matrix[0]); j++ {
		search(matrix, len(matrix) - 1, j, -1, adpt)
	}
	for i := 0; i < len(matrix) - 1; i++ {
		search(matrix, i, len(matrix[0]) - 1, -1, adpt)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if pdpt[i][j] & adpt[i][j] == 1 {
				dpt = append(dpt, []int{i, j})
			}
		}
	}
	return dpt
}

func search(matrix [][]int, i, j, val int, dpt [][]int) {
	if i < 0 || j < 0 || len(matrix) == 0 || len(matrix[0]) == 0 || i >= len(matrix) || j >= len(matrix[0]) || dpt[i][j] == 1 {
		return
	}
	if matrix[i][j] >= val {
		dpt[i][j] = 1
		search(matrix, i - 1, j, matrix[i][j], dpt)
		search(matrix, i, j + 1, matrix[i][j], dpt)
		search(matrix, i + 1, j, matrix[i][j], dpt)
		search(matrix, i, j - 1, matrix[i][j], dpt)
	}
}
