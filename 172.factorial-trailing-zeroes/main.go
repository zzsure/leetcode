package main

func main() {
	println(trailingZeroes(3))
	println(trailingZeroes(5))
	println(trailingZeroes(30))
	println(trailingZeroes(150))
}

func trailingZeroes(n int) int {
	count := 0
	if n < 5 {
		return 0
	}
	count = n / 5 + trailingZeroes(n / 5)
	return count
}