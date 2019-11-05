package array

import (
	"testing"
)

func TestRotate(t *testing.T) {
	var matrix, h [][]int

	matrix = [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	h = [][]int{
		{7,4,1},
		{8,5,2},
		{9,6,3},
	}
	rotate(matrix)
	t.Logf("\nmatrix=%v \n     h=%v", matrix, h)
}