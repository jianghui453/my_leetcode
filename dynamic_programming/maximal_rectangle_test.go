package dynamix_programming

import "testing"

func TestMaximalRectangle(t *testing.T) {
	var matrix [][]byte
	var hope, ret int

	matrix = [][]byte{[]byte("10100"), []byte("10111"), []byte("11111"), []byte("10010")}
	hope = 6
	ret = maximalRectangle(matrix)
	t.Logf("matrix=%v hope=%d ret=%d", matrix, hope, ret)
}
