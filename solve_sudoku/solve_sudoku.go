package solve_sudoku

import (
	"fmt"
)

func solveSudoku(board [][]byte) {
	boardMirror := make([][]byte, 9)
	vertical := make([][]int, 9)
	horizontal := make([][]int, 9)
	square := make([][]int, 9)

	for i := 0; i < 9; i++ {
		boardMirror[i] = make([]byte, 9)
		vertical[i] = make([]int, 9)
		horizontal[i] = make([]int, 9)
		square[i] = make([]int, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			boardMirror[i][j] = board[i][j]
			if board[i][j] != '.' {
				v := int(board[i][j]-'1')

				vertical[j][v]++
				horizontal[i][v]++
				idx := i/3*3+j/3
				square[idx][v]++
			}
		}
	}

	var f func(i, j, v int) bool
	f = func(i, j, v int) bool {
		fmt.Printf("i=%d j=%d v=%d\n", i, j, v)
		idx := i/3*3 + j/3
		if boardMirror[i][j] == '.' && (vertical[j][v] > 0 || horizontal[i][v] > 0 || square[idx][v] > 0) {
			return false
		}

		if boardMirror[i][j] == '.' {
			board[i][j] = byte('1'+v)
		}

		if i == 8 && j == 8 {
			return true
		}

		vertical[j][v]++
		horizontal[i][v]++
		square[idx][v]++
		
		if j >= 8 {
			i++
			j = 0
		} else {
			j++
		}

		var ret bool
		_v := 0
		for !ret && _v < 9 {
			ret = f(i, j, _v)
			_v++
		}

		if _v == 9 {
			vertical[j][v]--
			horizontal[i][v]--
			square[idx][v]--
		}
		
		return ret
	}

	val := 0
	var r bool
	for !r && val < 9 {
		r = f(0, 0, val)
		val++
	}
}

// func solveSudoku(board [][]byte) {
// 	var record, row, column, square [9][9]int
// 	for y := 0; y < 9; y++ {
// 		for x := 0; x < 9; x++ {
// 			if board[y][x] != '.' {
// 				record[y][x]++
// 				row[y][int(board[y][x]-'1')]++
// 				column[x][int(board[y][x]-'1')]++
// 				square[y/3*3+x/3][int(board[y][x]-'1')]++
// 			}
// 		}
// 	}
// 	for y := 0; y < 9; y++ {
// 	Loop:
// 		for x := 0; x < 9; x++ {
// 			squareIndex := y/3*3 + x/3
// 			if board[y][x] != '.' {
// 				continue
// 			}
// 			for i := 0; i < 9; i++ {
// 				if row[y][i]+column[x][i]+square[squareIndex][i] == 0 {
// 					row[y][i]++
// 					column[x][i]++
// 					square[squareIndex][i]++
// 					board[y][x] = byte(i + '1')
// 					continue Loop
// 				}
// 			}
// 			x0, y0 := (x+8)%9, y+(x-8)/8
// 			for ; x0 >= 0 && y0 >= 0; x0, y0 = (x0+8)%9, y0+(x0-8)/8 {
// 				squareIndex0 := y0/3*3 + x0/3
// 				val := int(board[y0][x0] - '1')
// 				if record[y0][x0] > 0 {
// 					continue
// 				}
// 				row[y0][val]--
// 				column[x0][val]--
// 				square[squareIndex0][val]--
// 				for i := int(board[y0][x0]-'1') + 1; i < 9; i++ {
// 					if row[y0][i]+column[x0][i]+square[squareIndex0][i] == 0 {
// 						row[y0][i]++
// 						column[x0][i]++
// 						square[squareIndex0][i]++
// 						board[y0][x0] = byte(i + '1')
// 						x, y = x0, y0
// 						continue Loop
// 					}
// 				}
// 				board[y0][x0] = '.'
// 			}
// 		}
// 	}
// }

// func solveSudokuV1(board [][]byte) {
// 	fmt.Println("Start.")
// 	var row, column, square [9][9]int
// 	solve(board, 0, 0, row, column, square)
// }

// func solve(board [][]byte, x, y int, row, column, square [9][9]int) bool {
// 	if x > 8 || y > 8 {
// 		return true
// 	}
// 	x0 := (x + 1) % 9
// 	y0 := y + (x+1)/9
// 	squareIndex := y/3*3 + x/3
// 	// fmt.Printf("board = %v, x = %d, y = %d, row = %v, column = %v, square = %v\n", board, x, y, row, column, square)
// 	if board[y][x] == '.' {
// 		for i := 0; i < 9; i++ {
// 			if row[y][i]+column[x][i]+square[squareIndex][i] > 0 {
// 				continue
// 			}
// 			// fmt.Printf("x = %d; y = %d; i = %d; x0 = %d; y0 = %d; row = %v; column = %v; square = %v, board = %v.\n", x, y, i, x0, y0, row, column, square, board)
// 			row[y][i]++
// 			column[x][i]++
// 			square[squareIndex][i]++
// 			board[y][x] = byte(i + '1')
// 			// fmt.Printf("board[y][x] = %d.\n", board[y][x])
// 			if solve(board, x0, y0, row, column, square) {
// 				return true
// 			} else {
// 				row[y][i]--
// 				column[x][i]--
// 				square[squareIndex][i]--
// 				board[y][x] = '.'
// 			}
// 		}
// 		return false
// 	} else {
// 		val := int(board[y][x] - '1')
// 		if row[y][val]+column[x][val]+square[squareIndex][val] > 0 {
// 			return false
// 		}
// 		column[x][val]++
// 		row[y][val]++
// 		square[squareIndex][val]++
// 		if !solve(board, x0, y0, row, column, square) {
// 			column[x][val]--
// 			row[y][val]--
// 			square[squareIndex][val]--
// 			return false
// 		} else {
// 			return true
// 		}
// 	}
// }
