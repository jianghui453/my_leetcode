package array

import (
	"testing"
)

func TestMaxCoins(t *testing.T) {
	var (
		nums []int
		hope, ret int
	)

	nums = []int{3,1,5,8}
	hope = 167
	ret = maxCoins(nums)
	t.Log(ret == hope, nums, hope, ret)

	// nums = []int{3}
	// hope = 3
	// ret = maxCoins(nums)
	// t.Log(ret == hope, nums, hope, ret)

	// nums = []int{3,1}
	// hope = 6
	// ret = maxCoins(nums)
	// t.Log(ret == hope, nums, hope, ret)

	// nums = []int{9,76,64,21,97,60}
	// hope = 1086136
	// ret = maxCoins(nums)
	// t.Log(ret == hope, nums, hope, ret)
}
