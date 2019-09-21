package array

import "testing"

func TestLongestConsecutive(t *testing.T) {
    var nums []int
    var h, r int

    nums = []int{100, 4, 200, 1, 3, 2}
    h = 4
    r = longestConsecutive(nums)
    t.Logf("nums=%v h=%d r=%d", nums, h, r)

    nums = []int{100}
    h = 1
    r = longestConsecutive(nums)
    t.Logf("nums=%v h=%d r=%d", nums, h, r)

    nums = []int{100, 2}
    h = 1
    r = longestConsecutive(nums)
    t.Logf("nums=%v h=%d r=%d", nums, h, r)

    nums = []int{1, 2, 3, 4, 5}
    h = 5
    r = longestConsecutive(nums)
    t.Logf("nums=%v h=%d r=%d", nums, h, r)
}
