//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]

package two_pointers

import "sort"

// import "fmt"

func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	if numsLen < 3 {
		return [][]int{}
	}

	sort.Ints(nums)

	ret := make([][]int, 0)

	for i := 0; i < numsLen; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 {
			break
		}

		j, k := i+1, numsLen-1
		for j < k {
			v := nums[i] + nums[j] + nums[k]

			if v < 0 {
				j++
			} else if v > 0 {
				k--
			} else {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})

				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}

				k--
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}

	return ret
}

//func threeSum(nums []int) [][]int {
//	var ret [][]int
//	var rptMap = make(map[string]bool)
//
//	sort.Ints(nums)
//
//	for i := 1; i+1 < len(nums); i++ {
//		left := i - 1
//		right := i + 1
//		for left >= 0 && right < len(nums) {
//			if nums[left]+nums[i]+nums[right] > 0 {
//				left--
//				continue
//			}
//			if nums[left]+nums[i]+nums[right] < 0 {
//				right++
//				continue
//			}
//			key := fmt.Sprintf("%d-%d", nums[left], nums[right])
//			fmt.Printf("rptMap = %v; key = %s\n", rptMap, key)
//			if _, ok := rptMap[key]; !ok {
//				ret_item := []int{
//					nums[left],
//					nums[i],
//					nums[right],
//				}
//				ret = append(ret, ret_item)
//				rptMap[key] = true
//			}
//			left--
//			right++
//		}
//	}
//	return ret
//}
