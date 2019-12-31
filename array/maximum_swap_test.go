package array

import (
	"testing"
)

func Test_maximumSwap(t *testing.T) {
	var (
		num, hope, ret int
	)

	num = 2736
	hope = 7236
	ret = maximumSwap(num)
	t.Log(ret == hope, num, hope, ret)

	num = 9973
	hope = 9973
	ret = maximumSwap(num)
	t.Log(ret == hope, num, hope, ret)

	num = 1000
	hope = 1000
	ret = maximumSwap(num)
	t.Log(ret == hope, num, hope, ret)

	num = 0
	hope = 0
	ret = maximumSwap(num)
	t.Log(ret == hope, num, hope, ret)

	num = 1
	hope = 1
	ret = maximumSwap(num)
	t.Log(ret == hope, num, hope, ret)

	num = 12
	hope = 21
	ret = maximumSwap(num)
	t.Log(ret == hope, num, hope, ret)
}
