package dynamic_programming

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var prices []int
	var hope, ret int

	prices = []int{3, 3, 5, 0, 0, 3, 1, 4}
	hope = 6
	ret = maxProfit(prices)
	t.Logf("%t prices=%v hope=%d ret=%d", ret == hope, prices, hope, ret)

	prices = []int{1, 2, 3, 4, 5}
	hope = 4
	ret = maxProfit(prices)
	t.Logf("%t prices=%v hope=%d ret=%d", ret == hope, prices, hope, ret)

	prices = []int{7, 6, 4, 3, 1}
	hope = 0
	ret = maxProfit(prices)
	t.Logf("%t prices=%v hope=%d ret=%d", ret == hope, prices, hope, ret)

	prices = []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	hope = 13
	ret = maxProfit(prices)
	t.Logf("%t prices=%v hope=%d ret=%d", ret == hope, prices, hope, ret)
}
