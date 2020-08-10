package main

import "fmt"

func main() {
	//s := "A man, a plan, a canal: Panama"
	//s := "race a car"
	s := "0P"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < len(s) && isValidChar(s[i]) == false {
			i++
		}
		for j >= 0 && isValidChar(s[j]) == false {
			j--
		}
		if i < len(s) && j >= 0 && s[i] != s[j] {
			if s[i] >= 'a' && s[i] <= 'z' && s[i]-32 == s[j] {
				continue
			}
			if s[i] >= 'A' && s[i] <= 'Z' && s[i]+32 == s[j] {
				continue
			}
			return false
		}
	}

	return true
}

func isValidChar(b byte) bool {
	if (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
		return true
	}
	return false
}
