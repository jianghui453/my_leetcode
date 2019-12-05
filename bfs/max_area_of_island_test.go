package bfs

import (
	"testing"
)

func Test_maxAreaOfIsland(t *testing.T) {
	var grid [][]int
	var hope, ret int

	grid = [][]int{
		{0,0,1,0,0,0,0,1,0,0,0,0,0},
 		{0,0,0,0,0,0,0,1,1,1,0,0,0},
 		{0,1,1,0,1,0,0,0,0,0,0,0,0},
		{0,1,0,0,1,1,0,0,1,0,1,0,0},
		{0,1,0,0,1,1,0,0,1,1,1,0,0},
		{0,0,0,0,0,0,0,0,0,0,1,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,0,0,0,0,0,0,1,1,0,0,0,0},
	}
	hope, ret = 6, maxAreaOfIsland(grid)
	t.Logf("%t grid=%v hope=%d ret=%d", ret == hope, grid, hope, ret)
}