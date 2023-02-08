package main

func main() {
	haystack := "hello"
	needle := "ll"
	println(strStr(haystack, needle))
	haystack = "aaaaa"
	needle = "bba"
	println(strStr(haystack, needle))
	needle = ""
	println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	index := -1
	for i, _ := range haystack {
		if i + len(needle) > len(haystack) {
			i = -1
			break
		}
		isHave := true
		for j, _ := range needle {
			if haystack[i+j] != needle[j] {
				isHave = false
				break
			}
		}
		if isHave {
			index = i
			break
		}
	}
	return index
}
