package main

import "fmt"

func main() {
	str := "Hello World"
	//str := "a "
	fmt.Println(lengthOfLastWord(str))
}

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}
	count := len(s)
	length := 0
	for i := count - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if length == 0 {
				continue
			}
			return length
		}
		length++
	}
	return length
}
