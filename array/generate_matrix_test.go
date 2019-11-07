package array

import "testing"

func TestGeneraMatrix(t *testing.T) {
	var n int
	var ret [][]int

	n = 3
	ret = generateMatrix(n)
	t.Logf("\nn=%d \nret=%v", n, ret)
}
