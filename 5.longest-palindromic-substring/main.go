package main

func main() {
	s := "babad"
	println(longestPalindrome(s))
	s = "cbbd"
	println(longestPalindrome(s))
	s = ""
	println(longestPalindrome(s))
	s = "a"
	println(longestPalindrome(s))
	s = "aa"
	println(longestPalindrome(s))
	s = "ab"
	println(longestPalindrome(s))
	s = "ccc"
	println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	sidx, eidx := 0, 0
	maxLen := 0
	msidx, meidx := 0, 0
	for i := 0; i < len(s); i++ {
		if i == 0 {
			eidx = 1
			meidx = 1
		} else {
			if s[i] == s[i-1] {
				for j := 0; j < i; j++ {
					if i+j >= len(s) || s[i-1-j] != s[i+j] {
						sidx = i - j
						eidx = i + j
						break
					} else {
						sidx = i - j - 1
						eidx = i + j + 1
					}
				}
				if eidx-sidx > maxLen {
					msidx = sidx
					meidx = eidx
					maxLen = eidx - sidx
				}
			}
			if i+1 < len(s) && s[i+1] == s[i-1] {
				for j := 0; j < i; j++ {
					if i+j+1 >= len(s) || s[i-1-j] != s[i+j+1] {
						sidx = i - j
						eidx = i + j + 1
						break
					} else {
						sidx = i - j - 1
						eidx = i + j + 2
					}
				}
				if eidx-sidx > maxLen {
					msidx = sidx
					meidx = eidx
					maxLen = eidx - sidx
				}
			}
		}
	}
	return s[msidx:meidx]
}
