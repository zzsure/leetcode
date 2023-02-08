package main

import "fmt"

func main() {
	//s := []int{2, 2, 1}
	s := []int{4, 1, 2, 1, 2}
	sn := singleNumber(s)
	fmt.Println(sn)
}

func singleNumber(nums []int) int {
	ret := 0
	for _, num := range nums {
		ret = ret ^ num
	}
	return ret
}
