// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
// 示例:

// 现有矩阵 matrix 如下：

// [
//   [1,   4,  7, 11, 15],
//   [2,   5,  8, 12, 19],
//   [3,   6,  9, 16, 22],
//   [10, 13, 14, 17, 24],
//   [18, 21, 23, 26, 30]
// ]
// 给定 target = 5，返回 true。

// 给定 target = 20，返回 false。

package array

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
		if matrix[i][n-1] < target {
			continue
		}

		min, max := 0, n-1
		for min <= max {
			mid := (min+max)/2
			if matrix[i][mid] > target {
				max = mid-1
			} else if matrix[i][mid] < target {
				min = mid+1
			} else {
				return true
			}
		}
	}

	return false
}
