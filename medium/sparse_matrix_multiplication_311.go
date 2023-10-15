package medium

import "interview_go/common"

// https://leetcode.com/problems/sparse-matrix-multiplication/description/
type SparseMatrixMultiplication struct {
}

func (s *SparseMatrixMultiplication) multiply_naive(mat1 [][]int, mat2 [][]int) [][]int {

	n := len(mat1)
	k := len(mat1[0])
	m := len(mat2[0])

	ans := make([][]int, n)

	for i := range ans {
		ans[i] = make([]int, m)
	}

	for row := 0; row < n; row++ {
		for i := 0; i < k; i++ {
			if mat1[row][i] != 0 {
				for col := 0; col < m; col++ {
					ans[row][col] += mat1[row][i] * mat2[i][col]
				}
			}
		}
	}
	return ans
}

func (m *SparseMatrixMultiplication) Test() {

	mat1 := [][]int{
		{1, 0, 0},
		{-1, 0, 3},
	}

	mat2 := [][]int{
		{7, 0, 0},
		{0, 0, 0},
		{0, 0, 1},
	}

	e := [][]int{
		{7, 0, 0},
		{-7, 0, 3},
	}

	r := m.multiply_naive(mat1, mat2)

	common.Print2DSlice(r, e)
}
