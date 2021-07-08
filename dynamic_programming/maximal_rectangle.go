//给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
//
//示例:
//
//输入:
//[
//  ["1","0","1","0","0"],
//  ["1","0","1","1","1"],
//  ["1","1","1","1","1"],
//  ["1","0","0","1","0"]
//]
//输出: 6

package dynamic_programming

import (
	"math"
)

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	area := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			if j > 0 {
				matrix[i][j] += (matrix[i][j-1] - '0')
			}

			min := matrix[i][j]
			for _i := i; _i >= 0; _i-- {
				if matrix[_i][j] == '0' {
					break
				}

				min = byte(math.Min(float64(min), float64(matrix[_i][j])))

				area = int(math.Max(float64(int(min-'0')*(i-_i+1)), float64(area)))
			}
		}
	}

	return area
}

// func maximalRectangle(matrix [][]byte) int {
// 	m := len(matrix)
// 	if m < 1 {
// 		return 0
// 	}
// 	n := len(matrix[0])
// 	if n < 1 {
// 		return 0
// 	}
// 	area := 0
// 	for y := 0; y < m; y++ {
// 		heights := make([]int, n)
// 		for j := 0; j < n; j++ {
// 			for i := y; i < m; i++ {
// 				if matrix[i][j] == byte('0') {
// 					break
// 				}
// 				heights[j]++
// 			}
// 		}
// 		s := []int{-1}
// 		lenS := 1
// 		for idxH := 0; idxH < n; idxH++ {
// 			for lenS > 1 && heights[idxH] < heights[s[lenS-1]] {
// 				width := idxH - s[lenS-2] - 1
// 				height := heights[s[lenS-1]]
// 				areaI := width * height
// 				if areaI > area {
// 					area = areaI
// 				}
// 				s = s[:lenS-1]
// 				lenS--
// 			}
// 			s = append(s, idxH)
// 			lenS++
// 		}
// 		for lenS > 1 {
// 			width := n - s[lenS-2] - 1
// 			height := heights[s[lenS-1]]
// 			areaI := width * height
// 			if areaI > area {
// 				area = areaI
// 			}
// 			s = s[:lenS-1]
// 			lenS--
// 		}
// 	}
// 	return area
// }
