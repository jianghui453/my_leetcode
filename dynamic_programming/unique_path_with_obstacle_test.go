package dynamic_programming

import "testing"

func TestUniquePathWithObstacle(t *testing.T) {
	var obstacleGrid [][]int
	var h, r int

	obstacleGrid = [][]int{
		{0,0,0},
		{0,1,0},
		{0,0,0},
	}
	h = 2
	r = uniquePathsWithObstacles(obstacleGrid)
	t.Logf("h=%d r=%d", h, r)
}