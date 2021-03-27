package sort

import (
	"fmt"
)

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	var max, min func(...int) int
	max = func(ints... int) int {
		maxInt := ints[0]
		for _, i := range ints {
			if i > maxInt {
				maxInt = i
			}
		}
		return maxInt
	}
	min = func(ints ...int) int {
		minInt := ints[0]
		for _, i := range ints {
			if i < minInt {
				minInt = i
			}
		}
		return minInt
	}

	type paires struct{
		min, max int
	}

	minNum, maxNum := min(nums...), max(nums...)
	avgGap := (maxNum - minNum) / (len(nums) - 1)
	fmt.Println("mixNum:", minNum, "maxNum:", maxNum, "avgGap:", avgGap)
	bucketWidth := max(1, avgGap)
	bucketsCnt := (maxNum - minNum) / bucketWidth + 1
	fmt.Println("bucketWidth:", bucketWidth, "bucktetsCnt:", bucketsCnt)
	buckets := make([]paires, bucketsCnt)
	for i := 0; i < bucketsCnt; i++ {
		buckets[i] = paires{
			min: -1,
			max: -1,
		}
	}
	for _, num := range nums {
		fmt.Println("num:", num)
		bucketIdx := (num - minNum) / bucketWidth
		if buckets[bucketIdx].min == -1 {
			buckets[bucketIdx].min, buckets[bucketIdx].max = num, num
		} else {
			buckets[bucketIdx].min = min(num, buckets[bucketIdx].min)
			buckets[bucketIdx].max = max(num, buckets[bucketIdx].max)
		}
	}

	maxGap := 0
	preIdx := -1
	for i := 0; i < bucketsCnt; i++ {
		if buckets[i].min == -1 {
			continue
		}
		if preIdx != -1 {
			maxGap = max(maxGap, buckets[i].min - buckets[preIdx].max)
		}
		preIdx = i
	}
	return maxGap
}
