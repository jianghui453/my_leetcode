//给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
//
//示例 1:
//
//输入:
//[
//  [1,1,1],
//  [1,0,1],
//  [1,1,1]
//]
//输出:
//[
//  [1,0,1],
//  [0,0,0],
//  [1,0,1]
//]
//示例 2:
//
//输入:
//[
//  [0,1,2,0],
//  [3,4,5,2],
//  [1,3,1,5]
//]
//输出:
//[
//  [0,0,0,0],
//  [0,4,5,0],
//  [0,3,1,0]
//]
//进阶:
//
//一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
//一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
//你能想出一个常数空间的解决方案吗？
package array

func setZeroes(matrix [][]int) {
	m := len(matrix)
	if m == 0 {
		return
	}

	n := len(matrix[0])
	if n == 0 {
		return
	}

	hashM := make(map[int]bool)
	hashN := make(map[int]bool)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				hashM[i] = true
				hashN[j] = true
			}
		}
	}

	for k := range hashM {
		for j := 0; j < n; j++ {
			matrix[k][j] = 0
		}
	}

	for k := range hashN {
		for i := 0; i < m; i++ {
			matrix[i][k] = 0
		}
	}
}

// func setZeroes(matrix [][]int) {
// 	m := len(matrix)
// 	if m == 0 {
// 		return
// 	}
// 	n := len(matrix[0])
// 	if n == 0 {
// 		return
// 	}
// 	record := make([]int, 0)
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if matrix[i][j] == 0 {
// 				record = append(record, i*n+j)
// 			}
// 		}
// 	}
// 	for _, idx := range record {
// 		raw := idx / n
// 		col := idx % n
// 		for i := 0; i < m; i++ {
// 			matrix[i][col] = 0
// 		}
// 		for j := 0; j < n; j++ {
// 			matrix[raw][j] = 0
// 		}
// 	}
// }
