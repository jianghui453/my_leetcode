package fsa

import (
	"testing"
)

func TestIsNumber(t *testing.T) {
	var s string
	var h, r bool

	s = "0"
	h = true
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = " 0.1 "
	h = true
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = "abc"
	h = false
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = "1 a"
	h = false
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = "2e10"
	h = true
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = " -90e3   "
	h = true
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = " 1e"
	h = false
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = "e3"
	h = false
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = " 6e-1"
	h = true
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = " 99e2.5 "
	h = false
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = "53.5e93"
	h = true
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = " --6 "
	h = false
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = "-+3"
	h = false
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = "95a54e53"
	h = false
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = ".1"
	h = true
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = ".0"
	h = true
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = "01"
	h = true
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = "3."
	h = true
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = "."
	h = false
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = "4e+"
	h = false
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = " -."
	h = false
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = " .1. "
	h = false
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = " 6.3.0 "
	h = false
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)

	s = " 46.e3 "
	h = true
	r = isNumber(s)
	t.Logf("%t s=%s h=%t r=%t", r == h, s, h, r)
}
