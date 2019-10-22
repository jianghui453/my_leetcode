package str_str

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	haystack := "a"
	needle := "a"

	ret := strStr(haystack, needle)
	t.Logf("ret = %d", ret)
}
