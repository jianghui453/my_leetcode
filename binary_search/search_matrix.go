//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数。
//示例 1:
//
//输入:
//matrix = [
//  [1,   3,  5,  7],
//  [10, 11, 16, 20],
//  [23, 30, 34, 50]
//]
//target = 3
//输出: true
//示例 2:
//
//输入:
//matrix = [
//  [1,   3,  5,  7],
//  [10, 11, 16, 20],
//  [23, 30, 34, 50]
//]
//target = 13
//输出: false

package binary_search

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	if n == 0 {
		return false
	}

	for i := 0; i < m; i++ {
		if matrix[i][0] > target {
			break
		}

		if matrix[i][0] <= target && matrix[i][n-1] >= target {
			min, max := 0, n-1
			for min <= max {
				mid := (min+max)/2

				if matrix[i][mid] < target {
					min = mid+1
				} else if matrix[i][mid] > target {
					max = mid-1
				} else {
					return true
				}
			}

		}
	}

	return false
}

// func searchMatrix(matrix [][]int, target int) bool {
// 	lenRaw := len(matrix)
// 	if lenRaw < 1 {
// 		return false
// 	}
// 	lenCol := len(matrix[0])
// 	if lenCol < 1 {
// 		return false
// 	}
// 	for i := 0; i < lenRaw; i++ {
// 		if matrix[i][0] < target && matrix[i][lenCol-1] > target {
// 			for j := 1; j < lenCol-1; j++ {
// 				if matrix[i][j] == target {
// 					return true
// 				}
// 			}
// 			return false
// 		} else if matrix[i][0] == target || matrix[i][lenCol-1] == target {
// 			return true
// 		} else if matrix[i][0] > target {
// 			break
// 		}
// 	}
// 	return false
// }
