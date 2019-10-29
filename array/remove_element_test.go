package array

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	var nums []int
	var val, h, r int

	nums = []int{3,2,2,3}
	val = 3
	h = 2
	r = removeElement(nums, val)
	t.Logf("%t nums=%v val=%d h=%d r=%d", r==h, nums, val, h, r)

	nums = []int{0,1,2,2,3,0,4,2}
	val = 2
	h = 5
	r = removeElement(nums, val)
	t.Logf("%t nums=%v val=%d h=%d r=%d", r==h, nums, val, h, r)

	nums = []int{3,2,2,3}
	val = 1
	h = 4
	r = removeElement(nums, val)
	t.Logf("%t nums=%v val=%d h=%d r=%d", r==h, nums, val, h, r)

	nums = []int{1}
	val = 1
	h = 0
	r = removeElement(nums, val)
	t.Logf("%t nums=%v val=%d h=%d r=%d", r==h, nums, val, h, r)
}