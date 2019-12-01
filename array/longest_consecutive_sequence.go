//给定一个未排序的整数数组，找出最长连续序列的长度。
//
//要求算法的时间复杂度为 O(n)。
//
//示例:
//
//输入: [100, 4, 200, 1, 3, 2]
//输出: 4
//解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

package array

import (
	"math"
	// "fmt"
)

func longestConsecutive(nums []int) int {
	ret, littler, bigger, l := 0, make(map[int]int), make(map[int]int), len(nums)
	for i := 0; i < l; i++ {
		num := nums[i]
		for {
			if _, ok := bigger[num+1]; ok {
				num = bigger[num+1]
			} else {
				break
			}
		}
		
		bigger[nums[i]] = num

		num = nums[i]
		for {
			if  _, ok := littler[num-1]; ok {
				num = littler[num-1]
			} else {
				break
			}
		}
		littler[nums[i]] = num

		ret = int(math.Max(float64(ret), float64(bigger[nums[i]]-littler[nums[i]]+1)))
	}

	return ret
}

// func longestConsecutive(nums []int) int {
//    numsLen := len(nums)
//    if numsLen == 0 {
//        return 0
//    }
//    record := make(map[int]int)
//    maxRecord := make(map[int]int)
//    max := 0
//    for i := range nums {
//        cur := nums[i]
//        if _, ok := record[cur]; ok {
//            continue
//        }

//        record[cur] = cur
//        maxRecord[cur] = 1
//        if maxRecord[cur] > max {
//            max = maxRecord[cur]
//        }
//        if _, ok := record[cur+1]; ok {
//            record[cur+1] = cur
//            maxRecord[cur] = maxRecord[cur+1]+1
//            if maxRecord[cur] > max {
//                max = maxRecord[cur]
//            }
//        }

//        j := cur-1
//        if _, ok := record[j]; ok {
//            for {
//                curRoot := record[j]
//                if curRoot == j {
//                    record[cur] = j
//                    maxRecord[j] += maxRecord[cur]
//                    if maxRecord[j] > max {
//                        max = maxRecord[j]
//                    }
//                    break
//                }
//                j = curRoot
//            }
//        }
//    }
//    return max
// }
