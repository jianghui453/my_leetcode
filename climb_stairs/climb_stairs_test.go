package climb_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
    var n, hope, ret int

    n = 2
    hope = 2
    ret = climbStairs(n)
    t.Logf("n=%d hope=%d ret=%d", n, hope, ret)

    n = 3
    hope = 3
    ret = climbStairs(n)
    t.Logf("n=%d hope=%d ret=%d", n, hope, ret)

    n = 5
    hope = 8
    ret = climbStairs(n)
    t.Logf("n=%d hope=%d ret=%d", n, hope, ret)
}