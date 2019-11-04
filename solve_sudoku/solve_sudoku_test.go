package solve_sudoku

import "testing"

func TestSolveSudoku(t *testing.T) {
	var board [][]byte
	var h, r [][]int

	var toIntSlice = func(b [][]byte) [][]int {
		ret := make([][]int, 9)
		for i := 0; i < 9; i++ {
			ret[i] = make([]int, 9)

			for j := 0; j < 9; j++ {
				ret[i][j] = int(b[i][j]-'0')
			}
		}

		return ret
	}

	board = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	r = toIntSlice(board)
	h = [][]int{
		{},
	}
	t.Logf("\nr=%v \nh=%v", r, h)
}
