package dynamic_programming

import (
	"testing"
)

func Test_maximalSquare(t *testing.T) {
	var (
		matrix [][]byte
		hope, ret int
	)

	matrix = [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	hope = 4
	ret = maximalSquare(matrix)
	t.Log(ret == hope, matrix, hope, ret)

	matrix = [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '0', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	hope = 1
	ret = maximalSquare(matrix)
	t.Log(ret == hope, matrix, hope, ret)

	matrix = [][]byte{
		{'0','0','1','0'},
		{'1','1','1','1'},
		{'1','1','1','1'},
		{'1','1','1','0'},
		{'1','1','0','0'},
		{'1','1','1','1'},
		{'1','1','1','0'},
	}
	hope = 9
	ret = maximalSquare(matrix)
	t.Log(ret == hope, matrix, hope, ret)

	matrix = [][]byte{
		{'0','0','0','1'},
		{'1','1','0','1'},
		{'1','1','1','1'},
		{'0','1','1','1'},
		{'0','1','1','1'},
	}
	hope = 9
	ret = maximalSquare(matrix)
	t.Log(ret == hope, matrix, hope, ret)
}