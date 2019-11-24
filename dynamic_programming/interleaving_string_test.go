package dynamic_programming

import "testing"

func TestIsInterleave(t *testing.T) {
	var s1, s2, s3 string
	var h, r bool

	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbcbcac"
	h = true
	r = isInterleave(s1, s2, s3)
	t.Logf("%t s1=%s s2=%s s3=%s h=%t r=%t", h == r, s1, s2, s3, h, r)

	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbbaccc"
	h = false
	r = isInterleave(s1, s2, s3)
	t.Logf("%t s1=%s s2=%s s3=%s h=%t r=%t", h == r, s1, s2, s3, h, r)

	s1 = "a"
	s2 = ""
	s3 = "a"
	h = true
	r = isInterleave(s1, s2, s3)
	t.Logf("%t s1=%s s2=%s s3=%s h=%t r=%t", h == r, s1, s2, s3, h, r)

	s1 = "a"
	s2 = ""
	s3 = "b"
	h = false
	r = isInterleave(s1, s2, s3)
	t.Logf("%t s1=%s s2=%s s3=%s h=%t r=%t", h == r, s1, s2, s3, h, r)

	s1 = "a"
	s2 = "b"
	s3 = "ab"
	h = true
	r = isInterleave(s1, s2, s3)
	t.Logf("%t s1=%s s2=%s s3=%s h=%t r=%t", h == r, s1, s2, s3, h, r)

	s1 = "a"
	s2 = "b"
	s3 = "bb"
	h = false
	r = isInterleave(s1, s2, s3)
	t.Logf("%t s1=%s s2=%s s3=%s h=%t r=%t", h == r, s1, s2, s3, h, r)
}
