// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

// 示例:

// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出: 6

package stack

import (
// "fmt"
)

func trap(height []int) int {
	ret := 0
	l := len(height)
	if l <= 2 {
		return ret
	}

	s := make([]int, 0)

	for i := range height {
		// fmt.Printf("i=%d s=%v ret=%d\n", i, s, ret)
		if len(s) == 0 {
			s = append(s, i)
			continue
		}

		if height[i] < height[s[len(s)-1]] {
			s = append(s, i)
			continue
		}

		sLen := len(s)
		if sLen == 1 {
			s = []int{i}
			continue
		}

		if height[i] >= height[s[0]] {
			for j := 0; j < sLen-1; j++ {
				ret += (i - s[j] - 1) * (height[s[j]] - height[s[j+1]])
			}
			s = []int{i}
		} else {
			for j := sLen - 2; j >= 0; j-- {
				if height[i] >= height[s[j]] {
					ret += (i - s[j] - 1) * (height[s[j]] - height[s[j+1]])
				} else {
					ret += (i - s[j] - 1) * (height[i] - height[s[j+1]])
				}

				if height[i] == height[s[j]] {
					s = s[:j]
					break
				}

				if height[i] < height[s[j]] {
					s = s[:j+1]
					break
				}
			}

			s = append(s, i)
		}
	}

	return ret
}

// func trap(height []int) int {
// 	if len(height) < 2 {
// 		return 0
// 	}
// 	left := 0
// 	area, ret := 0, 0
// 	for i := 1; i < len(height); i++ {
// 		if height[i] < height[left] {
// 			area += height[left] - height[i]
// 			continue
// 		} else {
// 			ret += area
// 			area = 0
// 			left = i
// 		}
// 	}
// 	if left < len(height)-2 {
// 		right := len(height) - 1
// 		area = 0
// 		for i := right - 1; i >= left; i-- {
// 			if height[i] < height[right] {
// 				area += height[right] - height[i]
// 			} else {
// 				ret += area
// 				area = 0
// 				right = i
// 			}
// 		}
// 	}
// 	return ret
// }
