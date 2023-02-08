package main

func main() {
	str := "23"
	// ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
	arr := letterCombinations(str)
	for _, v := range arr {
		println(v)
	}
	str = ""
	arr = letterCombinations(str)
	for _, v := range arr {
		println(v)
	}
	str = "7"
	arr = letterCombinations(str)
	for _, v := range arr {
		println(v)
	}
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	m := map[string][]string{"2":[]string{"a","b","c"},"3":[]string{"d","e","f"},"4":[]string{"g","h","i"},"5":[]string{"j","k","l"},"6":[]string{"m","n","o"},"7":[]string{"p","q","r","s"},"8":[]string{"t","u","v"},"9":[]string{"w","x","y","z"}}
	if len(digits) == 1 {
		return m[digits]
	}

	sub := letterCombinations(digits[1:])
	strArr := m[string(digits[0])]
	newArr := make([]string, len(strArr)*len(sub))

	for i:=0; i<len(strArr); i++ {
		for k, v := range sub {
			n := strArr[i] + v
			newArr[i*len(sub)+k] = n
		}
	}

	return newArr
}
