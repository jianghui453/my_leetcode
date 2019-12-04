package dynamic_programming

import (
	"testing"
)

func Test_coinChange(t *testing.T) {
	var coins []int
	var amount, hope, ret int

	coins, amount = []int{1, 2, 5}, 11
	hope, ret = 3, coinChange(coins, amount)
	t.Logf("%t coins=%v amount=%d hope=%d ret=%d", ret == hope, coins, amount, hope, ret)

	coins, amount = []int{186,419,83,408}, 6249
	hope, ret = 20, coinChange(coins, amount)
	t.Logf("%t coins=%v amount=%d hope=%d ret=%d", ret == hope, coins, amount, hope, ret)
}