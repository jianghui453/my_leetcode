package unique_path

import "testing"

func TestMinPathSum(t *testing.T) {
	var grid [][]int
	var hope, ret int

	grid = [][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}
	hope = 7
	ret = minPathSum(grid)
	t.Logf("hope=%d ret=%d", hope, ret)
}