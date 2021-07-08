package string

import (
	"testing"
)

func TestNearestPalindromic(t *testing.T) {
	var (
		n, ret, hope string
	)

	n = "123"
	hope = "121"
	ret = nearestPalindromic(n)
	t.Log(ret == hope, "n =", n, "hope =", hope, "ret =", ret)

	n = "100"
	hope = "99"
	ret = nearestPalindromic(n)
	t.Log(ret == hope, "n =", n, "hope =", hope, "ret =", ret)

	n = "101"
	hope = "99"
	ret = nearestPalindromic(n)
	t.Log(ret == hope, "n =", n, "hope =", hope, "ret =", ret)

	n = "111"
	hope = "101"
	ret = nearestPalindromic(n)
	t.Log(ret == hope, "n =", n, "hope =", hope, "ret =", ret)

	n = "2"
	hope = "1"
	ret = nearestPalindromic(n)
	t.Log(ret == hope, "n =", n, "hope =", hope, "ret =", ret)

	n = "200"
	hope = "202"
	ret = nearestPalindromic(n)
	t.Log(ret == hope, "n =", n, "hope =", hope, "ret =", ret)

	n = "213"
	hope = "212"
	ret = nearestPalindromic(n)
	t.Log(ret == hope, "n =", n, "hope =", hope, "ret =", ret)

	n = "1213"
	hope = "1221"
	ret = nearestPalindromic(n)
	t.Log(ret == hope, "n =", n, "hope =", hope, "ret =", ret)

	n = "99"
	hope = "101"
	ret = nearestPalindromic(n)
	t.Log(ret == hope, "n =", n, "hope =", hope, "ret =", ret)

	n = "1283"
	hope = "1331"
	ret = nearestPalindromic(n)
	t.Log(ret == hope, "n =", n, "hope =", hope, "ret =", ret)
}
