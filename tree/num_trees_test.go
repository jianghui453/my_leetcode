package tree

import "testing"

func TestNumTrees(t *testing.T) {
    var n, h, r int

    n = 3
    h = 5
    r = numTrees(n)
    t.Logf("%t n=%d h=%d r=%d", h==r, n, h, r)
}
