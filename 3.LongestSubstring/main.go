package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	// "abcabcbb"=>"abc"
	// "bbbbb"=>"b"
	// "pwwkew"=>"wke"
	if s == "" {
		return 0
	}
	maxLen := 0
	letterMap := make(map[string]bool)
	for i, j := 0, 0; i < len(s) && j < len(s); {
		ch := s[i]
		if 
	}
	return 0
}
