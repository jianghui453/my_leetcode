package number

import "testing"

func TestIsNumber(t *testing.T) {
	var s string
	var hope, ret bool

	s = "0"
	hope = true
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = " 0.1 "
	hope = true
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = "abc"
	hope = false
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = "1 a"
	hope = false
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = "2e10"
	hope = true
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = " -90e3   "
	hope = true
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = " 1e"
	hope = false
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = "e3"
	hope = false
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = " 6e-1"
	hope = true
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = " 99e2.5 "
	hope = false
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = "53.5e93"
	hope = true
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = " --6 "
	hope = false
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = "-+3"
	hope = false
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = "95a54e53"
	hope = false
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = ".1"
	hope = true
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = ".0"
	hope = true
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = "01"
	hope = true
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)

	s = "3."
	hope = true
	ret = isNumber(s)
	t.Logf("s=%s hope=%t ret=%t", s, hope, ret)
}
