package math_algorithm

import (
	"testing"
)

func TestIsPalindrome2(t *testing.T) {
	var x int
	var r, h bool

	x = 121
	h = true
	r = isPalindrome2(x)
	t.Logf("%t x=%d r=%t h=%t", r == h, x, r, h)

	x = -121
	h = false
	r = isPalindrome2(x)
	t.Logf("%t x=%d r=%t h=%t", r == h, x, r, h)

	x = 10
	h = false
	r = isPalindrome2(x)
	t.Logf("%t x=%d r=%t h=%t", r == h, x, r, h)

	x = 0
	h = true
	r = isPalindrome2(x)
	t.Logf("%t x=%d r=%t h=%t", r == h, x, r, h)

	x = 1
	h = true
	r = isPalindrome2(x)
	t.Logf("%t x=%d r=%t h=%t", r == h, x, r, h)
}
