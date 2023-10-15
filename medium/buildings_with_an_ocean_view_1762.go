package medium

import "interview_go/common"

// https://leetcode.com/problems/buildings-with-an-ocean-view/
type BuildingsWithAnOcenView struct {
}

// Time Complexity : O(n)
// Space Complexity : O(1)
func (v *BuildingsWithAnOcenView) findBuildings(heights []int) []int {

	l := len(heights)
	result := []int{l - 1}
	prev := heights[l-1]
	for x := l - 2; x >= 0; x-- {
		curr := heights[x]
		if curr > prev {
			result = append(result, x)
			prev = curr
		}
	}
	left, right := 0, len(result)-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}

	return result
}

func (v *BuildingsWithAnOcenView) Test() {
	h1 := []int{4, 2, 3, 1}
	e1 := []int{0, 2, 3}
	r1 := v.findBuildings(h1)
	common.PrintIntSlice(r1, e1)

	h2 := []int{4, 3, 2, 1}
	e2 := []int{0, 1, 2, 3}
	r2 := v.findBuildings(h2)
	common.PrintIntSlice(r2, e2)

	h3 := []int{1, 2, 3, 4}
	e3 := []int{3}
	r3 := v.findBuildings(h3)
	common.PrintIntSlice(r3, e3)

	h4 := []int{5, 3, 2, 4, 1, 1}
	e4 := []int{0, 3, 5}
	r4 := v.findBuildings(h4)
	common.PrintIntSlice(r4, e4)

}
