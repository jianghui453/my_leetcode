// 给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

// 示例 1:

// 输入: [10,2]
// 输出: 210
// 示例 2:

// 输入: [3,30,34,5,9]
// 输出: 9534330
// 说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。

package array

import (
	"strconv"
	// "fmt"
)

func largestNumber(nums []int) string {
	var quickSort func ([]string)
	quickSort = func (strs []string) {
		var (
			l = len(strs)
			strsCopy = make([]string, l)
			left, right = 0, l-1
		)
	
		if l <= 1 {
			return
		}
	
		copy(strsCopy, strs)
	
		for i := 1; i < l; i++ {
			if strsCopy[i] == strsCopy[0] {
				continue
			}
	
			if bigger(strsCopy[0], strsCopy[i]) {
				strs[left] = strsCopy[i]
				left++
			} else {
				strs[right] = strsCopy[i]
				right--
			}
		}
	
		for i := left; i <= right; i++ {
			strs[i] = strsCopy[0]
		}
		
		quickSort(strs[:left])
		if right < l-1 {
			quickSort(strs[right+1:])
		}
	}

	var (
		l = len(nums)
		strs []string
		ret string
		allZero = true
	)

	for i := 0; i < l; i++ {
		strs = append(strs, strconv.Itoa(nums[i]))
		if nums[i] > 0 {
			allZero = false
		}
	}
	if allZero {
		return "0"
	}
	
	quickSort(strs)
	
	for i := l-1; i >= 0; i-- {
		ret += strs[i]
	}
	
	return ret
}

func bigger(x, y string) bool {
	if x == y {
		return true
	}
	
	var (
		xLen, yLen = len(x), len(y)
		i = 0
	)

	for ; i < xLen && i < yLen; i++ {
		if x[i] > y[i] {
			return true
		}

		if x[i] < y[i] {
			return false
		}
	}
	
	if i < xLen {
		return bigger(x[i:], y)
	}
	
	return bigger(x, y[i:])
}
