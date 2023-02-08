package main

import "fmt"

func main() {
	fmt.Println(getRow(4))
}

func getRow(rowIndex int) []int {
	rowIndex = rowIndex + 1
	res := make([]int, rowIndex)
	for i := 1; i <= rowIndex; i++ {
		if i == 1 {
			res[0] = 1
			continue
		}
		pre := 1
		for j := 1; j < i; j++ {
			if j == i-1 {
				res[j] = 1
				break
			}
			temp := res[j] + pre
			pre = res[j]
			res[j] = temp
		}
	}
	return res
}
