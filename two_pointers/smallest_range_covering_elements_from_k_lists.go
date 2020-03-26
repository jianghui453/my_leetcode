// 你有 k 个升序排列的整数数组。找到一个最小区间，使得 k 个列表中的每个列表至少有一个数包含在其中。

// 我们定义如果 b-a < d-c 或者在 b-a == d-c 时 a < c，则区间 [a,b] 比 [c,d] 小。

// 示例 1:

// 输入:[[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
// 输出: [20,24]
// 解释: 
// 列表 1：[4, 10, 15, 24, 26]，24 在区间 [20,24] 中。
// 列表 2：[0, 9, 12, 20]，20 在区间 [20,24] 中。
// 列表 3：[5, 18, 22, 30]，22 在区间 [20,24] 中。
// 注意:

// 给定的列表可能包含重复元素，所以在这里升序表示 >= 。
// 1 <= k <= 3500
// -105 <= 元素的值 <= 105
// 对于使用Java的用户，请注意传入类型已修改为List<List<Integer>>。重置代码模板后可以看到这项改动。

package two_pointers

import (
	// "math"
	// "fmt"
)

func smallestRange(nums [][]int) []int {
	k := len(nums)
	idxs := make([]int, k)
	ret := make([]int, 2)
	min, mins, max := 0, 1, 2

	if k == 1 {
		if len(nums[0]) == 1 {
			ret = []int{nums[0][0], nums[0][0]}
		} else {
			ret = []int{nums[0][0], nums[0][1]}
			for i := range nums[0] {
				if i > 1 {
					if nums[0][i] - nums[0][i-1] < ret[1] - ret[0] {
						ret = []int{nums[0][i-1], nums[0][i]}
					}
				}
			}
		}
	} else if k == 2 {
		if nums[0][0] < nums[1][0] {
			ret = []int{nums[0][0], nums[1][0]}
		} else {
			ret = []int{nums[1][0], nums[0][0]}
		}
		
		for idxs[min] < len(nums[min]) {
			min, max = 0, 1
			if nums[1][idxs[1]] < nums[0][idxs[0]] {
				min, max = max, min
			}
			if nums[max][idxs[max]] - nums[min][idxs[min]] < ret[1] - ret[0] {
				ret = []int{nums[min][idxs[min]], nums[max][idxs[max]]}
			}
			idxs[min]++
		}
	} else {
		// 初始化
		for i := range nums {
			if nums[min][idxs[min]] > nums[i][idxs[i]] {
				min = i
			} else if nums[max][idxs[max]] < nums[i][idxs[i]] {
				max = i
			}
		}
		mins = max
		for i := range nums {
			if i != min && nums[i][idxs[i]] < nums[mins][idxs[mins]] {
				mins = i
			}
		}

		ret[0], ret[1] = nums[min][idxs[min]], nums[max][idxs[max]]
		idxs[min]++
		// fmt.Println(min, mins, max, nums[min][idxs[min]], nums[mins][idxs[mins]], nums[max][idxs[max]])

		for idxs[min] < len(nums[min]) && ret[0] < ret[1] {
			if nums[min][idxs[min]] > nums[mins][idxs[mins]] {
				if nums[min][idxs[min]] > nums[max][idxs[max]] {
					min, mins, max = mins, max, min
				} else {
					min, mins = mins, min
				}
				
				for i := range nums {
					if i != min && nums[i][idxs[i]] < nums[mins][idxs[mins]] {
						mins = i
					}
				}
			}
			// fmt.Println(min, mins, max, nums[min][idxs[min]], nums[mins][idxs[mins]], nums[max][idxs[max]])
			if (nums[max][idxs[max]] - nums[min][idxs[min]]) < (ret[1] - ret[0]) {
				ret[0], ret[1] = nums[min][idxs[min]], nums[max][idxs[max]]
			}

			idxs[min]++
		}
	}

	return ret
}
