package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	fmt.Println(countAndSay(6))
}

func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	if len(str) == 0 {
		return ""
	}
	lastChar := str[0]
	count := 0
	res := ""
	for i := 0; i < len(str); i++ {
		if str[i] != lastChar {
			res += strconv.Itoa(count) + string(lastChar)
			lastChar = str[i]
			count = 0
		}
		count++
		if i == len(str)-1 {
			res += strconv.Itoa(count) + string(lastChar)
		}
	}
	return res
}
