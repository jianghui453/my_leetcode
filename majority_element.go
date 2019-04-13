package leet_code

import (
	"fmt"
	"math"
)

func majorityElementMine(nums []int) int {
	iMap := make(map[int]int)
	iLen := int(math.Ceil(float64(len(nums)) / 2))
	for _, num := range nums {
		iMap[num]++
		if iMap[num] >= iLen {
			fmt.Printf("num=%d", num)
			return num
		}
	}
	return -1
}
