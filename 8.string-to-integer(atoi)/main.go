package main

import (
	"math"
	"strings"
)

func main() {
	str := "42"
	println(myAtoi(str))
	str = "   -42"
	println(myAtoi(str))
	str = "4193 with words"
	println(myAtoi(str))
	str = "words and 987"
	println(myAtoi(str))
	str = "-91283472332"
	println(myAtoi(str))
	str = "  0000000000012345678"
	println(myAtoi(str))
	str = "2147483647"
	println(myAtoi(str))
}

func myAtoi(str string) int {
	var num int64
	str = strings.TrimLeft(str, " ")
	if str == "" {
		return 0
	}

	orgStr := str
	if str[0] == '+' || str[0] == '-' {
		str = str[1:]
	}

	idx := 0
	for _, ch := range str {
		ch = ch - '0'
		if ch > 9 || ch < 0 {
			break
		}
		idx++
		num = num*10 + int64(ch)
		if num >= math.MaxInt32 {
			break
		}
	}

	if orgStr[0] == '-' {
		num = -1 * num
	}

	if num >= math.MaxInt32 {
		return math.MaxInt32
	}
	if num <= math.MinInt32 {
		num = math.MinInt32
	}

	return int(num)
}
