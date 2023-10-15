package medium

import (
	"fmt"
	"interview_go/common"
)

type WallsAndGates_286 struct {
}

const (
	empty = 2147483647
	gate  = 0
	wall  = -1
)

// https://leetcode.com/problems/walls-and-gates/description/
func (w *WallsAndGates_286) wallsAndGates(rooms [][]int) {

	rowSize := len(rooms)
	if rowSize == 0 {
		return
	}
	colSize := len(rooms[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := [][]int{}

	for r := 0; r < rowSize; r++ {
		for c := 0; c < colSize; c++ {
			if rooms[r][c] == gate {
				queue = append(queue, []int{r, c})
			}
		}
	}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		row := pos[0]
		col := pos[1]

		for _, dir := range directions {
			r := row + dir[0]
			c := col + dir[1]

			if r < 0 || r >= rowSize || c < 0 || c >= colSize || rooms[r][c] != empty {
				continue
			}

			rooms[r][c] = rooms[row][col] + 1
			queue = append(queue, []int{r, c})
		}
	}
}

func (w *WallsAndGates_286) Test() {

	rooms := [][]int{
		{empty, wall, gate, empty},
		{empty, empty, empty, wall},
		{empty, wall, empty, wall},
		{gate, wall, empty, empty},
	}
	fmt.Println(rooms)
	w.wallsAndGates(rooms)
	fmt.Println(rooms)

	expected := [][]int{
		{3, -1, 0, 1},
		{2, 2, 1, -1},
		{1, -1, 2, -1},
		{0, -1, 3, 4},
	}

	common.Print2DSlice(rooms, expected)
}
