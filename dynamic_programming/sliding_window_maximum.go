// 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

// 返回滑动窗口中的最大值。

// 示例:

// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
// 输出: [3,3,5,5,6,7]
// 解释:

//   滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7
//

// 提示：

// 你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

// 进阶：

// 你能在线性时间复杂度内解决此题吗？

package dynamic_programming

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	if k <= 1 {
		return nums
	}

	l := len(nums)
	if l == 0 {
		return []int{}
	}

	var (
		left, right []int = make([]int, l), make([]int, l)
	)

	left[0], right[l-1] = nums[0], nums[l-1]
	for i := 1; i < l; i++ {
		if i%k == 0 {
			left[i] = nums[i]
		} else {
			left[i] = max(left[i-1], nums[i])
		}

		if (l-1-i)%k == k-1 {
			right[l-1-i] = nums[l-1-i]
		} else {
			right[l-1-i] = max(right[l-i], nums[l-1-i])
		}
	}

	var ret []int
	if k < l {
		ret = make([]int, l-k+1)
		for i := 0; i < l-k+1; i++ {
			ret[i] = max(left[i+k-1], right[i])
		}
	} else {
		ret = append(ret, max(left[l-1], right[0]))
	}

	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
