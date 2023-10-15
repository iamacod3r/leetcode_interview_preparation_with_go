package medium

import (
	"fmt"
	"strconv"
	"time"
)

type ValidSudoku_36 struct{}

// https://leetcode.com/problems/valid-sudoku/description/
func (v *ValidSudoku_36) isValidSudoku(board [][]byte) bool {
	size := len(board)
	if size == 0 {
		return false
	}
	empty := byte('.')
	isValid := func(row, col int, num byte) bool {
		r := 3 * (row / 3)
		c := 3 * (col / 3)

		for i := 0; i < size; i++ {

			if board[row][i] == num && i != col {
				return false
			}

			if board[i][col] == num && i != row {
				return false
			}
			// (r/3) * 3 + (c/3)
			br := r + i/3 // box row info
			bc := c + i%3 // box col info
			if board[br][bc] == num && br != row && bc != col {
				return false
			}
		}
		return true
	}

	for r := 0; r < size; r++ {

		for c := 0; c < size; c++ {
			if board[r][c] != empty {
				if !isValid(r, c, board[r][c]) {
					return false
				}
			}
		}
	}
	return true
}

func (v *ValidSudoku_36) isValidSudokuHashSet(board [][]byte) bool {

	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}

	rows := make(map[int]int)
	cols := make(map[int]int)

	squares := make([][]int, 3)
	squares[0] = []int{1, 1, 1}
	squares[1] = []int{1, 1, 1}
	squares[2] = []int{1, 1, 1}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {

			if board[i][j] != '.' {

				boardPosition, _ := strconv.Atoi(string(board[i][j]))

				if rows[boardPosition] == 0 {
					rows[boardPosition] = 1
				}
				if cols[boardPosition] == 0 {
					cols[boardPosition] = 1
				}

				if rows[boardPosition]%primes[i] == 0 ||
					cols[boardPosition]%primes[j] == 0 ||
					squares[i/3][j/3]%primes[boardPosition-1] == 0 {
					return false
				}

				cols[boardPosition] *= primes[j]
				rows[boardPosition] *= primes[i]
				squares[i/3][j/3] *= primes[boardPosition-1]
			}
		}
	}
	return true
}

func (v *ValidSudoku_36) Test() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	// expected := true
	start := time.Now()
	// result := v.isValidSudoku(board)
	v.isValidSudoku(board)
	fmt.Println("array -> ", time.Since(start).Nanoseconds())
	// common.PrintBool(result, expected)

	start = time.Now()
	// result = v.isValidSudokuHashSet(board)
	v.isValidSudokuHashSet(board)
	fmt.Println("hashset -> ", time.Since(start).Nanoseconds())
	// common.PrintBool(result, expected)

	board = [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	// expected = false
	start = time.Now()
	// result = v.isValidSudoku(board)
	v.isValidSudoku(board)
	fmt.Println("array -> ", time.Since(start).Nanoseconds())
	// common.PrintBool(result, expected)

	start = time.Now()
	// result = v.isValidSudokuHashSet(board)
	v.isValidSudokuHashSet(board)
	fmt.Println("hashset -> ", time.Since(start).Nanoseconds())
	// common.PrintBool(result, expected)

}
