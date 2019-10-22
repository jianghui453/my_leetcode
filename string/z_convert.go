package string

import (
	"strings"
)

func convert(s string, numRows int) string {
	sLen := len(s)
	if numRows < 2 || sLen <= numRows {
		return s
	}
	matrix := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		matrix[i] = make([]int, sLen)
	}
	i, j := 0, 0
	inc := false
	for k := 0; k < sLen; k++ {
		matrix[i][j] = int(s[k]) + 1
		if inc {
			if i == numRows-1 {
				inc = !inc
				i--
				j++
			} else {
				i++
			}
		} else {
			if i == 0 {
				inc = !inc
				i++
			} else {
				i--
				j++
			}
		}
	}
	sb := new(strings.Builder)
	for i := 0; i < numRows; i++ {
		for j := 0; j < sLen; j++ {
			if matrix[i][j] > 0 {
				sb.WriteByte(byte(matrix[i][j] - 1))
			}
		}
	}
	return sb.String()
}

// func convert(s string, numRows int) string {
// 	sLen := len(s)
// 	if sLen <= numRows || numRows <= 1 {
// 		return s
// 	}
// 	sSlice := make([]strings.Builder, numRows, numRows)
// 	offset := 0
// 	increaseFlag := true
// 	for i := 0; i < sLen; i ++ {
// 		sSlice[offset].WriteByte(s[i])
// 		if offset == numRows - 1 {
// 			increaseFlag = false
// 		}
// 		if offset == 0 {
// 			increaseFlag = true
// 		}
// 		if increaseFlag {
// 			offset ++
// 		} else {
// 			offset --
// 		}
// 	}
// 	for i := 1; i < numRows; i ++ {
// 		sSlice[0].WriteString(sSlice[i].String())
// 	}
// 	return sSlice[0].String()
// }

// func convert(s string, numRows int) string {
// 	sLen := len(s)
// 	if sLen <= numRows || numRows <= 1 {
// 		return s
// 	}
// 	var sb strings.Builder
// 	for floor := 1; floor <= numRows; floor ++ {
// 		for i := 0; i < sLen + numRows; i += (numRows - 1) * 2 {
// 			offset1 := i - (floor - 1)
// 			if offset1 >= 0 && offset1 < sLen {
// 				sb.WriteByte(s[offset1])
// 			}
// 			offset2 := i + (floor - 1)
// 			if floor != numRows && floor != 1 && offset2 < sLen {
// 				sb.WriteByte(s[offset2])
// 			}
// 		}
// 	}
// 	return sb.String()
// }
