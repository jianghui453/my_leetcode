package backtracking

import (
	"testing"
)

func TestNumberOfPatterns(t *testing.T) {
	var (
		m, n, ret, hope int
	)

	m, n = 1, 1
	hope = 9
	ret = numberOfPatterns(m, n)
	t.Log(ret == hope, m, n, hope, ret)

	m, n = 1, 2
	hope = 65
	ret = numberOfPatterns(m, n)
	t.Log(ret == hope, m, n, hope, ret)

	m, n = 2, 2
	hope = 56
	ret = numberOfPatterns(m, n)
	t.Log(ret == hope, m, n, hope, ret)
}