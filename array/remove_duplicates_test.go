package array

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var nums []int
	var h, r int

	nums = []int{1, 1, 2}
	h = 2
	r = removeDuplicates(nums)
	t.Logf("%t nums=%v h=%d r=%d", r==h, nums, h, r)

	nums = []int{0,0,1,1,1,2,2,3,3,4}
	h = 5
	r = removeDuplicates(nums)
	t.Logf("%t nums=%v h=%d r=%d", r==h, nums, h, r)
}
