package main

import "math"

func main() {
	//println(divide(10, 3))
	//println(divide(7, -3))
	//println(divide(-2147483648, -1))
	//println(divide(2147483647, 2))
	println(divide(-2147483648, -3))
}

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -dividend
		}
		return math.MaxInt32
	}
	if dividend == math.MinInt32 && divisor == math.MinInt32 {
		return 1
	} else if dividend > math.MinInt32 && divisor == math.MinInt32 {
		return 0
	}

	sign := 1
	remain := 0

	if dividend > 0 && divisor < 0 {
		sign = -1
		divisor = -divisor
	} else if dividend < 0 && divisor > 0 {
		sign = -1
		if dividend == math.MinInt32 {
			dividend = math.MaxInt32
			remain = 1
		} else {
			dividend = -dividend
		}
	} else if dividend < 0 && divisor < 0 {
		if dividend == math.MinInt32 {
			dividend = math.MaxInt32
			remain = 1
		} else {
			dividend = -dividend
		}
		divisor = -divisor
	}

	x, y := div(dividend, divisor)
	if y+remain == divisor {
		x = x + 1
	}

	if sign < 0 {
		x = -x
	}
	return x
}

func div(a, b int) (int, int) {
	if a < b {
		return 0, a
	}
	c := 1
	x := b
	for x+x <= a {
		c = c + c
		x = x + x
	}
	d, e := div(a-x, b)
	d = c + d
	return d, e
}
