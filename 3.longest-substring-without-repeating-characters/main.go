package main

func main() {
	// 3
	str := "abcabcbb"
	println(lengthOfLongestSubstring(str))
	// 1
	str = "bbbbb"
	println(lengthOfLongestSubstring(str))
	// 3
	str = "pwwkew"
	println(lengthOfLongestSubstring(str))
	// 0
	str = ""
	println(lengthOfLongestSubstring(str))
	// 1
	str = " "
	println(lengthOfLongestSubstring(str))
	// 3
	str = "dvdf"
	println(lengthOfLongestSubstring(str))
	// 2
	str = "abba"
	println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	sidx := 0
	maxMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := maxMap[s[i]]; ok {
			if i-sidx > maxLen {
				maxLen = i - sidx
			}
			for j := sidx; j < maxMap[s[i]]; j++ {
				delete(maxMap, s[j])
			}
			sidx = maxMap[s[i]] + 1
		}
		maxMap[s[i]] = i
		if i+1 == len(s) && i+1-sidx > maxLen {
			maxLen = i + 1 - sidx
		}
	}
	return maxLen
}
