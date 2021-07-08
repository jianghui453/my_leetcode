package dijkstra

import "testing"

func TestDijkstra(t *testing.T) {
	var graph [][]int
	var hope, ret []int

	graph = [][]int{
		{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0},
	}
	hope = []int{0, 4, 12, 19, 21, 11, 9, 8, 14}
	ret = dijkstra(9, graph)
	t.Logf("graph=%v \nhope=%v \n ret=%v", graph, hope, ret)
}
