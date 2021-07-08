// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

// Example:

// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Follow up:

// If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

package dynamic_programming

import (
	"math"
)

func maxSubArray(nums []int) int {
	max := math.MinInt32
	l := len(nums)
	if l == 0 {
		return max
	}

	dp := make([]int, l)
	dp[0] = nums[0]
	max = int(math.Max(float64(max), float64(nums[0])))

	for i := 1; i < l; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		max = int(math.Max(float64(max), float64(dp[i])))
	}

	return max
}

// func maxSubArray(nums []int) int {
// 	nLen := len(nums)
// 	if nLen == 0 {
// 		return 0
// 	}
// 	max := nums[0]
// 	sum := nums[0]
// 	if nums[0] < 0 {
// 		sum = 0
// 	}
// 	for i := 1; i < nLen; i++ {
// 		sum += nums[i]
// 		if sum <= 0 {
// 			if nums[i] > max {
// 				max = nums[i]
// 			}
// 			sum = 0
// 			continue
// 		}
// 		if sum > max {
// 			max = sum
// 		}
// 	}
// 	return max
// }

// func maxSubArray(nums []int) int {
// 	nLen := len(nums)
// 	if nLen == 0 {
// 		return 0
// 	}
// 	max := nums[0]
// 	for i := 0; i < nLen-1; i++ {
// 		sum := nums[i]
// 		if sum > max {
// 			max = sum
// 		}
// 		_maxSubArray(nums[i+1:], nLen-1-i, sum, &max)
// 	}
// 	if nums[nLen-1] > max {
// 		max = nums[nLen-1]
// 	}
// 	return max
// }

// func _maxSubArray(nums []int, nLen, sum int, max *int) {
// 	fmt.Printf("nums=%v nLen=%d sum=%d max=%d\n", nums, nLen, sum, *max)
// 	if nLen == 0 {
// 		return
// 	}
// 	sum += nums[0]
// 	if sum > *max {
// 		*max = sum
// 	}
// 	if nLen > 1 {
// 		_maxSubArray(nums[1:], nLen-1, sum, max)
// 	}
// }
