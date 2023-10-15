package easy

import (
	"fmt"
	"math"
)

type TicTacToe struct{}

// https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/

// let n be the length of the moves
// let m be the length of input moves
// Time Complexity : O(m)
/* ->
For every move, we update the value of for a row, column, diagonal, and anti-diagonal.
Each update takes constant time. We also check if any of these lines satisfies the
winning condition which also takes constant time
*/
// Space Complexity : O(n)
/* ->
We use two arrays of size n to record the value for each row and column, and two
integers of constant space to record to value for diagonal and anti-diagonal.
*/
func (*TicTacToe) tictactoe(moves [][]int) string {

	size := 3
	rows := make([]int, size)
	cols := make([]int, size)
	diag, anti_diag := 0, 0

	abs := func(x int) int {
		return int(math.Abs(float64(x)))
	}

	player := 1
	for i, v := range moves {
		r, c := v[0], v[1]
		rows[r] += player
		cols[c] += player

		if r == c {
			diag += player
		}
		if r+c == size-1 {
			anti_diag += player
		}

		if i > 3 {
			// for this comparation we need to at least 5 moves.
			// i start from 0
			if abs(rows[r]) == size || abs(cols[c]) == size ||
				abs(diag) == size || abs(anti_diag) == size {
				msg := "A"
				if player != 1 {
					msg = "B"
				}
				return msg
			}
		}

		player *= -1
	}

	if len(moves) == size*size {
		return "Draw"
	}
	return "Pending"
}

func (t *TicTacToe) Test() {
	moves := [][]int{
		{0, 0},
		{2, 0},
		{1, 1},
		{2, 1},
		{2, 2},
	}
	fmt.Println(t.tictactoe(moves))
}
