package main

import (
	"fmt"
	"strings"
)

func main() {
	//path := "/home/"
	//path := "/../"
	//path := "/home//foo/"
	//path := "/a/./b/../../c/"
	//path := "/a/../../b/../c//.//"
	path := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(path))
}

func simplifyPath(path string) string {
	strs := strings.Split(path, "/")
	stack := new(Stack)
	for _, str := range strs {
		if str == "" || str == "." {
			continue
		}
		if str == ".." {
			stack.Pop()
			continue
		}
		stack.Push(str)
	}
	res := "/"
	for stack.Length() > 0 {
		if res != "/" {
			res = "/" + stack.Pop() + res
		} else {
			res = "/" + stack.Pop()
		}
	}
	return res
}

type Stack struct {
	length int
	strs   []string
}

func (s *Stack) Push(str string) {
	if s.length >= len(s.strs) {
		s.strs = append(s.strs, str)
	}
	s.strs[s.length] = str
	s.length++
}

func (s *Stack) Pop() string {
	if s.length == 0 {
		return ""
	}
	s.length--
	if s.length < len(s.strs) {
		return s.strs[s.length]
	}
	return ""
}

func (s *Stack) Length() int {
	return s.length
}
