// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

// candidates 中的数字可以无限制重复被选取。

// 说明：

// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。
// 示例 1:

// 输入: candidates = [2,3,6,7], target = 7,
// 所求解集为:
// [
//   [7],
//   [2,2,3]
// ]
// 示例 2:

// 输入: candidates = [2,3,5], target = 8,
// 所求解集为:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]

package backtracking

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	l := len(candidates)
	if l == 0 {
		return ret
	}

	sort.Ints(candidates)

	var f func(nums []int, t int) [][]int
	f = func(nums []int, t int) [][]int {
		_ret := make([][]int, 0)
		
		for i := range nums {
			if nums[i] > t {
				break
			}

			if nums[i] == t {
				_ret = append(_ret, []int{nums[i]})
				break
			}

			for _, item := range f(nums[i: ], t-nums[i]) {
				_ret = append(_ret, append([]int{nums[i]}, item...))
			}
		}

		return _ret
	}

	ret = f(candidates, target)
	return ret
}

// import "sort"
// import "fmt"
// import "strings"
// import "strconv"

// func combinationSum(candidates []int, target int) [][]int {
// 	ret := [][]int{}
// 	sort.Ints(candidates)
// 	getSum(candidates, target, &ret, []int{})
// 	return ret
// }

// func getSum(candidates []int, target int, ret *[][]int, temp []int) {
// 	for i := 0; i < len(candidates); i++ {
// 		if candidates[i] < target {
// 			getSum(candidates[i:], target-candidates[i], ret, append(temp, candidates[i]))
// 			continue
// 		} else if candidates[i] == target {
// 			item := make([]int, len(temp)+1)
// 			copy(item, append(temp, candidates[i]))
// 			(*ret) = append((*ret), item)
// 		}
// 		break
// 	}
// }

// func combinationSum(candidates []int, target int) [][]int {
// 	ans := make([][]int, 0)
// 	n := len(candidates)
// 	sort.Ints(candidates)

// 	backtrack(&ans, []int{}, candidates, n, 0, target)
// 	return ans
// }

// func backtrack(ans *[][]int, current, candidates []int, n, idx, target int) {
// 	for i := idx; i < n; i++ {
// 		if candidates[i] < target {
// 			backtrack(ans, append(current, candidates[i]), candidates, n, i, target-candidates[i])
// 		} else if candidates[i] > target {
// 			return
// 		} else {
// 			res := make([]int, len(current)+1)
// 			copy(res, append(current, candidates[i]))
// 			*ans = append(*ans, res)
// 		}
// 	}
// }

// func combinationSum(candidates []int, target int) [][]int {
// 	if len(candidates) == 0 {
// 		return [][]int{}
// 	}
// 	sort.Ints(candidates)
// 	if candidates[0] > target {
// 		return [][]int{}
// 	}
// 	maxCount := 1
// 	sum := candidates[0]
// 	for {
// 		sum += candidates[0]
// 		if sum <= target {
// 			maxCount++
// 		} else {
// 			break
// 		}
// 	}
// 	ret := [][]int{}
// 	his := [][]int{}
// 	rptMap := make(map[string]bool)
// 	for _, i := range candidates {
// 		if i < target {
// 			his = append(his, []int{i})
// 		} else if i == target {
// 			ret = append(ret, []int{i})
// 		} else {
// 			break
// 		}
// 	}
// 	fmt.Printf("maxCount = %v\n", maxCount)
// 	for ; maxCount > 1; maxCount-- {
// 		fmt.Printf("his = %v\n", his)
// 		newHis := [][]int{}
// 		rptHisMap := make(map[string]bool)
// 		for _, nums := range his {
// 			sum := 0
// 			for _, num := range nums {
// 				sum += num
// 			}
// 			for _, i := range candidates {
// 				if sum+i < target {
// 					item := make([]int, len(nums)+1, len(nums)+1)
// 					copy(item, nums)
// 					item[len(item)-1] = i
// 					sort.Ints(item)
// 					sb := strings.Builder{}
// 					for j := 0; j < len(item); j++ {
// 						sb.WriteString(strconv.Itoa(item[j]))
// 					}
// 					if _, ok := rptHisMap[sb.String()]; !ok {
// 						newHis = append(newHis, item)
// 						rptHisMap[sb.String()] = true
// 					}
// 				} else if sum+i == target {
// 					item := make([]int, len(nums)+1, len(nums)+1)
// 					copy(item, nums)
// 					item[len(item)-1] = i
// 					sort.Ints(item)
// 					sb := strings.Builder{}
// 					for j := 0; j < len(item); j++ {
// 						sb.WriteString(strconv.Itoa(item[j]))
// 					}
// 					if _, ok := rptMap[sb.String()]; !ok {
// 						ret = append(ret, item)
// 						rptMap[sb.String()] = true
// 					}
// 				} else {
// 					break
// 				}
// 			}
// 		}
// 		his = newHis
// 	}
// 	return ret
// }
