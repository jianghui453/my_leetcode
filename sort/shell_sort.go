package sort

import (
	"fmt"
)

func shellSort(nums []int) {
	l := len(nums)
	for i := l/2; i > 0; i /= 2 {
		for j := i; j < l; j++ {
			k, v := j, nums[j]
			for ; k >= i && nums[k-i] > v; k -= i {
				nums[k] = nums[k-i]
			}
			nums[k] = v
		}
	}
}
