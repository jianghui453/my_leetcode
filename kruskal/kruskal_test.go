package kruskal

import "testing"

func TestKruskal(t *testing.T) {
    g := Graph{
        4,
        5,
        []edge{
            {0, 1,10},
            {0, 2,6},
            {0, 3,5},
            {1, 3,15},
            {2, 3,4},
        },
    }
    mst := kruskal(g)
    t.Logf("\ng=%v \nmst=%v", g, mst)
}