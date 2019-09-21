package bfs

import "testing"

func TestMinCut(t *testing.T) {
	var s string
	var h, r int

	s = "aab"
	h = 1
	r = minCut(s)
	t.Logf("s=%s h=%d r=%d", s, h, r)

	s = "abc"
	h = 2
	r = minCut(s)
	t.Logf("s=%s h=%d r=%d", s, h, r)

	s = "aaa"
	h = 0
	r = minCut(s)
	t.Logf("s=%s h=%d r=%d", s, h, r)

	s = "a"
	h = 0
	r = minCut(s)
	t.Logf("s=%s h=%d r=%d", s, h, r)

	s = "abbab"
	h = 1
	r = minCut(s)
	t.Logf("s=%s h=%d r=%d", s, h, r)

	s = "fifgbeajcacehiicccfecbfhhgfiiecdcjjffbghdidbhbdbfbfjccgbbdcjheccfbhafehieabbdfeigbiaggchaeghaijfbjhi"
	h = 75
	r = minCut(s)
	t.Logf("s=%s h=%d r=%d", s, h, r)
}
