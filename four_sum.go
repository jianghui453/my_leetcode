package leet_code

import (
	"fmt"
	"sort"
	"strings"
	"strconv"
)

func fourSum(nums []int, target int) [][]int {
	var i int
	var ret [][]int
	var rptMap = make(map[string]bool)
	sort.Ints(nums)
	for i = 1; i < len(nums) - 2; i ++ {
		for j := i + 1; j < len(nums) - 1; j ++ {
			left := i - 1
			right := j + 1
			for left >= 0 && right < len(nums) {
				if nums[left]+nums[i]+nums[j]+nums[right] > target {
					left--
					continue
				}
				if nums[left]+nums[i]+nums[j]+nums[right] < target {
					right++
					continue
				}
				numSlice := []int{nums[left], nums[i], nums[j], nums[right]}
				sort.Ints(numSlice)
				sb := strings.Builder{}
				for _, v := range numSlice {
					sb.WriteString(strconv.Itoa(v))
				}
				key := sb.String()
				if _, ok := rptMap[key]; !ok {
					rptMap[key] = true
					ret = append(ret, []int{nums[left], nums[i], nums[j], nums[right]})
				}
				left --
				right ++
			}
		}
	}
	return ret
}

func fourSumV1(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	var rptMap = make(map[string]bool)
	threeSum := func(nums []int, target int, num int) [][]int {
		var left, i, right int
		for i = 1; i+1 < len(nums); i++ {
			left = i - 1
			right = i + 1
			for left >= 0 && right < len(nums) {
				if nums[left]+nums[i]+nums[right] > target-num {
					left--
					continue
				}
				if nums[left]+nums[i]+nums[right] < target-num {
					right++
					continue
				}
				numSlice := []int{nums[left], nums[i], nums[right], num}
				sort.Ints(numSlice)
				sb := strings.Builder{}
				for _, v := range numSlice {
					sb.WriteString(strconv.Itoa(v))
				}
				key := sb.String()
				if _, ok := rptMap[key]; !ok {
					rptMap[key] = true
					ret = append(ret, []int{nums[left], nums[i], nums[right], num})
				}
				left --
				right ++
				// fmt.Printf("rptMap = %v\n", rptMap)
			}
		}
		return ret
	}
	for k, num := range nums {
		numsNew := []int{}
		for _, v := range nums[:k] {
			numsNew = append(numsNew, v)
		}
		for _, v := range nums[k + 1:] {
			numsNew = append(numsNew, v)
		}
		fmt.Printf("numsNew = %v; num = %d\n", numsNew, num)
		threeSum(numsNew, target, num)
	}
	return ret
}
