package string

import (
	"testing"
)

func TestMyAtoI(t *testing.T) {
	var str string
	var r, h int

	str = "4193 with words"
	h = 4193
	r = myAtoi(str)
	t.Logf("%t str=%s h=%d r=%d", h == r, str, h, r)

	str = "42"
	h = 42
	r = myAtoi(str)
	t.Logf("%t str=%s h=%d r=%d", h == r, str, h, r)

	str = " -42"
	h = -42
	r = myAtoi(str)
	t.Logf("%t str=%s h=%d r=%d", h == r, str, h, r)

	str = "words and 987"
	h = 0
	r = myAtoi(str)
	t.Logf("%t str=%s h=%d r=%d", h == r, str, h, r)

	str = "-91283472332"
	h = -2147483648
	r = myAtoi(str)
	t.Logf("%t str=%s h=%d r=%d", h == r, str, h, r)

	str = "+-2"
	h = 0
	r = myAtoi(str)
	t.Logf("%t str=%s h=%d r=%d", h == r, str, h, r)

	str = "-+2"
	h = 0
	r = myAtoi(str)
	t.Logf("%t str=%s h=%d r=%d", h == r, str, h, r)

	str = "  +0 123"
	h = 0
	r = myAtoi(str)
	t.Logf("%t str=%s h=%d r=%d", h == r, str, h, r)

	str = "9223372036854775808"
	h = 2147483647
	r = myAtoi(str)
	t.Logf("%t str=%s h=%d r=%d", h == r, str, h, r)
}
