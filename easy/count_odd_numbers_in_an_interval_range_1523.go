package easy

import "interview_go/common"

type CountOddNumbersInAnIntervalRange struct {
}

func (x *CountOddNumbersInAnIntervalRange) countOdds(low int, high int) int {
	if low&1 == 0 {
		low++
	}
	if low > high {
		return 0
	}
	return (high-low)/2 + 1
}

func (x *CountOddNumbersInAnIntervalRange) Test() {

	e1 := 3
	r1 := x.countOdds(3, 7)
	common.PrintInt(r1, e1)

	e2 := 1
	r2 := x.countOdds(8, 10)
	common.PrintInt(r2, e2)

}
