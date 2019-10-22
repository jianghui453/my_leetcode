package jump

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	for {
		if len(nums) < 2 {
			return true
		}
		maxVal := nums[0]
		maxIdx := 0
		for i := 1; i <= nums[0]; i++ {
			canIdx := nums[i] + i
			if canIdx >= len(nums)-1 {
				return true
			}
			if canIdx > maxVal {
				maxIdx = i
				maxVal = canIdx
			}
		}
		if maxIdx == 0 {
			return false
		}
		nums = nums[maxIdx:]
	}
}

func canJumpV1(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	if nums[0]+1 >= len(nums) {
		return true
	}
	for i := 1; i <= nums[0]; i++ {
		if canJump(nums[i:]) {
			return true
		}
	}
	return false
}
