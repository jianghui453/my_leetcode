package array

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	var numRows int
	var hope, ret [][]int

	numRows = 5
	ret = generate(numRows)
	hope = [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}
	t.Logf("\nhope=%v \n ret=%v", hope, ret)

	numRows = 1
	ret = generate(numRows)
	hope = [][]int{
		{1},
	}
	t.Logf("\nhope=%v \n ret=%v", hope, ret)
}
