package dynamic_programming

import (
	"testing"
)

func TestVideoStitching(t *testing.T) {
	var (
		clips [][]int
		T int
		hope, ret int
	)

	clips = [][]int{{0,2},{4,6},{8,10},{1,9},{1,5},{5,9}}
	T = 10
	hope = 3
	ret = videoStitching(clips, T)
	t.Log(ret == hope, "clips =", clips, "hope =", hope, "ret =", ret)

	clips = [][]int{{0,1},{1,2}}
	T = 5
	hope = -1
	ret = videoStitching(clips, T)
	t.Log(ret == hope, "clips =", clips, "hope =", hope, "ret =", ret)

	clips = [][]int{{0,1},{6,8},{0,2},{5,6},{0,4},{0,3},{6,7},{1,3},{4,7},{1,4},{2,5},{2,6},{3,4},{4,5},{5,7},{6,9}}
	T = 9
	hope = 3
	ret = videoStitching(clips, T)
	t.Log(ret == hope, "clips =", clips, "hope =", hope, "ret =", ret)

	clips = [][]int{{0,4},{2,8}}
	T = 5
	hope = 2
	ret = videoStitching(clips, T)
	t.Log(ret == hope, "clips =", clips, "hope =", hope, "ret =", ret)

	clips = [][]int{{1,4},{2,8}}
	T = 5
	hope = -1k
	ret = videoStitching(clips, T)
	t.Log(ret == hope, "clips =", clips, "hope =", hope, "ret =", ret)

	clips = [][]int{{0,5},{6,8}}
	T = 7
	hope = -1
	ret = videoStitching(clips, T)
	t.Log(ret == hope, "clips =", clips, "hope =", hope, "ret =", ret)
}
