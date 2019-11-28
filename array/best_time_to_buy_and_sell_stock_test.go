package array

import "testing"

func TestMaxProfit(t *testing.T) {
	var prices []int
	var h, r int

	prices = []int{7, 1, 5, 3, 6, 4}
	h = 5
	r = maxProfit(prices)
	t.Logf("%t prices=%v h=%d r=%d", h == r, prices, h, r)

	prices = []int{7, 6, 4, 3, 1}
	h = 0
	r = maxProfit(prices)
	t.Logf("%t prices=%v h=%d r=%d", h == r, prices, h, r)
}
