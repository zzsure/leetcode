package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		temp := make([]int, i)
		for j := 0; j < i; j++ {
			if j == 0 {
				temp[j] = 1
			} else if j == i-1 {
				temp[j] = 1
			} else {
				temp[j] = res[i-2][j] + res[i-2][j-1]
			}
		}
		res[i-1] = temp
	}
	return res
}
