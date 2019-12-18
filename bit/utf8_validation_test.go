package bit

import (
	"testing"
)

func Test_validUtf8(t *testing.T) {
	var data []int
	var hope, ret bool

	data = []int{197, 130, 1}
	hope = true
	ret = validUtf8(data)
	t.Log(ret == hope, data, hope, ret)

	data = []int{235, 140, 4}
	hope = false
	ret = validUtf8(data)
	t.Log(ret == hope, data, hope, ret)

	data = []int{237}
	hope = false
	ret = validUtf8(data)
	t.Log(ret == hope, data, hope, ret)
}