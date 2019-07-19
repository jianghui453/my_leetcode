package trap

import "testing"

func TestTrap(t *testing.T) {
	var height []int
	var hope, ret int

	height = []int{0,1,0,2,1,0,1,3,2,1,2,1}
	hope = 6
	ret = trap(height)
	t.Logf("height = %v, hope = %d, ret = %d.\n", height, hope, ret)

	height = []int{0,1,0,2,1,0,1,3,2,1,2, 0, 1}
	hope = 7
	ret = trap(height)
	t.Logf("height = %v, hope = %d, ret = %d.\n", height, hope, ret)
}