// 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
// 注意：
// 答案中不可以包含重复的四元组。
// 示例：
// 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
// 满足要求的四元组集合为：
// [
//  [-1,  0, 0, 1],
//  [-2, -1, 1, 2],
//  [-2,  0, 0, 2]
// ]

package array

import (
	"sort"
	// "fmt"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	ret := make([][]int, 0)
	retItem := make([]int, 0)
	k := 4
	twoPointers(nums, 0, len(nums), target, k, &ret, &retItem)
	return ret
}

func twoPointers(nums []int, numsIdx, numsLen, target, k int, retPtr *[][]int, retItemPtr *[]int) {
	// fmt.Printf("nums=%v numsIdx=%d retItem=%v\n", nums, numsIdx, *retItemPtr)
	if k == 2 {
		i, j := numsIdx, numsLen-1
		for i < j {
			v := nums[i] + nums[j]
			if v > target {
				j--
			} else if v < target {
				i++
			} else {
				newRetItem := make([]int, len(*retItemPtr))
				copy(newRetItem, *retItemPtr)
				newRetItem = append(newRetItem, nums[i])
				newRetItem = append(newRetItem, nums[j])
				*retPtr = append(*retPtr, newRetItem)

				j--
				for j > i && nums[j] == nums[j+1] {
					j--
				}
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			}
		}
	} else {
		for i := numsIdx; i <= numsLen-k; i++ {
			if nums[i] > 0 && nums[i] > target {
				break
			}
			if i > numsIdx && nums[i] == nums[i-1] {
				continue
			}

			newRetItem := make([]int, len(*retItemPtr))
			copy(newRetItem, *retItemPtr)
			newRetItem = append(newRetItem, nums[i])

			twoPointers(nums, i+1, numsLen, target-nums[i], k-1, retPtr, &newRetItem)
		}
	}
}

// func fourSum(nums []int, target int) [][]int {
// 	var i int
// 	var ret [][]int
// 	var rptMap = make(map[string]bool)
// 	sort.Ints(nums)
// 	for i = 1; i < len(nums)-2; i++ {
// 		for j := i + 1; j < len(nums)-1; j++ {
// 			left := i - 1
// 			right := j + 1
// 			for left >= 0 && right < len(nums) {
// 				if nums[left]+nums[i]+nums[j]+nums[right] > target {
// 					left--
// 					continue
// 				}
// 				if nums[left]+nums[i]+nums[j]+nums[right] < target {
// 					right++
// 					continue
// 				}
// 				numSlice := []int{nums[left], nums[i], nums[j], nums[right]}
// 				sort.Ints(numSlice)
// 				sb := strings.Builder{}
// 				for _, v := range numSlice {
// 					sb.WriteString(strconv.Itoa(v))
// 				}
// 				key := sb.String()
// 				if _, ok := rptMap[key]; !ok {
// 					rptMap[key] = true
// 					ret = append(ret, []int{nums[left], nums[i], nums[j], nums[right]})
// 				}
// 				left--
// 				right++
// 			}
// 		}
// 	}
// 	return ret
// }

// func fourSum(nums []int, target int) [][]int {
// 	var i int
// 	var ret [][]int
// 	var rptMap = make(map[string]bool)
// 	sort.Ints(nums)
// 	for i = 1; i < len(nums)-2; i++ {
// 		for j := i + 1; j < len(nums)-1; j++ {
// 			left := i - 1
// 			right := j + 1
// 			for left >= 0 && right < len(nums) {
// 				if nums[left]+nums[i]+nums[j]+nums[right] > target {
// 					left--
// 					continue
// 				}
// 				if nums[left]+nums[i]+nums[j]+nums[right] < target {
// 					right++
// 					continue
// 				}
// 				numSlice := []int{nums[left], nums[i], nums[j], nums[right]}
// 				sort.Ints(numSlice)
// 				sb := strings.Builder{}
// 				for _, v := range numSlice {
// 					sb.WriteString(strconv.Itoa(v))
// 				}
// 				key := sb.String()
// 				if _, ok := rptMap[key]; !ok {
// 					rptMap[key] = true
// 					ret = append(ret, []int{nums[left], nums[i], nums[j], nums[right]})
// 				}
// 				left--
// 				right++
// 			}
// 		}
// 	}
// 	return ret
// }

// func fourSum(nums []int, target int) [][]int {
// 	sort.Ints(nums)
// 	var ret [][]int
// 	var rptMap = make(map[string]bool)
// 	threeSum := func(nums []int, target int, num int) [][]int {
// 		var left, i, right int
// 		for i = 1; i+1 < len(nums); i++ {
// 			left = i - 1
// 			right = i + 1
// 			for left >= 0 && right < len(nums) {
// 				if nums[left]+nums[i]+nums[right] > target-num {
// 					left--
// 					continue
// 				}
// 				if nums[left]+nums[i]+nums[right] < target-num {
// 					right++
// 					continue
// 				}
// 				numSlice := []int{nums[left], nums[i], nums[right], num}
// 				sort.Ints(numSlice)
// 				sb := strings.Builder{}
// 				for _, v := range numSlice {
// 					sb.WriteString(strconv.Itoa(v))
// 				}
// 				key := sb.String()
// 				if _, ok := rptMap[key]; !ok {
// 					rptMap[key] = true
// 					ret = append(ret, []int{nums[left], nums[i], nums[right], num})
// 				}
// 				left--
// 				right++
// 				// fmt.Printf("rptMap = %v\n", rptMap)
// 			}
// 		}
// 		return ret
// 	}
// 	for k, num := range nums {
// 		numsNew := []int{}
// 		for _, v := range nums[:k] {
// 			numsNew = append(numsNew, v)
// 		}
// 		for _, v := range nums[k+1:] {
// 			numsNew = append(numsNew, v)
// 		}
// 		fmt.Printf("numsNew = %v; num = %d\n", numsNew, num)
// 		threeSum(numsNew, target, num)
// 	}
// 	return ret
// }
