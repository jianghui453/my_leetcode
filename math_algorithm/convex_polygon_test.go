package math_algorithm

import (
	"testing"
)

func TestIsConvex(t *testing.T) {
	var (
		points [][]int
		ret, hope bool
	)

	points = [][]int{{0,0},{0,1},{1,1},{1,0}}
	hope = true
	ret = isConvex(points)
	t.Log(ret == hope, "points =", points, "hope =", hope, "ret =", ret)

	points = [][]int{{0,0},{0,10},{10,10},{10,0},{5,5}}
	hope = false
	ret = isConvex(points)
	t.Log(ret == hope, "points =", points, "hope =", hope, "ret =", ret)

	points = [][]int{{0,0},{1, 0},{1,1},{0,1}}
	hope = true
	ret = isConvex(points)
	t.Log(ret == hope, "points =", points, "hope =", hope, "ret =", ret)

	points = [][]int{{0,0},{1,1},{0,2},{-1,2},{-1,1}}
	hope = true
	ret = isConvex(points)
	t.Log(ret == hope, "points =", points, "hope =", hope, "ret =", ret)
}