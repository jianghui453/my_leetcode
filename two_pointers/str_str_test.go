package two_pointers

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	var haystack, needle string
	var h, r int

	haystack = "a"
	needle = "a"
	r = strStr(haystack, needle)
	t.Logf("%t haystack=%s needle=%s h=%d r=%d", r == h, haystack, needle, h, r)
}
