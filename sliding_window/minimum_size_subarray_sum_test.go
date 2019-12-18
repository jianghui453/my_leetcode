package sliding_window

import (
	"testing"
)

func Test_minSubArrayLen(t *testing.T) {
	var nums []int
	var s, hope, ret int

	nums = []int{2,3,1,2,4,3}
	s = 7
	hope = 2
	ret = minSubArrayLen(s, nums)
	t.Log(ret == hope, nums, s, hope, ret)

	nums = []int{2,3}
	s = 7
	hope = 0
	ret = minSubArrayLen(s, nums)
	t.Log(ret == hope, nums, s, hope, ret)

	nums = []int{2,3,1,2,4,3, 7}
	s = 7
	hope = 1
	ret = minSubArrayLen(s, nums)
	t.Log(ret == hope, nums, s, hope, ret)
}