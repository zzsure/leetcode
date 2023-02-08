package main

import "strconv"

func main() {
	println(isPalindrome(121))
	println(isPalindrome(-121))
	println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i < len(s) / 2; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
