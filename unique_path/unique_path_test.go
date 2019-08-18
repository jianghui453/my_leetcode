package unique_path

import "testing"

func TestUniquePath(t *testing.T) {
	var m, n, hope, ret int

	m = 3
	n = 2
	hope = 3
	ret = uniquePaths(m, n)
	t.Logf("m=%d n=%d hope=%d ret=%d", m, n, hope, ret)

	m = 7
	n = 3
	hope = 28
	ret = uniquePaths(m, n)
	t.Logf("m=%d n=%d hope=%d ret=%d", m, n, hope, ret)
}