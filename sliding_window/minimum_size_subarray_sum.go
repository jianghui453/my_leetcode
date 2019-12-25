// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。

// 示例:

// 输入: s = 7, nums = [2,3,1,2,4,3]
// 输出: 2
// 解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
// 进阶:

// 如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。

package sliding_window

func minSubArrayLen(s int, nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	var ret int
	sums := make([]int, l)

	for i := 0; i < l; i++ {
		if i > 0 {
			sums[i] = sums[i-1] + nums[i]
		} else {
			sums[i] = nums[i]
		}

		if sums[i] < s {
			continue
		}

		min, max := 0, i
		for sums[i] >= s && min <= max {
			mid := (min + max) / 2

			if mid > 0 && sums[i]-sums[mid-1] < s {
				max = mid - 1
			} else {
				if ret == 0 || ret > (i+1-mid) {
					ret = i + 1 - mid
				}

				if sums[i] > s {
					min = mid + 1
				} else {
					break
				}
			}
		}
	}

	return ret
}

// func minSubArrayLen(s int, nums []int) int {
// 	l := len(nums)
// 	if l == 0 {
// 		return 0
// 	}

// 	var ret, sum, left, right int
// 	sum = nums[0]
// 	for left <= right && right < l {
// 		if sum < s {
// 			right++
// 			if right < l {
// 				sum += nums[right]
// 			}
// 		} else {
// 			if ret == 0 || right-left+1 < ret {
// 				ret = right+1-left
// 			}
// 			sum -= nums[left]
// 			left++
// 		}
// 	}

// 	return ret
// }
