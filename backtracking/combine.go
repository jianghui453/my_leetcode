//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
//示例:
//
//输入: n = 4, k = 2
//输出:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]

package backtracking

func combine(n int, k int) [][]int {
	ret := make([][]int, 0)
	if k == 0 {
		return ret
	}

	if n == 0 {
		return ret
	}

	for i := n; i >= 1; i-- {
		if k == 1 {
			ret = append(ret, []int{i})
		} else {
			for _, retItem := range combine(i-1, k-1) {
				ret = append(ret, append([]int{i}, retItem...))
			}
		}
	}

	return ret
}

// func combine(n int, k int) [][]int {
// 	var backtrackCombine func(int, []int)
// 	var res = make([][]int, 0)
// 	backtrackCombine = func(first int, curr []int) {
// 		if len(curr) == k {
// 			var tmp = make([]int, k)
// 			copy(tmp, curr)
// 			res = append(res, tmp)
// 			return
// 		}
// 		for i := first; i <= n-(k-len(curr))+1; i++ {
// 			curr = append(curr, i)
// 			backtrackCombine(i+1, curr)
// 			curr = curr[:len(curr)-1]
// 		}
// 	}
// 	backtrackCombine(1, make([]int, 0))
// 	return res
// }

//func combine(n int, k int) [][]int {
//	ret := make([][]int, 0)
//	if k == 0 || n < k {
//		return ret
//	}
//	nums := make([]int, n)
//	for i := 0; i < n; i++ {
//		nums[i] = i + 1
//	}
//	recursion(nums, k, &ret, []int{})
//	return ret
//}
//
//func recursion(nums []int, k int, ret *[][]int, retItem []int) {
//	lenRetItem := len(retItem)
//	lenNums := len(nums)
//	for i, num := range nums {
//		retItemCopy := make([]int, lenRetItem)
//		copy(retItemCopy, retItem)
//		retItemCopy = append(retItemCopy, num)
//		if k == 1 {
//			*ret = append(*ret, retItemCopy)
//		} else {
//			if i < lenNums-1 {
//				recursion(nums[i+1:], k-1, ret, retItemCopy)
//			}
//		}
//	}
//}
