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

func longestConsecutive(nums []int) int {
	hash := make(map[int]int)
	var max int

	for _, v := range nums {
		if _, ok := hash[v]; !ok {
			left := hash[v-1]
			right := hash[v+1]
			curry := 1 + left + right
			if curry > max {
				max = curry
			}
			hash[v] = curry
			hash[v-left] = curry
			hash[v+right] = curry
		}
	}
	return max
}

//func longestConsecutive(nums []int) int {
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
//
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
//
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
//}
