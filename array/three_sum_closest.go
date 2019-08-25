//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
//
//例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
//
//与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

package array

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
