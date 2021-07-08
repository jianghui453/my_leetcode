package string

import "testing"

func TestCountAndSay(t *testing.T) {
	var n int
	var h, r string

	n = 1
	h = "1"
	r = countAndSay(n)
	t.Logf("n=%d h=%s r=%s", n, h, r)

	n = 2
	h = "11"
	r = countAndSay(n)
	t.Logf("n=%d h=%s r=%s", n, h, r)

	n = 3
	h = "21"
	r = countAndSay(n)
	t.Logf("n=%d h=%s r=%s", n, h, r)

	n = 4
	h = "1211"
	r = countAndSay(n)
	t.Logf("n=%d h=%s r=%s", n, h, r)

	n = 5
	h = "111221"
	r = countAndSay(n)
	t.Logf("n=%d h=%s r=%s", n, h, r)
}
