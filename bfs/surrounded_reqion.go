//给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
//
//找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
//
//示例:
//
//X X X X
//X O O X
//X X O X
//X O X X
//运行你的函数后，矩阵变为：
//
//X X X X
//X X X X
//X X X X
//X O X X
//解释:
//
//被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

package bfs

func solve(board [][]byte) {
	m := len(board)
	if m <= 2 {
		return
	}
	n := len(board[0])
	if n <= 2 {
		return
	}

	matrix := make([][2]int, 0)
	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			board[0][i] = '.'
			matrix = append(matrix, [2]int{0, i})
		}
	}
	for i := 1; i < m; i++ {
		if board[i][n-1] == 'O' {
			board[i][n-1] = '.'
			matrix = append(matrix, [2]int{i, n-1})
		}
	}
	for i := n-2; i >= 0; i-- {
		if board[m-1][i] == 'O' {
			board[m-1][i] = '.'
			matrix = append(matrix, [2]int{m-1, i})
		}
	}
	for i := m-2; i > 0; i-- {
		if board[i][0] == 'O' {
			board[i][0] = '.'
			matrix = append(matrix, [2]int{i, 0})
		}
	}
	for len(matrix) > 0 {
		newMatrix := make([][2]int, 0)
		for i := range matrix {
			x, y := matrix[i][0], matrix[i][1]
			if x > 0 && board[x-1][y] == 'O' {
				board[x-1][y] = '.'
				newMatrix = append(newMatrix, [2]int{x-1, y})
			}
			if y > 0 && board[x][y-1] == 'O' {
				board[x][y-1] = '.'
				newMatrix = append(newMatrix, [2]int{x, y-1})
			}
			if x < m-1 && board[x+1][y] == 'O' {
				board[x+1][y] = '.'
				newMatrix = append(newMatrix, [2]int{x+1, y})
			}
			if y < n-1 && board[x][y+1] == 'O' {
				board[x][y+1] = '.'
				newMatrix = append(newMatrix, [2]int{x, y+1})
			}
		}
		matrix = newMatrix
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

// func solve(board [][]byte) {
// 	height := len(board)
// 	if height < 3 {
// 		return
// 	}
// 	width := len(board[0])
// 	if width < 3 {
// 		return
// 	}

// 	record := make([][]bool, height)
// 	for i := 0; i < height; i++ {
// 		record[i] = make([]bool, width)
// 	}

// 	var his [][]int
// 	for i := 0; i < width; i++ {
// 		if board[0][i] == 'O' {
// 			his = append(his, []int{0, i})
// 			record[0][i] = true
// 		}
// 		if board[height-1][i] == 'O' {
// 			his = append(his, []int{height - 1, i})
// 			record[height-1][i] = true
// 		}
// 	}
// 	for i := 1; i < height-1; i++ {
// 		if board[i][0] == 'O' {
// 			his = append(his, []int{i, 0})
// 			record[i][0] = true
// 		}
// 		if board[i][width-1] == 'O' {
// 			his = append(his, []int{i, width - 1})
// 			record[i][width-1] = true
// 		}
// 	}
// 	for len(his) > 0 {
// 		hisLen := len(his)
// 		newHis := make([][]int, 0)
// 		for i := 0; i < hisLen; i++ {
// 			x := his[i][0]
// 			y := his[i][1]
// 			if x > 0 && board[x-1][y] == 'O' && !record[x-1][y] {
// 				newHis = append(newHis, []int{x - 1, y})
// 				record[x-1][y] = true
// 			}
// 			if x < height-1 && board[x+1][y] == 'O' && !record[x+1][y] {
// 				newHis = append(newHis, []int{x + 1, y})
// 				record[x+1][y] = true
// 			}
// 			if y > 0 && board[x][y-1] == 'O' && !record[x][y-1] {
// 				newHis = append(newHis, []int{x, y - 1})
// 				record[x][y-1] = true
// 			}
// 			if y < width-1 && board[x][y+1] == 'O' && !record[x][y+1] {
// 				newHis = append(newHis, []int{x, y + 1})
// 				record[x][y+1] = true
// 			}
// 		}
// 		his = newHis
// 	}
// 	for i := 0; i < height; i++ {
// 		for j := 0; j < width; j++ {
// 			if board[i][j] == 'O' && !record[i][j] {
// 				board[i][j] = 'X'
// 			}
// 		}
// 	}
// }
