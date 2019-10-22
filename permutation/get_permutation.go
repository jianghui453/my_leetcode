package permutation

import (
	"fmt"
	"strings"
)

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

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	for ; k > 1; k-- {
		fmt.Printf("nums=%v\n", nums)
		for i := n - 1; i > 0; i-- {
			if nums[i] > nums[i-1] {

			}
		}
	}
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteByte(byte(nums[i] + '0'))
	}
	return sb.String()
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
