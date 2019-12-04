package bfs

import (
	"testing"
)

func Test_numIslands(t *testing.T) {
	var grid [][]byte
	var hope, ret int

	grid = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'}, 
		{'1', '1', '0', '0', '0'}, 
		{'0', '0', '0', '0', '0'}, 
	}
	hope, ret = 1, numIslands(grid)
	t.Logf("%t grid=%v hope=%d ret=%d", ret == hope, grid, hope, ret)

	grid = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'}, 
		{'0', '0', '1', '0', '0'}, 
		{'0', '0', '0', '1', '1'}, 
	}
	hope, ret = 3, numIslands(grid)
	t.Logf("%t grid=%v hope=%d ret=%d", ret == hope, grid, hope, ret)
}