package array

import "testing"

func TestExist(t *testing.T) {
	var board [][]byte
	var word string
	var hope, ret bool

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word = "ABCCED"
	hope = true
	ret = exist(board, word)
	t.Logf("hope=%t ret=%t", hope, ret)

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word = "SEE"
	hope = true
	ret = exist(board, word)
	t.Logf("hope=%t ret=%t", hope, ret)

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word = "ABCB"
	hope = false
	ret = exist(board, word)
	t.Logf("hope=%t ret=%t", hope, ret)
}
