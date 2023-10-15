package easy

import (
	"interview_go/common"
	"math"
)

// https://leetcode.com/problems/degree-of-an-array/
type DegreeOfAnArray697 struct {
}

type PosInfo struct {
	start, end, count int
}

func (w *DegreeOfAnArray697) findShortestSubArray_withStruct(nums []int) int {
	infoMap := map[int]PosInfo{}
	degree := 0
	result := len(nums)

	for i, v := range nums {
		info, exists := infoMap[v]
		if exists {
			info.end = i
			info.count++
		} else {

			info = PosInfo{start: i, end: i, count: 1}
		}
		infoMap[v] = info

		if info.count > degree {
			degree = info.count
			result = info.end - info.start + 1
		} else if info.count == degree {
			l := info.end - info.start + 1
			if result > l {
				result = l
			}
		}
	}

	return result
}

// Time Complexity : O(n)
// Space Complexity : O(n)
func (w *DegreeOfAnArray697) findShortestSubArray_withMultiMap(nums []int) int {

	left, right, count := map[int]int{}, map[int]int{}, map[int]int{}

	degree := 0
	for i := 0; i < len(nums); i++ {
		curr := nums[i]

		if _, ok := left[curr]; !ok {
			left[curr] = i
		}
		right[curr] = i
		count[curr]++
		if count[curr] > degree {
			degree = count[curr]
		}
	}

	result := len(nums)

	for k, v := range count {

		if v == degree {
			tmp := float64(right[k] - left[k] + 1)
			result = int(math.Min(float64(result), tmp))
		}
	}

	return result
}

func (w *DegreeOfAnArray697) Test() {
	nums := []int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2}
	e := 7
	r := w.findShortestSubArray_withStruct(nums)
	common.PrintInt(r, e)
}
