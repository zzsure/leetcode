package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 0-255
	// ["255.255.11.135", "255.255.111.35"]
	fmt.Println(restoreIpAddresses("25525511135"))
}

func restoreIpAddresses(s string) []string {
	count := len(s)
	var res, temp []string
	if count < 4 || count > 12 {
		return res
	}
	dfs(&s, &res, &temp, 0, 0)
	return res
}

func dfs(s *string, res *[]string, temp *[]string, depth, i int) {
	if depth == 4 {
		r := strings.Join(*temp, ".")
		if len(r) == len(*s)+3 {
			*res = append(*res, r)
		}
		return
	}
	count := len(*s)
	for j := i; j < i+3 && j < count; j++ {
		t := (*s)[i:(j + 1)]
		if t[0] == '0' && j > i {
			continue
		}
		tmp, err := strconv.Atoi(t)
		if err != nil {
			continue
		}
		if tmp > 255 {
			continue
		}
		*temp = append(*temp, t)
		dfs(s, res, temp, depth+1, j+1)
		*temp = (*temp)[0 : len(*temp)-1]
	}
}
