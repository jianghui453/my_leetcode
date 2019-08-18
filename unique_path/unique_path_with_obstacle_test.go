package unique_path

import "testing"

func TestUniquePathWithObstacle(t *testing.T) {
	var obstacleGrid [][]int
	var hope, ret int

	obstacleGrid = [][]int{
		{0,0,0},
		{0,1,0},
		{0,0,0},
	}
	hope = 2
	ret = uniquePathsWithObstacles(obstacleGrid)
	t.Logf("hope=%d ret=%d", hope, ret)
}