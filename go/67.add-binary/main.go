package main

import "fmt"

func main() {
	//a, b := "11", "1"
	//a, b := "1010", "1011"
	a, b := "1", "111"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	i, j := m-1, n-1
	k := m
	if n > m {
		k = n
	}
	k++

	res := make([]byte, k)
	var carry, r uint8

	for i >= 0 || j >= 0 {
		k--
		if i >= 0 && j >= 0 {
			r = a[i] - '0' + b[j] - '0' + carry
			i--
			j--
		} else if i == -1 && j >= 0 {
			r = b[j] - '0' + carry
			j--
		} else if i >= 0 && j == -1 {
			r = a[i] - '0' + carry
			i--
		}
		res[k] = '0' + r%2
		carry = r / 2
	}
	if carry == 1 {
		res[0] = '1'
	} else {
		res = res[1:]
	}
	return string(res)
}
