package multiply

import "testing"

func TestMutiply(t *testing.T) {
	var num1, num2, hope, ret string

	// num1 = "2"
	// num2 = "3"
	// hope = "6"
	// ret = multiply(num1, num2)
	// t.Logf("num1 = %s, num2 = %s, hope = %s, ret = %s.\n", num1, num2, hope, ret)

	// num1 = "123"
	// num2 = "456"
	// hope = "56088"
	// ret = multiply(num1, num2)
	// t.Logf("num1 = %s, num2 = %s, hope = %s, ret = %s.\n", num1, num2, hope, ret)

	// num1 = "0"
	// num2 = "0"
	// hope = "0"
	// ret = multiply(num1, num2)
	// t.Logf("num1 = %s, num2 = %s, hope = %s, ret = %s.\n", num1, num2, hope, ret)

	num1 = "498828660196"
	num2 = "840477629533"
	hope = "419254329864656431168468"
	ret = multiply(num1, num2)
	t.Logf("num1 = %s, num2 = %s, hope = %s, ret = %s.\n", num1, num2, hope, ret)

	// num1 = "9133"
	// num2 = "0"
	// hope = "0"
	// ret = multiply(num1, num2)
	// t.Logf("num1 = %s, num2 = %s, hope = %s, ret = %s.\n", num1, num2, hope, ret)
}