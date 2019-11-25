// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

// candidates 中的每个数字在每个组合中只能使用一次。

// 说明：

// 所有数字（包括目标数）都是正整数。
// 解集不能包含重复的组合。
// 示例 1:

// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 所求解集为:
// [
//   [1, 7],
//   [1, 2, 5],
//   [2, 6],
//   [1, 1, 6]
// ]
// 示例 2:

// 输入: candidates = [2,5,2,1,2], target = 5,
// 所求解集为:
// [
//   [1,2,2],
//   [5]
// ]

package backtracking

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	l := len(candidates)
	ret := make([][]int, 0)
	if l == 0 {
		return ret
	}

	sort.Ints(candidates)

	var f func(nums []int, t int) [][]int
	f = func(nums []int, t int) [][]int {
		_l := len(nums)
		_ret := make([][]int, 0)

		for i := 0; i < _l; i++ {
			if nums[i] > t {
				break
			}

			if i > 0 && nums[i] == nums[i-1] {
				continue
			}

			if nums[i] == t {
				_ret = append(_ret, []int{nums[i]})
				break
			}

			if i < _l-1 {
				for _, item := range f(nums[i+1:], t-nums[i]) {
					_ret = append(_ret, append([]int{nums[i]}, item...))
				}
			}
		}

		return _ret
	}

	ret = f(candidates, target)
	return ret
}

// func combinationSum2(candidates []int, target int) [][]int {
// 	sort.Ints(candidates)
// 	var ret [][]int
// 	var his []int
// 	sum(&ret, his, candidates, target)
// 	return ret
// }

// func sum(ret *[][]int, his []int, candidates []int, target int) {
// 	for i := 0; i < len(candidates); i++ {
// 		if i > 0 && candidates[i] == candidates[i-1] {
// 			continue
// 		}
// 		if candidates[i] < target && i < len(candidates)-1 {
// 			newHis := make([]int, len(his)+1)
// 			copy(newHis, his)
// 			newHis[len(his)] = candidates[i]
// 			sum(ret, newHis, candidates[i+1:], target-candidates[i])
// 		} else if candidates[i] == target {
// 			newHis := make([]int, len(his)+1)
// 			copy(newHis, his)
// 			newHis[len(his)] = candidates[i]
// 			(*ret) = append((*ret), newHis)
// 			break
// 		} else {
// 			break
// 		}
// 	}
// }
