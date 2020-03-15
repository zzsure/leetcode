package main

func main() {
	res := generateParenthesis(3)
	for _, s := range res {
		println(s)
	}
}

func generateParenthesis(n int) []string {
	var res []string
	if n == 0 {
		return []string{""}
	}
	for i := 0; i < n; i++ {
		leftRes := generateParenthesis(i)
		rightRes := generateParenthesis(n - 1 - i)
		for _, left := range leftRes {
			for _, right := range rightRes {
				res = append(res, "("+left+")"+right)
			}
		}
	}

	return res
}
