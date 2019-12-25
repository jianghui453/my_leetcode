package math_algorithm

import (
	"testing"
)

func Test_canMeasureWater(t *testing.T) {
	var (
		x, y, z   int
		hope, ret bool
	)

	x, y, z = 3, 5, 4
	hope = true
	ret = canMeasureWater(x, y, z)
	t.Log(ret == hope, x, y, z, hope, ret)

	x, y, z = 2, 6, 5
	hope = false
	ret = canMeasureWater(x, y, z)
	t.Log(ret == hope, x, y, z, hope, ret)
}
