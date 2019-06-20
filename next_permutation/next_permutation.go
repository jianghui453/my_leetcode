package next_permutation

import (
	"fmt"
)

func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= nums[i+1] {
			continue
		}
		fmt.Printf("i = %d\n", i)
		var j int
		for j = i + 1; j < len(nums) - 1; j++ {
			if nums[j] <= nums[i] {
				break
			}
		}
		fmt.Printf("j = %d\n", j)
		if j == len(nums) - 1 && nums[j] > nums[i] {
			fmt.Printf("before = %v\n", nums)
			nums[i], nums[j] = nums[j], nums[i]
			fmt.Printf("after = %v\n", nums)
		} else {
			nums[i], nums[j - 1] = nums[j - 1], nums[i]
		}
		if i == len(nums) - 2 {
			return
		}
		for j = (len(nums) - 1 - i) / 2; j > 0; j-- {
			fmt.Printf("j = %d\v", j)
			nums[i+j], nums[len(nums)-j] = nums[len(nums)-j], nums[i+j]
		}
		return
	}
	for j := len(nums) / 2 - 1 ; j >= 0; j -- {
		nums[j], nums[len(nums) - 1 - j] = nums[len(nums) - 1 - j], nums[j]
	}
}
