package my_pow

import "testing"

func TestMyPow(t *testing.T) {
    var x, hope, ret float64
    var n int

    x = 2.0
    n = -2
    hope = 0.25
    ret = myPow(x, n)
    t.Logf("x=%g n=%d hope=%g ret=%g", x, n, hope, ret)
}