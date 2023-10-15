package algorithm

import (
	"interview_go/common"
	"sort"
)

/*
	--Solution Steps--
	1. Visualize Example(s)
	2. Find an appropriate subproblem
	3. Find releationships among subproblems
	4. Generalize the relationship
	5. Implement by solving subproblems in order
*/

type DynamicProgramming struct {
}

/*
Longest Increasing Subsequence - LIS

For a sequence a1, a2,...an, find the length of the
longest increasing subsequence ai1, ai2,...,aik
Constraints: i1 < i2 < ... < ik; ai1 < ai2 < ... < aik

LIS([3, 1, 8, 2, 5]) -> 3 -> (1,2,5)
LIST([5, 2, 8, 6, 3, 6, 9, 5]) -> 4 -. (2, 3, 6, 9)

We will focus on the length of the Longest Increasing Subsequence - LIS
*/
func (d *DynamicProgramming) LongestIncreasingSubsequence(A []int) int {
	/* LIS = Longest Path in DAG + 1
	 	1. Visualize Example(s)
	 	2. Find an appropriate subproblem
			- All increasing subsequences are subsets of original sequence.
			- All increasing subsequences have a start and end.
			- Let's focus on the end index of an increasing subsequence
				- Subproblem: LIS[k] = LIS ending at index k
				- What subproblems are needed to solve LIST[4] ?
				- How do we use these subproblems to solve LIS[4] ?
					LIS[4] = 1 + max(LIS[0],LIS[1],LIS[3]) = 3
		3. Find releationships among subproblems
		4. Generalize the relationship
			- How do we solve subproblem LIS[n] ?
				LIS[n] = 1 + max{LIS[k] | k < n, A[k] < A[n]}
		5. Implement by solving subproblems in order
	*/

	L := make([]int, len(A))
	for i := 0; i < len(L); i++ {
		L[i] = 1
	}

	result := 0
	for i := 1; i < len(L); i++ {
		max := 0
		for k := 0; k < i; k++ {
			if A[k] < A[i] {
				if L[k] > max {
					max = L[k]
				}
			}
		}
		L[i] = 1 + max
		if L[i] > result {
			result = L[i]
		}
	}

	return result
}

func (d *DynamicProgramming) TallestBoxStacking(boxes [][]int) int {
	sort.SliceStable(boxes, func(i, j int) bool {
		return boxes[i][0] < boxes[j][0]
	})
	l := len(boxes)
	heights := []int{}
	for i := 0; i < l; i++ {
		heights = append(heights, boxes[i][2])
	}

	for i := 1; i < l; i++ {
		box := boxes[i]
		stacked := []int{}
		for j := 0; j < i; j++ {
			if d.canBeStacked(boxes[j], box) {
				stacked = append(stacked, j)
			}
		}
		// heights[i] = box[2] + d.findMaxBox(stacked, heights)
		heights[i] += d.findMaxBox(stacked, heights)
	}
	return d.findMaxValue(heights)
}

func (d *DynamicProgramming) findMaxBox(stacked []int, heights []int) int {

	result := 0
	for i := 0; i < len(stacked); i++ {
		if heights[stacked[i]] > result {
			result = heights[stacked[i]]
		}
	}
	return result
}

func (d *DynamicProgramming) findMaxValue(values []int) int {
	if len(values) == 0 {
		return 0
	}
	sort.Ints(values)
	return values[len(values)-1]
}

func (d *DynamicProgramming) canBeStacked(top, bottom []int) bool {
	// 0 -> Length, 1 -> Width, 2 -> Height
	return top[0] < bottom[0] && top[1] < bottom[1]
}

func (d *DynamicProgramming) Test() {
	// slice := []int{5, 2, 8, 6, 3, 6, 9, 5}
	// e := 4
	// r := d.LongestIncreasingSubsequence(slice)
	// common.PrintInt(r, e)

	// indexes
	// 0 -> Length
	// 1 -> Width
	// 2 -> Height
	boxes := [][]int{
		{4, 5, 3},
		{2, 3, 2},
		{3, 6, 2},
		{1, 5, 4},
		{2, 4, 1},
		{1, 2, 2},
	}
	e := 7
	r := d.TallestBoxStacking(boxes)
	common.PrintInt(r, e)
}
