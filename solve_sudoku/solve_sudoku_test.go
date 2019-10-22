package solve_sudoku

import "testing"

func TestSolveSudoku(t *testing.T) {
	var abord [][]byte

	abord = [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	ret := [][]string{}
	for _, bytes := range abord {
		vs := []string{}
		for _, v := range bytes {
			vs = append(vs, string(v))
		}
		ret = append(ret, vs)
	}
	t.Logf("before = %v", ret)
	solveSudoku(abord)
	ret = [][]string{}
	for _, bytes := range abord {
		vs := []string{}
		for _, v := range bytes {
			vs = append(vs, string(v))
		}
		ret = append(ret, vs)
	}
	t.Logf("after = %v", ret)

	// t.Logf("%d", -9 / 8)
}
