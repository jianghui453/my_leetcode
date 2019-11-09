package plus_one

import "testing"

func TestPlusOne(t *testing.T) {
	var digits, hope, ret []int

	digits = []int{0}
	hope = []int{1}
	ret = plusOne(digits)
	t.Logf("digits=%v \nhope=%v \nret=%v\n", digits, hope, ret)

	digits = []int{9, 9}
	hope = []int{1, 0, 0}
	ret = plusOne(digits)
	t.Logf("digits=%v \nhope=%v \nret=%v\n", digits, hope, ret)

	digits = []int{9}
	hope = []int{1, 0}
	ret = plusOne(digits)
	t.Logf("digits=%v \nhope=%v \nret=%v\n", digits, hope, ret)

	digits = []int{8}
	hope = []int{9}
	ret = plusOne(digits)
	t.Logf("digits=%v \nhope=%v \nret=%v\n", digits, hope, ret)

	digits = []int{9, 8}
	hope = []int{9, 9}
	ret = plusOne(digits)
	t.Logf("digits=%v \nhope=%v \nret=%v\n", digits, hope, ret)

	digits = []int{8, 9}
	hope = []int{9, 0}
	ret = plusOne(digits)
	t.Logf("digits=%v \nhope=%v \nret=%v\n", digits, hope, ret)
}
