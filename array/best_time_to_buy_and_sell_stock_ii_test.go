package array

import "testing"

func TestMaxProfit2(t *testing.T) {
	var prices []int
	var h, r int

	prices = []int{7, 1, 5, 3, 6, 4}
	h = 7
	r = maxProfit2(prices)
	t.Logf("%t prices=%v h=%d r=%d", h == r, prices, h, r)

	prices = []int{1, 2, 3, 4, 5}
	h = 4
	r = maxProfit2(prices)
	t.Logf("%t prices=%v h=%d r=%d", h == r, prices, h, r)

	prices = []int{7, 6, 4, 3, 1}
	h = 0
	r = maxProfit2(prices)
	t.Logf("%t prices=%v h=%d r=%d", h == r, prices, h, r)
}
