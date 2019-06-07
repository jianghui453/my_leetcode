package leet_code

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var ret, i, left, right int
	ret = nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
Loop:
	for i = 1; i+1 < len(nums); i++ {
		left = i - 1
		right = i + 1
		for left >= 0 && right < len(nums) {
			if math.Abs(float64(nums[left]+nums[i]+nums[right]-target)) <= math.Abs(float64(ret-target)) {
				ret = nums[left] + nums[i] + nums[right]
			}
			if nums[left]+nums[i]+nums[right] > target {
				left--
				continue
			}
			if nums[left]+nums[i]+nums[right] < target {
				right++
				continue
			}
			break Loop
		}
	}
	return ret
}
