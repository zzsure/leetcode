package main

import "fmt"

func main() {
	//boardStrArr := [][]string{
	//	{"5", "3", ".", ".", "7", ".", ".", ".", "."},
	//	{"6", ".", ".", "1", "9", "5", ".", ".", "."},
	//	{".", "9", "8", ".", ".", ".", ".", "6", "."},
	//	{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
	//	{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
	//	{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
	//	{".", "6", ".", ".", ".", ".", "2", "8", "."},
	//	{".", ".", ".", "4", "1", "9", ".", ".", "5"},
	//	{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	//}
	boardStrArr := [][]string{
		{"8", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	boardByteArr := make([][]byte, len(boardStrArr))
	for i, boardStr := range boardStrArr {
		boardByte := make([]byte, len(boardStr))
		for j, board := range boardStr {
			boardByte[j] = board[0]
		}
		boardByteArr[i] = boardByte
	}
	fmt.Println(isValidSudoku(boardByteArr))

}

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 {
		return false
	}
	if len(board[0]) != 9 {
		return false
	}
	const column = 9
	const row = 9
	for _, b := range board {
		if isValid(b) == false {
			return false
		}
	}
	for j := 0; j < column; j++ {
		b := make([]byte, row)
		for i := 0; i < row; i++ {
			b[i] = board[i][j]
		}
		if isValid(b) == false {
			return false
		}
	}
	for k := 0; k < row; k++ {
		b := make([]byte, row)
		for i := 0; i < row/3; i++ {
			for j := 0; j < column/3; j++ {
				b[3*i+j] = board[3*(k/3)+i][3*(k%3)+j]
			}
		}
		if isValid(b) == false {
			return false
		}
	}
	return true
}

func isValid(board []byte) bool {
	if len(board) != 9 {
		return false
	}
	boardMap := make(map[byte]bool)
	for _, b := range board {
		if b == '.' {
			continue
		}
		if _, ok := boardMap[b]; ok {
			return false
		}
		boardMap[b] = true
	}
	return true
}
