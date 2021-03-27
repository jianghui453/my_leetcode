// The set [1,2,3,...,n] contains a total of n! unique permutations.

// By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

// "123"
// "132"
// "213"
// "231"
// "312"
// "321"
// Given n and k, return the kth permutation sequence.

// Note:

// Given n will be between 1 and 9 inclusive.
// Given k will be between 1 and n! inclusive.
// Example 1:

// Input: n = 3, k = 3
// Output: "213"
// Example 2:

// Input: n = 4, k = 9
// Output: "2314"

package math_algorithm

import (
// "fmt"
// "strings"
)

func getPermutation(n int, k int) string {
	if n == 0 {
		return ""
	}

	if n == 1 {
		return "1"
	}

	var nums []byte
	cnt := 1
	for i := n; i > 1; i-- {
		cnt *= i
	}
	if k != 0 && k%cnt == 0 {
		k = cnt
	} else {
		k = k % cnt
	}

	if k <= cnt-k {
		nums = make([]byte, n)
		for i := 0; i < n; i++ {
			nums[i] = byte(i + '1')
		}

	loop1:
		for ; k > 1; k-- {
			for i := n - 1; i >= 0; i-- {
				if i > 0 && nums[i] > nums[i-1] {
					idx := i - 1
					_idx := n - 1
					for j := i + 1; j < n; j++ {
						if nums[j] <= nums[i-1] {
							_idx = j - 1
							break
						}
					}
					nums[idx], nums[_idx] = nums[_idx], nums[idx]
					for x := 0; x < (n-i)/2; x++ {
						nums[i+x], nums[n-1-x] = nums[n-1-x], nums[i+x]
					}
					continue loop1
				}
			}
		}
	} else {
		nums = make([]byte, n)
		for i := 0; i < n; i++ {
			nums[n-i-1] = byte(i + '1')
		}

	loop2:
		for ; cnt > k; cnt-- {
			for i := n - 1; i >= 0; i-- {
				if i > 0 && nums[i] < nums[i-1] {
					idx := i - 1
					_idx := n - 1

					for j := i + 1; j < n; j++ {
						if nums[j] >= nums[i-1] {
							_idx = j - 1
							break
						}
					}

					nums[idx], nums[_idx] = nums[_idx], nums[idx]

					for x := 0; x < (n-i)/2; x++ {
						nums[i+x], nums[n-1-x] = nums[n-1-x], nums[i+x]
					}
					continue loop2
				}
			}
		}
	}

	return string(nums)
}

//func getPermutation(n int, k int) string {
//	nums := make([]int, 0)
//	for i := 1; i <= n; i++ {
//		nums = append(nums, i)
//	}
//	fact := make(map[int]int)
//	fact[0] = 1
//	for i := 1; i < n; i++ {
//		fact[i] = i * fact[i-1]
//	}
//	k--
//	res := ""
//	for i := n; i >= 1; i-- {
//		pos := k / fact[i-1]
//		res += strconv.Itoa(nums[pos])
//		k = k % fact[i-1]
//		nums = sliceRm(nums, pos)
//	}
//	return res
//}
//func sliceRm(nums []int, pos int) []int {
//	res := make([]int, 0)
//	for k, v := range nums {
//		if k != pos {
//
//			res = append(res, v)
//		}
//	}
//	return res
//}

func getPermutation2(n int, k int) string {
	candidate := make([]byte, n)
	for i := 0; i < n; i++ {
		candidate[i] = '0' + byte(i+1)
	}

	fact := 1
	for i := 1; i < n; i++ {
		fact *= i
	}
	k--
	for i := 0; i < n-1; i++ {
		idx := i + k/fact
		tmp := candidate[idx]
		for j := idx; j > i; j-- {
			candidate[j] = candidate[j-1]
		}
		candidate[i] = tmp
		k %= fact
		fact /= (n - i - 1)
	}
	return string(candidate)
}

//
//func quickSort(nums []int) {
//	nLen := len(nums)
//	if nLen < 2 {
//		return
//	}
//	copyNums := make([]int, nLen)
//	copy(copyNums, nums)
//	left := 0
//	right := nLen - 1
//	for i := 1; i < nLen; i ++ {
//		if copyNums[i] > copyNums[0] {
//			nums[right] = copyNums[i]
//			right --
//		} else if copyNums[i] < copyNums[0] {
//			nums[left] = copyNums[i]
//			left ++
//		}
//	}
//	for i := left; i <= right; i ++ {
//		nums[i] = copyNums[0]
//	}
//	quickSort(nums[:left])
//	if right < nLen-1 {
//		quickSort(nums[right+1:])
//	}
//}
