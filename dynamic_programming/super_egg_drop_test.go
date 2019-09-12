package dynamic_programming

import "testing"

func TestSuperEggDrop(t *testing.T) {
	var K, N, hope, ret int

	K = 2
	N = 6
	hope = 3
	ret = superEggDrop(K, N)
	t.Logf("K=%d N=%d hope=%d ret=%d", K, N, hope, ret)

	K = 3
	N = 14
	hope = 4
	ret = superEggDrop(K, N)
	t.Logf("K=%d N=%d hope=%d ret=%d", K, N, hope, ret)
}
