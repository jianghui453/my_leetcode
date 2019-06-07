package leet_code

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ret [][]int
	var rptMap = make(map[string]bool)

	sort.Ints(nums)

	for i := 1; i+1 < len(nums); i++ {
		left := i - 1
		right := i + 1
		for left >= 0 && right < len(nums) {
			if nums[left]+nums[i]+nums[right] > 0 {
				left--
				continue
			}
			if nums[left]+nums[i]+nums[right] < 0 {
				right++
				continue
			}
			key := fmt.Sprintf("%d-%d", nums[left], nums[right])
			fmt.Printf("rptMap = %v; key = %s\n", rptMap, key)
			if _, ok := rptMap[key]; !ok {
				ret_item := []int{
					nums[left],
					nums[i],
					nums[right],
				}
				ret = append(ret, ret_item)
				rptMap[key] = true
			}
			left--
			right++
		}
	}
	return ret
}
