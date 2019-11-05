// Given a collection of distinct integers, return all possible permutations.

// Example:

// Input: [1,2,3]
// Output:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

package backtracking

import (
	// "fmt"
)

func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	l := len(nums)
	if l == 0 {
		return ret
	}

	if l == 1 {
		ret = append(ret, []int{nums[0]})
		return ret
	}

	for i := 0; i < l; i++ {
		newNums := make([]int, 0)
		for j := 0; j < l; j++ {
			if j != i {
				newNums = append(newNums, nums[j])
			}
		}

		for _, item := range permute(newNums) {
			ret = append(ret, append([]int{nums[i]}, item...))
		}
	}

	return ret
}

// func permute(nums []int) [][]int {
// 	rets := [][]int{}
// 	if len(nums) > 0 {
// 		recurity(&rets, []int{}, nums)
// 	}
// 	return rets
// }

// func recurity(rets *[][]int, ret []int, nums []int) {
// 	if len(nums) == 0 {
// 		*rets = append(*rets, ret)
// 		return
// 	}

// 	for i, num := range nums {
// 		newNums := []int{}
// 		if i == 0 {
// 			newNums = append(newNums, nums[1:]...)
// 		} else if i == len(nums)-1 {
// 			newNums = append(newNums, nums[:i]...)
// 		} else {
// 			newNums = append(newNums, nums[:i]...)
// 			newNums = append(newNums, nums[i+1:]...)
// 		}
// 		newRet := make([]int, len(ret))
// 		copy(newRet, ret)
// 		recurity(rets, append(newRet, num), newNums)
// 	}
// }
