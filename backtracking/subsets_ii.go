//给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
//输入: [1,2,2]
//输出:
//[
//  [2],
//  [1],
//  [1,2,2],
//  [2,2],
//  [1,2],
//  []
//]

package backtracking

import (
	// "fmt"
	// "strconv"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}
	
	sort.Ints(nums)

	var f func(ns []int) [][]int
	f = func(ns []int) [][]int {
		ln := len(ns)

		r := make([][]int, 0)
		r = append(r, make([]int, 0))
		if ln == 0 {
			return r
		}

		for i := 0; i < ln; i++ {
			if i > 0 && ns[i] == ns[i-1] {
				continue
			}

			var _r [][]int
			if i == ln-1 {
				_r = [][]int{[]int{}}
			} else {
				_r = f(ns[i+1: ])
			}
			for _, item := range _r {
				r = append(r, append([]int{ns[i]}, item...))
			}
		}

		return r
	}

	ret := f(nums)
	return ret
}

// func subsetsWithDup(nums []int) [][]int {
// 	l := len(nums)
// 	r := make([][]int, 0)
// 	r = append(r, []int{})
// 	if l < 1 {
// 		return r
// 	}
// 	quickSort(nums)
// 	fmt.Printf("nums=%v\n", nums)
// 	memo := make(map[string]bool)
// 	var f func([]int, []int, int, int, string)
// 	f = func(n, rItem []int, nL, rL int, k string) {
// 		if nL == 0 {
// 			return
// 		}
// 		for i := 0; i < nL; i++ {
// 			rItemCopy := make([]int, rL+1)
// 			copy(rItemCopy, rItem)
// 			rItemCopy[rL] = n[i]
// 			key := k + strconv.Itoa(n[i])
// 			if _, ok := memo[key]; !ok {
// 				r = append(r, rItemCopy)
// 				memo[key] = true
// 			}
// 			if i < nL-1 {
// 				f(n[i+1:], rItemCopy, nL-1-i, rL+1, key)
// 			}
// 		}
// 	}
// 	f(nums, []int{}, l, 0, "")
// 	return r
// }

// func quickSort(nums []int) {
// 	l := len(nums)
// 	if l < 2 {
// 		return
// 	}
// 	n := make([]int, l)
// 	copy(n, nums)
// 	le := 0
// 	ri := l - 1
// 	for i := 1; i < l; i++ {
// 		if n[i] > n[0] {
// 			nums[ri] = n[i]
// 			ri--
// 		} else if n[i] < n[0] {
// 			nums[le] = n[i]
// 			le++
// 		}
// 	}
// 	for i := le; i <= ri; i++ {
// 		nums[i] = n[0]
// 	}
// 	if le > 0 {
// 		quickSort(nums[:le])
// 	}
// 	if ri < l-1 {
// 		quickSort(nums[ri+1:])
// 	}
// }
