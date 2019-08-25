package array

import "testing"

func TestRemoveDuplicates2(t *testing.T) {
	var nums []int
	var hope, ret int

	nums = []int{1,1,1,2,2,3}
	hope = 5
	ret = remove_duplicates2(nums)
	t.Logf("nums=%v hope=%d ret=%d", nums, hope, ret)

	nums = []int{0,0,1,1,1,1,2,3,3}
	hope = 7
	ret = remove_duplicates2(nums)
	t.Logf("nums=%v hope=%d ret=%d", nums, hope, ret)
}