package backtracking

import "testing"

func TestPartition(t *testing.T) {
	var s string
	var h, r [][]string

	s = "aab"
	h = [][]string{
		{"aa", "b"},
		{"a", "a", "b"},
	}
	r = partition(s)
	t.Logf("\n h=%v \n r=%v", h, r)

	s = ""
	h = [][]string{}
	r = partition(s)
	t.Logf("\n h=%v \n r=%v", h, r)

	s = "aa"
	h = [][]string{
		{"aa"},
		{"a", "a"},
	}
	r = partition(s)
	t.Logf("\n h=%v \n r=%v", h, r)

	s = "aba"
	h = [][]string{
		{"a", "b", "a"},
		{"aba"},
	}
	r = partition(s)
	t.Logf("\n h=%v \n r=%v", h, r)

	s = "abbab"
	h = [][]string{
		{"a", "b", "b", "a", "b"},
		{"a", "b", "bab"},
		{"a", "bb", "a", "b"},
		{"abba", "b"},
	}
	r = partition(s)
	t.Logf("\n h=%v \n r=%v", h, r)
}
