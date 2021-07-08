// 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
// 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
// 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

package two_pointers

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	numsLen := len(nums)

	sort.Ints(nums)

	ret := nums[0] + nums[1] + nums[2]
	for i := 0; i < numsLen; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]-int(math.Abs(float64(ret))) > target {
			break
		}

		j, k := i+1, numsLen-1
		for j < k {
			v := nums[i] + nums[j] + nums[k]
			if math.Abs(float64(v-target)) < math.Abs(float64(ret-target)) {
				ret = v
			}

			if v > target {
				k--
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			} else if v < target {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else {
				return target
			}
		}
	}
	return ret
}

// func threeSumClosest(nums []int, target int) int {
// 	var ret, i, left, right int
// 	ret = nums[0] + nums[1] + nums[2]
// 	sort.Ints(nums)
// Loop:
// 	for i = 1; i+1 < len(nums); i++ {
// 		left = i - 1
// 		right = i + 1
// 		for left >= 0 && right < len(nums) {
// 			if math.Abs(float64(nums[left]+nums[i]+nums[right]-target)) <= math.Abs(float64(ret-target)) {
// 				ret = nums[left] + nums[i] + nums[right]
// 			}
// 			if nums[left]+nums[i]+nums[right] > target {
// 				left--
// 				continue
// 			}
// 			if nums[left]+nums[i]+nums[right] < target {
// 				right++
// 				continue
// 			}
// 			break Loop
// 		}
// 	}
// 	return ret
// }
