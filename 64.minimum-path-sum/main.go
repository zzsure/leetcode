package main

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	min := minPathSum(grid)
	println(min)
}

func minPathSum(grid [][]int) int {
	if nil == grid || 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	path := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		path[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i <= 0 && j <= 0 {
				path[i][j] = grid[i][j]
			} else if i <= 0 && j > 0 {
				path[i][j] = path[i][j - 1] + grid[i][j]
			} else if i > 0 && j <= 0 {
				path[i][j] = path[i - 1][j] + grid[i][j]
			} else {
				path[i][j] = minVal(path[i - 1][j], path[i][j - 1]) + grid[i][j]
			}
		}
	}
	return path[len(grid) - 1][len(grid[0]) - 1]
}

func minVal(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}
	return v1
}