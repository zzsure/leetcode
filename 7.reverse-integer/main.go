package main

import (
	"math"
	"strconv"
)

func main() {
	println(reverse(123))
	println(reverse(-123))
	println(reverse(120))
	println(reverse(1534236469))
}

func reverse(x int) int {
	s := strconv.Itoa(x)
	if x < 0 {
		s = s[1:len(s)]
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	s = string(r)
	i, _ := strconv.Atoi(s)
	if x < 0 {
		i = -1*i
	}
	if i > math.MaxInt32 || i < math.MinInt32 {
		i = 0
	}
	return i
}
