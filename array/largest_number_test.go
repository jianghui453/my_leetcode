package array

import (
	"testing"
)

func TestLargestNumber(t *testing.T) {
	// t.Log(bigger("3", "33"))
	// return
	var (
		nums []int
		hope, ret string
	)

	nums = []int{3,43,48,94,85,33,64,32,63,66}
	hope = "9485666463484333332"
	ret = largestNumber(nums)
	t.Log(ret == hope, nums, hope, ret)
	
	nums = []int{20, 1}
	hope = "201"
	ret = largestNumber(nums)
	t.Log(ret == hope, nums, hope, ret)
// return
	nums = []int{10, 2}
	hope = "210"
	ret = largestNumber(nums)
	t.Log(ret == hope, nums, hope, ret)

	nums = []int{3,30,34,5,9}
	hope = "9534330"
	ret = largestNumber(nums)
	t.Log(ret == hope, nums, hope, ret)

	nums = []int{10, 102, 100}
	hope = "10210100"
	ret = largestNumber(nums)
	t.Log(ret == hope, nums, hope, ret)

	nums = []int{0, 0}
	hope = "0"
	ret = largestNumber(nums)
	t.Log(ret == hope, nums, hope, ret)
}