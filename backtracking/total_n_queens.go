// The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.



// Given an integer n, return the number of distinct solutions to the n-queens puzzle.

// Example:

// Input: 4
// Output: 2
// Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
// [
//  [".Q..",  // Solution 1
//   "...Q",
//   "Q...",
//   "..Q."],

//  ["..Q.",  // Solution 2
//   "Q...",
//   "...Q",
//   ".Q.."]
// ]

package backtracking

func totalNQueens(n int) int {
	ret := 0
	if n == 0 {
		return ret
	}

	horizontal := make([]bool, n)
	vertical := make([]bool, n)

	var f func(x, y, cnt int, strs [][]byte)
	f = func(x, y, cnt int, strs [][]byte) {
		if cnt < x {
			return
		}

		if x == n {
			if cnt != n {
				return
			}
			ret++
			return
		}

		_x, _y := x, y
		if _y == n-1 {
			_x++
			_y = 0
		} else {
			_y++
		}

		strs[x][y] = '.'
		f(_x, _y, cnt, strs)

		if horizontal[x] || vertical[y] {
			return
		}
		for i := 1; x-i >= 0 && y-i >= 0; i++ {
			if strs[x-i][y-i] == 'Q' {
				return
			}
		}
		for i := 1; x-i >= 0 && y+i < n; i++ {
			if strs[x-i][y+i] == 'Q' {
				return
			}
		}

		horizontal[x] = true
		vertical[y] = true
		strs[x][y] = 'Q'
		
		f(_x, _y, cnt+1, strs)

		horizontal[x] = false
		vertical[y] = false
		strs[x][y] = '.'
	}

	_strs := make([][]byte, n)
	for i := 0; i < n; i++ {
		_strs[i] = make([]byte, n)
	}
	f(0, 0, 0, _strs)
	return ret
}

// func totalNQueens(n int) int {
// 	ret := 0
// 	if n > 0 {
// 		column := make([]int, n)
// 		strs := make([][]int, n)
// 		for i := 0; i < n; i++ {
// 			str := make([]int, n)
// 			strs[i] = str
// 		}
// 		totalNQueens_(n, 0, column, strs, &ret)
// 	}
// 	return ret
// }

// func totalNQueens_(n, y int, column []int, strs [][]int, ret *int) {
// 	if y >= n {
// 		*ret++
// 		return
// 	}
// Loop:
// 	for x := 0; x < n; x++ {
// 		if column[x] > 0 {
// 			continue
// 		}
// 		for i := 1; x-i >= 0 && y-i >= 0; i++ {
// 			if strs[y-i][x-i] == 1 {
// 				continue Loop
// 			}
// 		}
// 		for i := 1; x+i < n && y-i >= 0; i++ {
// 			if strs[y-i][x+i] == 1 {
// 				continue Loop
// 			}
// 		}
// 		column[x]++
// 		strs[y][x]++
// 		totalNQueens_(n, y+1, column, strs, ret)
// 		column[x]--
// 		strs[y][x]--
// 	}
// }
