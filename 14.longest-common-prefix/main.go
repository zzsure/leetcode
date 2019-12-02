package main

func main() {
	strs1 := []string{"flower","flow","flight"}
	println(longestCommonPrefix(strs1))
	strs2 := []string{"dog","racecar","car"}
	println(longestCommonPrefix(strs2))
	strs3 := []string{"aa", "a"}
	println(longestCommonPrefix(strs3))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	common := ""
	for i := 0; i < len(strs[0]); i++ {
		var isSame bool = true
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				isSame = false
				break
			}
			if strs[0][i] != strs[j][i] {
				isSame = false
				break
			}
		}
		if isSame == false {
			common = strs[0][0:i]
			break
		}
		if i == len(strs[0]) - 1 {
			common = strs[0]
		}
	}
	return common
}
