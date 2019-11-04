package stack

import "testing"

func TestTrap(t *testing.T) {
	var height []int
	var h, r int

	height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	h = 6
	r = trap(height)
	t.Logf("height = %v, h = %d, r = %d.\n", height, h, r)

	height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 0, 1}
	h = 7
	r = trap(height)
	t.Logf("height = %v, h = %d, r = %d.\n", height, h, r)

	height = []int{0, 1, 2}
	h = 0
	r = trap(height)
	t.Logf("height = %v, h = %d, r = %d.\n", height, h, r)

	height = []int{2, 1, 0}
	h = 0
	r = trap(height)
	t.Logf("height = %v, h = %d, r = %d.\n", height, h, r)

	height = []int{2, 1, 0, 1}
	h = 1
	r = trap(height)
	t.Logf("height = %v, h = %d, r = %d.\n", height, h, r)

	height = []int{2, 1, 0, 2}
	h = 3
	r = trap(height)
	t.Logf("height = %v, h = %d, r = %d.\n", height, h, r)

	height = []int{6,4,2,0,3,2,0,3,1,4,5,3,2,7,5,3,0,1,2,1,3,4,6,8,1,3}
	h = 83
	r = trap(height)
	t.Logf("height = %v, h = %d, r = %d.\n", height, h, r)
}
