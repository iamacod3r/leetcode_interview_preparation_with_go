package easy

import "interview_go/common"

// https://leetcode.com/problems/plus-one/description/
type PlusOne_66 struct {
}

func (t *PlusOne_66) plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i]%10 != 0 {
			return digits
		}
		digits[i] = 0
	}

	digits = append([]int{1}, digits...)
	return digits
}

func (t *PlusOne_66) Test() {
	// create test cases for plusOne function
	// case 1
	// digits1 := []int{1, 2, 3}
	// e1 := []int{1, 2, 4}
	// r1 := t.plusOne(digits1)
	// common.PrintIntSlice(r1, e1)

	// // case 2
	// digits2 := []int{4, 3, 2, 1}
	// e2 := []int{4, 3, 2, 2}
	// r2 := t.plusOne(digits2)
	// common.PrintIntSlice(r2, e2)
	// case 3
	digits3 := []int{1, 9}
	e3 := []int{2, 0}
	r3 := t.plusOne(digits3)
	common.PrintIntSlice(r3, e3)

	// empty case
	// empty := []int{}
	// e4 := []int{1}
	// r4 := t.plusOne(empty)
	// common.PrintIntSlice(r4, e4)
}
