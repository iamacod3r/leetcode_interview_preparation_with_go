package easy

import "interview_go/common"

// https://leetcode.com/problems/toeplitz-matrix/description/
type ToeplitzMatrix766 struct {
}

func (t *ToeplitzMatrix766) isToeplitzMatrix(matrix [][]int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[0]); c++ {
			if matrix[r-1][c-1] != matrix[r][c] {
				return false
			}
		}
	}
	return true
}

func (t *ToeplitzMatrix766) Test() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}
	e := true
	r := t.isToeplitzMatrix(matrix)
	common.PrintBool(r, e)

	matrix2 := [][]int{
		{1, 2},
		{2, 2},
	}
	e2 := false
	r2 := t.isToeplitzMatrix(matrix2)
	common.PrintBool(r2, e2)
}
