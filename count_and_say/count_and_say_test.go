package count_and_say

import "testing"

func TestCountAndSay(t *testing.T) {
	var n int
	var hope, ret string

	n = 1
	hope = "1"
	ret = countAndSay(n)
	t.Logf("n = %d, hope = %s, ret = %s", n, hope, ret)

	n = 2
	hope = "11"
	ret = countAndSay(n)
	t.Logf("n = %d, hope = %s, ret = %s", n, hope, ret)

	n = 3
	hope = "21"
	ret = countAndSay(n)
	t.Logf("n = %d, hope = %s, ret = %s", n, hope, ret)

	n = 4
	hope = "1211"
	ret = countAndSay(n)
	t.Logf("n = %d, hope = %s, ret = %s", n, hope, ret)

	n = 5
	hope = "111221"
	ret = countAndSay(n)
	t.Logf("n = %d, hope = %s, ret = %s", n, hope, ret)
}
