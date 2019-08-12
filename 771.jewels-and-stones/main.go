package main

import "fmt"

func main() {
	//j := "aA"
	//s := "aAAbbbb"
	j := "z"
	s := "ZZ"
	fmt.Println(numJewelsInStones(j, s))
}

func numJewelsInStones(J string, S string) int {
	if len(J) == 0 || len(S) == 0 {
		return 0
	}
	jMap := make(map[string]bool)
	for i := 0; i < len(J); i++ {
		jMap[string(J[i])] = true
	}
	num := 0
	for i := 0; i < len(S); i++ {
		if _, ok := jMap[string(S[i])]; ok {
			num++
		}
	}
	return num
}
