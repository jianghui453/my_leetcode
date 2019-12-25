package array

import (
	"testing"
)

func TestGetRow(t *testing.T) {
	var rowIndex int
	var hope, ret []int

	rowIndex = 4
	ret = getRow(rowIndex)
	hope = []int{1, 4, 6, 4, 1}
	t.Logf("\nhope=%v \n ret=%v", hope, ret)

	rowIndex = 0
	ret = getRow(rowIndex)
	hope = []int{1}
	t.Logf("\nhope=%v \n ret=%v", hope, ret)
}
