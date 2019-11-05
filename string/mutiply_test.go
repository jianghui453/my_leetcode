package string

import "testing"

func TestMutiply(t *testing.T) {
	var num1, num2, h, r string

	num1 = "2"
	num2 = "3"
	h = "6"
	r = multiply(num1, num2)
	t.Logf("num1=%s num2=%s h=%s r=%s\n", num1, num2, h, r)

	num1 = "123"
	num2 = "456"
	h = "56088"
	r = multiply(num1, num2)
	t.Logf("num1=%s num2=%s h=%s r=%s\n", num1, num2, h, r)

	num1 = "0"
	num2 = "0"
	h = "0"
	r = multiply(num1, num2)
	t.Logf("num1=%s num2=%s h=%s r=%s\n", num1, num2, h, r)

	num1 = "498828660196"
	num2 = "840477629533"
	h = "419254329864656431168468"
	r = multiply(num1, num2)
	t.Logf("num1=%s num2=%s h=%s r=%s\n", num1, num2, h, r)

	num1 = "9133"
	num2 = "0"
	h = "0"
	r = multiply(num1, num2)
	t.Logf("num1=%s num2=%s h=%s r=%s\n", num1, num2, h, r)
}
