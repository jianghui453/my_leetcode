package array

import "testing"

func TestSetZeroes(t *testing.T) {
	var martix, hope [][]int

	martix = [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	hope = [][]int{{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1}}
	setZeroes(martix)
	t.Logf("\nmatrix=%v \n  hope=%v", martix, hope)

	martix = [][]int{
		{1, 0},
	}
	hope = [][]int{{0, 0}}
	setZeroes(martix)
	t.Logf("\nmatrix=%v \n  hope=%v", martix, hope)
}
