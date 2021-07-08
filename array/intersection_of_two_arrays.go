// 给定两个数组，编写一个函数来计算它们的交集。

// 示例 1:

// 输入: nums1 = [1,2,2,1], nums2 = [2,2]
// 输出: [2]
// 示例 2:

// 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出: [9,4]
// 说明:

// 输出结果中的每个元素一定是唯一的。
// 我们可以不考虑输出结果的顺序。

package array

import (
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	var (
		l1, l2 int   = len(nums1), len(nums2)
		ret    []int = make([]int, 0)
		i1, i2 int   = 0, 0
	)

	sort.Ints(nums1)
	sort.Ints(nums2)

	for i1 < l1 && i2 < l2 {
		if nums1[i1] > nums2[i2] {
			i2++
		} else if nums1[i1] < nums2[i2] {
			i1++
		} else {
			if len(ret) == 0 || nums1[i1] != ret[len(ret)-1] {
				ret = append(ret, nums1[i1])
			}
			i1, i2 = i1+1, i2+1
		}
	}

	return ret
}
