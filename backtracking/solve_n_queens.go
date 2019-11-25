// The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

// Given an integer n, return all distinct solutions to the n-queens puzzle.

// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

// Example:

// Input: 4
// Output: [
//  [".Q..",  // Solution 1
//   "...Q",
//   "Q...",
//   "..Q."],

//  ["..Q.",  // Solution 2
//   "Q...",
//   "...Q",
//   ".Q.."]
// ]
// Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.

package backtracking

func solveNQueens(n int) [][]string {
	ret := make([][]string, 0)
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
			_strs := make([]string, n)
			for i := 0; i < n; i++ {
				_strs[i] = string(strs[i])
			}
			ret = append(ret, _strs)
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

// func solveNQueens(n int) [][]string {
// 	var ret [][]string
// 	if n > 0 {
// 		var b []byte
// 		for i := 0; i < n; i++ {
// 			b = append(b, '.')
// 		}
// 		var strs [][]byte
// 		column := make([]int, n)
// 		for i := 0; i < n; i++ {
// 			b_ := make([]byte, n)
// 			copy(b_, b)
// 			strs = append(strs, b_)
// 		}
// 		solve(0, n, column, &ret, strs)
// 	}
// 	return ret
// }

// func solve(y, n int, column []int, ret *[][]string, strs [][]byte) {
// 	//fmt.Printf("y=%d colume=%v strs=%v\n", y, column, strs)
// 	if y >= n {
// 		ret_ := make([]string, n)
// 		for i, v := range strs {
// 			ret_[i] = string(v)
// 		}
// 		*ret = append(*ret, ret_)
// 		return
// 	}
// Loop:
// 	for x := 0; x < n; x++ {
// 		if column[x] > 0 {
// 			continue
// 		}
// 		i := 1
// 		for ; x-i >= 0 && y-i >= 0; i++ {
// 			if strs[y-i][x-i] == 'Q' {
// 				continue Loop
// 			}
// 		}
// 		i = 1
// 		for ; x+i < n && y-i >= 0; i++ {
// 			if strs[y-i][x+i] == 'Q' {
// 				continue Loop
// 			}
// 		}
// 		column[x]++
// 		strs[y][x] = 'Q'
// 		solve(y+1, n, column, ret, strs)
// 		column[x]--
// 		strs[y][x] = '.'
// 	}
// 	//fmt.Println()
// }
