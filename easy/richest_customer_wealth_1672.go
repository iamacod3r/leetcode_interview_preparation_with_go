package easy

import "interview_go/common"

// https://leetcode.com/problems/richest-customer-wealth/
type RichestCustomerWealth struct {
}

func (w *RichestCustomerWealth) maximumWealth(accounts [][]int) int {

	result := 0

	for _, c := range accounts {
		customerSum := 0
		for _, a := range c {
			customerSum += a
		}

		if customerSum > result {
			result = customerSum
		}
	}
	return result
}

func (w *RichestCustomerWealth) Test() {
	accounts := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	e := 6
	r := w.maximumWealth(accounts)
	common.PrintInt(r, e)

	acc2 := [][]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}
	e2 := 10
	r2 := w.maximumWealth(acc2)
	common.PrintInt(r2, e2)

	acc3 := [][]int{
		{2, 8, 7},
		{7, 1, 3},
		{1, 9, 5},
	}
	e3 := 17
	r3 := w.maximumWealth(acc3)

	common.PrintInt(r3, e3)

}
