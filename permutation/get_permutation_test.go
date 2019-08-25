package permutation

import "testing"

func TestGetPermutation(t *testing.T) {
	var str string
	//str = getPermutation(3, 3)
	//t.Log(str)

	str = getPermutation(4, 9)
	t.Log(str)
}