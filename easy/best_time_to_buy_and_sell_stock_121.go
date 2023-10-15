package easy

import (
	"interview_go/common"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
type BestTimeToBuyAndSellStock struct{}

func (b *BestTimeToBuyAndSellStock) maxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}

	buy := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if buy < prices[i] {
			profit = common.Max(profit, prices[i]-buy)
		}

		buy = common.Min(buy, prices[i])
	}

	return profit
}

func (b *BestTimeToBuyAndSellStock) Test() {

	prices := []int{7, 1, 5, 3, 6, 4}
	e := 5
	r := b.maxProfit(prices)
	common.PrintInt(r, e)
}
