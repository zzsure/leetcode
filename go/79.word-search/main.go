package main

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	println(exist(board, "ABCCED"))
	println(exist(board, "SEE"))
	println(exist(board, "ABCB"))
}

func exist(board [][]byte, word string) bool {
	if word == "" {
		return true
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] == board[i][j] {
				if search(board, word, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func search(board [][]byte, word string, i, j, idx int) bool {
	if idx == len(word) {
		return true
	}
	if !isValidIdx(board, i, j) || board[i][j] != word[idx] {
		return false
	}
	board[i][j] = '*'
	res := search(board, word, i - 1, j, idx + 1) || search(board, word, i, j + 1, idx + 1) || search(board, word, i + 1, j, idx + 1) || search(board, word, i, j - 1, idx + 1)
	board[i][j] = word[idx]
	return res
}

func isValidIdx(board [][]byte, i, j int) bool {
	return i >= 0 && i < len(board) && j >= 0 && j < len(board[0])
}