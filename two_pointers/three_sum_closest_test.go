package two_pointers

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	var nums []int
	var target, hope, ret int

	nums = []int{-1, 2, 1, -4}
	target = 1
	hope = 2
	ret = threeSumClosest(nums, target)
	if ret == hope {
		t.Logf("success: nums = %v; target = %d; hope = %d; ret = %d", nums, target, hope, ret)
	} else {
		t.Errorf("fail: nums = %v; target = %d; hope = %d; ret = %d", nums, target, hope, ret)
	}
}
