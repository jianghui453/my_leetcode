package sort

import (
// "fmt"
// "math"
)

func radioSort(nums []int) {
	nLen := len(nums)
	if nLen < 2 {
		return
	}

	i, cntBucket := 1, 10
	for {
		bucket := make([][]int, cntBucket)
		for j := 0; j < nLen; j++ {
			idxBucket := nums[j] / i % 10
			bucket[idxBucket] = append(bucket[idxBucket], nums[j])
		}

		if len(bucket[0]) == nLen {
			break
		}

		idx := 0
		for x := 0; x < cntBucket; x++ {
			lenBucket := len(bucket[x])
			for y := 0; y < lenBucket; y++ {
				nums[idx] = bucket[x][y]
				idx++
			}
		}

		i *= 10
	}
}
