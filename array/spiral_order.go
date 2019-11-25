// Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

// Example 1:

// Input:
// [
//  [ 1, 2, 3 ],
//  [ 4, 5, 6 ],
//  [ 7, 8, 9 ]
// ]
// Output: [1,2,3,6,9,8,7,4,5]
// Example 2:

// Input:
// [
//   [1, 2, 3, 4],
//   [5, 6, 7, 8],
//   [9,10,11,12]
// ]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]

package array

func spiralOrder(matrix [][]int) []int {
	ret := make([]int, 0)
	lm := len(matrix)
	if lm == 0 {
		return ret
	}

	ln := len(matrix[0])
	if ln == 0 {
		return ret
	}

	top, bottom, left, right := 0, lm-1, 0, ln-1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			ret = append(ret, matrix[top][i])
		}

		if top == bottom {
			break
		}

		for i := top + 1; i <= bottom; i++ {
			ret = append(ret, matrix[i][right])
		}

		for i := right - 1; i >= left; i-- {
			ret = append(ret, matrix[bottom][i])
		}

		if left == right {
			break
		}

		for i := bottom - 1; i > top; i-- {
			ret = append(ret, matrix[i][left])
		}

		top++
		bottom--
		left++
		right--
	}

	return ret
}

// func spiralOrder(matrix [][]int) []int {
// 	var ret []int
// 	m := len(matrix)
// 	if m == 0 {
// 		return ret
// 	}
// 	n := len(matrix[0])
// 	if n == 0 {
// 		return ret
// 	}
// 	for m > 0 && n > 0 {
// 		fmt.Printf("m=%d n=%d matrix=%v\n", m, n, matrix)
// 		if m == 1 {
// 			return append(ret, matrix[0]...)
// 		}
// 		if n == 1 {
// 			for i := 0; i < m; i++ {
// 				ret = append(ret, matrix[i][0])
// 			}
// 			return ret
// 		}
// 		ret = append(ret, matrix[0]...)
// 		for i := 1; i < m; i++ {
// 			ret = append(ret, matrix[i][n-1])
// 		}
// 		for i := n - 2; i >= 0; i-- {
// 			ret = append(ret, matrix[m-1][i])
// 		}
// 		for i := m - 2; i > 0; i-- {
// 			ret = append(ret, matrix[i][0])
// 		}
// 		if m == 2 || n == 2 {
// 			return ret
// 		}
// 		matrix = matrix[1 : m-1]
// 		for i := 0; i < m-2; i++ {
// 			fmt.Printf("matrix=%v i=%d\n", matrix, i)
// 			matrix[i] = matrix[i][1 : n-1]
// 		}
// 		m -= 2
// 		n -= 2
// 	}
// 	return ret
// }
