package main

func main() {
	s := "()"
	println(isValid(s))
	s = "()[]{}"
	println(isValid(s))
	s = "(]"
	println(isValid(s))
	s = "([)]"
	println(isValid(s))
	s = "{[]}"
	println(isValid(s))
	s = "{[({"
	println(isValid(s))
	s = ""
	println(isValid(s))
	s = "){"
	println(isValid(s))
}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if  len(s) % 2 != 0 {
		return false
	}
	stack := make([]byte, len(s)/2)
	index := 0

	left := []byte{'(', '{', '['}
	right := []byte{')', '}', ']'}

	for _, v := range s {
		c := byte(v)
		if c == left[0] || c == left[1] || c == left[2] {
			if index >= len(s) / 2 {
				return false
			}
			stack[index] = c
			index++
		} else if index > 0 && ((c == right[0] && stack[index-1] == left[0]) || (c == right[1] && stack[index-1] == left[1]) || (c == right[2] && stack[index-1] == left[2])) {
			index--
			if index < 0 {
				return false
			}
		} else {
			return false
		}
	}
	if index != 0 {
		return false
	}
	return true
}