// Given an array of non-negative integers, you are initially positioned at the first index of the array.

// Each element in the array represents your maximum jump length at that position.

// Your goal is to reach the last index in the minimum number of jumps.

// Example:

// Input: [2,3,1,1,4]
// Output: 2
// Explanation: The minimum number of jumps to reach the last index is 2.
//     Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Note:

// You can assume that you can always reach the last index.

package greedy

import (
	// "math"
)

func jump(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return 0
	}

	cnt := 1
	idx := 0
	for idx < l && nums[idx]+idx < l-1 {
		maxV := 0
		maxI := 0
		for i := idx+1; i <= idx+nums[idx]; i++ {
			if i+nums[i] > maxV {
				maxV = i+nums[i]
				maxI = i
			}
		}

		idx = maxI
		cnt++
	}

	return cnt
}

// func jump(nums []int) int {
// 	nLen := len(nums)
// 	if nLen < 2 {
// 		return 0
// 	}
// 	idx := 0
// 	step := 0
// 	for idx < nLen {
// 		fmt.Printf("idx = %d.\n", idx)
// 		if idx+nums[idx] >= nLen-1 {
// 			return step + 1
// 		}
// 		max := idx
// 		for i := idx + 1; i <= idx+nums[idx]; i++ {
// 			fmt.Printf("i = %d, nums[i] = %d.\n", i, nums[i])
// 			if i+nums[i] > max+nums[max] {
// 				max = i
// 			}
// 		}
// 		idx = max
// 		step++
// 	}
// 	return step
// }
