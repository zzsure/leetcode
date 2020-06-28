package main

import "fmt"

func main() {
	//a := 4
	//a := 8
	a := 1
	fmt.Println(mySqrt(a))
}

func mySqrt(x int) int {
	i, j := 0, x
	for i <= j {
		k := (i + j) / 2
		a := k * k
		b := (k + 1) * (k + 1)
		if a <= x && b > x {
			return k
		}
		if a < x {
			i = k + 1
		} else {
			j = k - 1
		}
	}
	return 0
}
