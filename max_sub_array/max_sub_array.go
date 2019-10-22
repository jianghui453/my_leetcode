package max_sub_array

import "fmt"

func maxSubArray(nums []int) int {
	nLen := len(nums)
	if nLen == 0 {
		return 0
	}
	max := nums[0]
	sum := nums[0]
	if nums[0] < 0 {
		sum = 0
	}
	for i := 1; i < nLen; i++ {
		sum += nums[i]
		if sum <= 0 {
			if nums[i] > max {
				max = nums[i]
			}
			sum = 0
			continue
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func maxSubArray1(nums []int) int {
	nLen := len(nums)
	if nLen == 0 {
		return 0
	}
	max := nums[0]
	for i := 0; i < nLen-1; i++ {
		sum := nums[i]
		if sum > max {
			max = sum
		}
		_maxSubArray(nums[i+1:], nLen-1-i, sum, &max)
	}
	if nums[nLen-1] > max {
		max = nums[nLen-1]
	}
	return max
}

func _maxSubArray(nums []int, nLen, sum int, max *int) {
	fmt.Printf("nums=%v nLen=%d sum=%d max=%d\n", nums, nLen, sum, *max)
	if nLen == 0 {
		return
	}
	sum += nums[0]
	if sum > *max {
		*max = sum
	}
	if nLen > 1 {
		_maxSubArray(nums[1:], nLen-1, sum, max)
	}
}
