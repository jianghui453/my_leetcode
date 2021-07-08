package backtracking

import (
	"testing"
)

func TestSmallestFactorization(t *testing.T) {
	var (
		a, ret, hope int
	)

	a = 48
	hope = 68
	ret = smallestFactorization(a)
	t.Log(ret == hope, "a =", a, "hope =", hope, "ret =", ret)

	a = 15
	hope = 35
	ret = smallestFactorization(a)
	t.Log(ret == hope, "a =", a, "hope =", hope, "ret =", ret)
}