//给定一个二维网格和一个单词，找出该单词是否存在于网格中。
//
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//示例:
//
//board =
//[
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
//]
//
//给定 word = "ABCCED", 返回 true.
//给定 word = "SEE", 返回 true.
//给定 word = "ABCB", 返回 false.

package array

func exist(board [][]byte, word string) bool {
	l := len(word)
	if l == 0 {
		return true
	}
	
	m := len(board)
	if m == 0 {
		return false
	}

	n := len(board[0])
	if n == 0 {
		return false
	}

	var f func(i, x, y int) bool
	f = func(i, x, y int) bool {
		if i == l {
			return true
		}

		if x > 0 {
			if board[x-1][y] != ' ' && board[x-1][y] == word[i] {
				board[x-1][y] = ' '
				if f(i+1, x-1, y) {
					return true
				}
				board[x-1][y] = word[i] 
			}
		}
		if x < m-1 {
			if board[x+1][y] != ' ' && board[x+1][y] == word[i] {
				board[x+1][y] = ' '
				if f(i+1, x+1, y) {
					return true
				}
				board[x+1][y] = word[i]
			}
		}
		if y > 0 {
			if board[x][y-1] != ' ' && board[x][y-1] == word[i] {
				board[x][y-1] = ' '
				if f(i+1, x, y-1) {
					return true
				}
				board[x][y-1] = word[i]
			}
		}
		if y < n-1 {
			if board[x][y+1] != ' ' && board[x][y+1] == word[i] {
				board[x][y+1] = ' '
				if f(i+1, x, y+1) {
					return true
				}
				board[x][y+1] = word[i]
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				board[i][j] = ' '
				if f(1, i, j) {
					return true
				}
				board[i][j] = word[0]
			}
		}
	}

	return false
}

// func exist(board [][]byte, word string) bool {
// 	lenI := len(board)
// 	if lenI < 1 {
// 		return false
// 	}
// 	lenJ := len(board[0])
// 	if lenJ < 1 {
// 		return false
// 	}
// 	for i := 0; i < lenI; i++ {
// 		for j := 0; j < lenJ; j++ {
// 			if board[i][j] == word[0] {
// 				if len(word) < 2 {
// 					return true
// 				}
// 				boardIJ := board[i][j]
// 				board[i][j] = '0'
// 				if searchByte(board, i, j, word[1:]) {
// 					return true
// 				}
// 				board[i][j] = boardIJ
// 			}
// 		}
// 	}
// 	return false
// }

// func searchByte(board [][]byte, i, j int, word string) bool {
// 	if i > 0 && board[i-1][j] == word[0] {
// 		if len(word) < 2 {
// 			return true
// 		}
// 		boardIJ := board[i-1][j]
// 		board[i-1][j] = '0'
// 		if searchByte(board, i-1, j, word[1:]) {
// 			return true
// 		}
// 		board[i-1][j] = boardIJ
// 	}
// 	if i < len(board)-1 && board[i+1][j] == word[0] {
// 		if len(word) < 2 {
// 			return true
// 		}
// 		boardIJ := board[i+1][j]
// 		board[i+1][j] = '0'
// 		if searchByte(board, i+1, j, word[1:]) {
// 			return true
// 		}
// 		board[i+1][j] = boardIJ
// 	}
// 	if j > 0 && board[i][j-1] == word[0] {
// 		if len(word) < 2 {
// 			return true
// 		}
// 		boardIJ := board[i][j-1]
// 		board[i][j-1] = '0'
// 		if searchByte(board, i, j-1, word[1:]) {
// 			return true
// 		}
// 		board[i][j-1] = boardIJ
// 	}
// 	if j < len(board[0])-1 && board[i][j+1] == word[0] {
// 		if len(word) < 2 {
// 			return true
// 		}
// 		boardIJ := board[i][j+1]
// 		board[i][j+1] = '0'
// 		if searchByte(board, i, j+1, word[1:]) {
// 			return true
// 		}
// 		board[i][j+1] = boardIJ
// 	}
// 	return false
// }
