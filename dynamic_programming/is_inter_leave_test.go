package dynamic_programming

import "testing"

func TestIsInterLeave(t *testing.T) {
	var s1, s2, s3 string
	var h, r bool

	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbcbcac"
	h = true
	r = isInterLeave(s1, s2, s3)
	t.Logf("%t s1=%s s2=%s s3=%s h=%t r=%t", h == r, s1, s2, s3, h, r)

	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbbaccc"
	h = false
	r = isInterLeave(s1, s2, s3)
	t.Logf("%t s1=%s s2=%s s3=%s h=%t r=%t", h == r, s1, s2, s3, h, r)
}
