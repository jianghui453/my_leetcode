package dynamic_programming

import (
	"testing"
)

func Test_numSquares(t *testing.T) {
	var n, hope, ret int

	n = 12
	hope = 3
	ret = numSquares(n)
	t.Log(ret == hope, n, hope, ret)

	n = 13
	hope = 2
	ret = numSquares(n)
	t.Log(ret == hope, n, hope, ret)

	n = 0
	hope = 0
	ret = numSquares(n)
	t.Log(ret == hope, n, hope, ret)

	n = 1
	hope = 1
	ret = numSquares(n)
	t.Log(ret == hope, n, hope, ret)
}