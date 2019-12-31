// 给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

// 示例 1 :

// 输入: 2736
// 输出: 7236
// 解释: 交换数字2和数字7。
// 示例 2 :

// 输入: 9973
// 输出: 9973
// 解释: 不需要交换。
// 注意:

// 给定数字的范围是 [0, 108]

package array

import (
// "fmt"
)

func maximumSwap(num int) int {
	var (
		nums   []int = make([]int, 0)
		l, ret int
	)

	for ; num > 0; num /= 10 {
		nums = append(nums, num%10)
		l++
	}

	for i := l - 1; i >= 0; i-- {
		if nums[i] == 9 {
			ret = ret*10 + nums[i]
			continue
		}

		idx := i
		for j := i - 1; j >= 0; j-- {
			if nums[j] >= nums[idx] {
				idx = j
			}
		}

		if nums[idx] > nums[i] {
			nums[i], nums[idx] = nums[idx], nums[i]
			for ; i >= 0; i-- {
				ret = ret*10 + nums[i]
			}
		} else {
			ret = ret*10 + nums[i]
		}
	}

	return ret
}
