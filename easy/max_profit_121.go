package easy

import (
	"interview_go/common"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func MaxProfit121(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}

	buy := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {

		if buy < prices[i] {
			profit = max(profit, prices[i]-buy)
		}

		buy = min(buy, prices[i])
	}

	return profit
}

func max(a, b int) int {

	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {

	if a > b {
		return b
	}

	return a
}

func MaxProfit121_Test() {
	common.PrintlnStr("MaxProfit_121_Test")
	p1 := [6]int{7, 1, 5, 3, 6, 4}
	e1 := 5
	r1 := MaxProfit121(p1[:])
	common.PrintInt(r1, e1)

	p2 := [5]int{7, 6, 4, 3, 1}
	e2 := 0
	r2 := MaxProfit121(p2[:])
	common.PrintInt(r2, e2)
}
