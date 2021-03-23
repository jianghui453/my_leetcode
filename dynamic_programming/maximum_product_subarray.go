// 给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。

// 示例 1:

// 输入: [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。
// 示例 2:

// 输入: [-2,0,-1]
// 输出: 0
// 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

package dynamic_programming

import (
// "fmt"
)

func maxProduct(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	dp := make([][2]int, l)
	if nums[0] > 0 {
		dp[0][1] = nums[0]
	} else {
		dp[0][0] = nums[0]
	}
	ret := nums[0]

	for i := 1; i < l; i++ {
		if nums[i] > 0 {
			dp[i][1] = nums[i]

			if dp[i-1][1] > 0 {
				dp[i][1] *= dp[i-1][1]
			}
			dp[i][0] = dp[i-1][0] * nums[i]
		} else {
			if dp[i-1][0] < 0 {
				dp[i][1] = dp[i-1][0] * nums[i]
			}

			dp[i][0] = nums[i]
			if dp[i-1][1] > 0 {
				dp[i][0] *= dp[i-1][1]
			}
		}

		ret = max(ret, dp[i][1])
	}

	return ret
}
