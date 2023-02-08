package main

import "fmt"

func main() {
	//digits := []int{1, 2, 3}
	//digits := []int{4, 3, 2, 1}
	//digits := []int{9, 9, 9, 9}
	//digits := []int{0}
	digits := []int{9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	count := len(digits)
	res := make([]int, count+1)
	carry := 1
	for i := count - 1; i >= 0; i-- {
		if carry == 0 {
			res[i+1] = digits[i]
		} else {
			n := digits[i] + carry
			if n >= 10 {
				carry = 1
				res[i+1] = n - 10
			} else {
				carry = 0
				res[i+1] = n
			}
		}
	}
	if carry == 1 {
		res[0] = 1
	} else {
		res = res[1:]
	}
	return res
}
