package graph

import (
	"testing"
)

func Test_kruskal(t *testing.T) {
	var V int
	var g, hope, ret Graph

	V = 4
	g = Graph{
		4,
		5,
		[]Edge{
			Edge{0, 1, 10},
			Edge{0, 2, 6},
			Edge{0, 3, 5},
			Edge{1, 3, 15},
			Edge{2, 3, 4},
		},
	}
	hope = Graph{
		4,
		3,
		[]Edge{
			Edge{2, 3, 4},
			Edge{0, 3, 5},
			Edge{0, 1, 10},
		},
	}
	ret = kruskal(g)
	t.Logf("V=%d g=%v \nhope=%v \n ret=%v", V, g, hope, ret)
}
