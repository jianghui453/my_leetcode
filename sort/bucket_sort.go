package sort

import (
	gosort "sort"
)

func bucketSort(nums []float64) {
	l := len(nums)
	bucket := make([][]float64, 10)

	for i := 0; i < l; i++ {
		bucket[int(nums[i]*10)] = append(bucket[int(nums[i]*10)], nums[i])
	}

	k := 0
	for i := 0; i < 10; i++ {
		gosort.Float64s(bucket[i])
		for j := range bucket[i] {
			nums[k] = bucket[i][j]
			k++
		}
	}
}
