// 给定一个无序的整数数组，找到其中最长上升子序列的长度。

// 示例:

// 输入: [10,9,2,5,3,7,101,18]
// 输出: 4 
// 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
// 说明:

// 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
// 你算法的时间复杂度应该为 O(n2) 。
// 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

package dynamic_programming

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	tails := make([]int, l)
	ret := 0
	for i := 0; i < l; i++ {
		min, max := 0, ret-1

		for min <= max {
			mid := (min+max)/2
			if tails[mid] > nums[i] {
				max = mid-1
			} else if tails[mid] < nums[i] {
				min = mid+1
			} else {
				tails[mid] = nums[i]
				break
			}
		}

		if min > max {
			tails[min] = nums[i]
			if min == ret {
				ret++
			}
		}

		fmt.Println(tails, ret)
	}

	return ret
}

// func lengthOfLIS(nums []int) int {
// 	l := len(nums)
// 	if l == 0 {
// 		return 0
// 	}

// 	s := make([]int, l)
// 	dp := make([]int, l)
// 	dp[0] = 1
// 	ret := 1
// 	for i := 1; i < l; i++ {
// 		dp[i] = 1

// 		for j := i-1; j >= 0; j-- {
// 			if nums[i] > nums[s[j]] {
// 				dp[i] = max(dp[i], dp[s[j]]+1)
// 				break
// 			}
// 		}

// 		s[i] = i
// 		for j := i-1; j >= 0; j-- {
// 			if dp[s[j]] > dp[s[j+1]] {
// 				s[j], s[j+1] = s[j+1], s[j]
// 			} else {
// 				break
// 			}
// 		}

// 		ret = max(ret, dp[i])
// 	}
	
// 	return ret
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }