package easy

import "interview_go/common"

// https://leetcode.com/problems/add-two-integers/
type AddTwoIntegers struct{}

func (t *AddTwoIntegers) sum(num1 int, num2 int) int {
	return num1 + num2
}

func (t *AddTwoIntegers) Test() {
	num1, num2 := 12, 5
	e := 17
	r := t.sum(num1, num2)
	common.PrintInt(r, e)

	num3, num4 := -10, 4
	e2 := -6
	r2 := t.sum(num3, num4)
	common.PrintInt(r2, e2)

}
