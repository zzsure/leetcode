package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myPow(2.0, 10))
	fmt.Println(myPow(2.1, 3))
	fmt.Println(myPow(2.0, -2))
	math.Pow(2.0, 2.0)
	fmt.Println(myPow(0.00001, 2147483647))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n > 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -1*n)
}

func quickMul(x float64, n int) float64 {
	res := 1.0
	for n > 0 {
		if n%2 == 1 {
			res *= x
		}
		x *= x
		n = n / 2
	}
	return res
}
