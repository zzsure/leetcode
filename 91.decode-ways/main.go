package main

import "fmt"

func main() {
	// 1
	fmt.Println(numDecodings("110"))
	// 0
	fmt.Println(numDecodings("100"))
	// 0
	fmt.Println(numDecodings("012"))
	// 0
	fmt.Println(numDecodings("01"))
	// 0
	fmt.Println(numDecodings("0"))
	// 1
	fmt.Println(numDecodings("10"))
	// 0
	fmt.Println(numDecodings("00"))
	// 2
	fmt.Println(numDecodings("12"))
	// 3
	fmt.Println(numDecodings("226"))
	// 2
	fmt.Println(numDecodings("227"))
	// 0
	fmt.Println(numDecodings("230"))
}

func numDecodings(s string) int {
	count := len(s)
	if count == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	cur, pre := 1, 1

	for i := 1; i < count; i++ {
		tmp := cur
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			cur = cur + pre
		}
		pre = tmp
	}

	return cur
}
