package string

import "testing"

func TestAddBinary(t *testing.T) {
	var a, b, hope, ret string

	a = ""
	b = ""
	hope = ""
	ret = addBinary(a, b)
	t.Logf("%t a=%s b=%s hope=%s ret=%s\n", ret == hope, a, b, hope, ret)

	a = ""
	b = "1"
	hope = "1"
	ret = addBinary(a, b)
	t.Logf("%t a=%s b=%s hope=%s ret=%s\n", ret == hope, a, b, hope, ret)

	a = "0"
	b = ""
	hope = "0"
	ret = addBinary(a, b)
	t.Logf("%t a=%s b=%s hope=%s ret=%s\n", ret == hope, a, b, hope, ret)

	a = "11"
	b = "1"
	hope = "100"
	ret = addBinary(a, b)
	t.Logf("%t a=%s b=%s hope=%s ret=%s\n", ret == hope, a, b, hope, ret)

	a = "10"
	b = "1"
	hope = "11"
	ret = addBinary(a, b)
	t.Logf("%t a=%s b=%s hope=%s ret=%s\n", ret == hope, a, b, hope, ret)

	a = "1010"
	b = "1011"
	hope = "10101"
	ret = addBinary(a, b)
	t.Logf("%t a=%s b=%s hope=%s ret=%s\n", ret == hope, a, b, hope, ret)
}
