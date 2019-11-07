// Given an array of non-negative integers, you are initially positioned at the first index of the array.

// Each element in the array represents your maximum jump length at that position.

// Determine if you are able to reach the last index.

// Example 1:

// Input: [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Example 2:

// Input: [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum
// 			 jump length is 0, which makes it impossible to reach the last index.
			 
package greedy

func canJump(nums []int) bool {
	l := len(nums)

	if l <= 1 {
		return true
	}

	i := 0
	for nums[i] > 0 {
		if i+nums[i] >= l-1 {
			return true
		}

		maxL := -1
		maxI := i

		for j := i+1; j <= i+nums[i]; j++ {
			if j+nums[j] > maxL {
				maxL = j+nums[j]
				maxI = j
			}
		}

		i = maxI
	}

	return false
}

// func canJump(nums []int) bool {
// 	if len(nums) < 2 {
// 		return true
// 	}
// 	for {
// 		if len(nums) < 2 {
// 			return true
// 		}
// 		maxVal := nums[0]
// 		maxIdx := 0
// 		for i := 1; i <= nums[0]; i++ {
// 			canIdx := nums[i] + i
// 			if canIdx >= len(nums)-1 {
// 				return true
// 			}
// 			if canIdx > maxVal {
// 				maxIdx = i
// 				maxVal = canIdx
// 			}
// 		}
// 		if maxIdx == 0 {
// 			return false
// 		}
// 		nums = nums[maxIdx:]
// 	}
// }

// func canJump(nums []int) bool {
// 	if len(nums) < 2 {
// 		return true
// 	}
// 	if nums[0]+1 >= len(nums) {
// 		return true
// 	}
// 	for i := 1; i <= nums[0]; i++ {
// 		if canJump(nums[i:]) {
// 			return true
// 		}
// 	}
// 	return false
// }
