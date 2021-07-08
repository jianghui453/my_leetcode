package divide

import (
	"math"
	"testing"
)

func TestDevice(t *testing.T) {
	var divident, divisor, h, r int

	divident = -2147483648
	divisor = -1
	h = math.MaxInt32
	r = divide(divident, divisor)
	t.Logf("%t divident=%d divisor=%d h=%d r=%d", r == h, divident, divisor, h, r)

	divident = -5
	divisor = -1
	h = 5
	r = divide(divident, divisor)
	t.Logf("%t divident=%d divisor=%d h=%d r=%d", r == h, divident, divisor, h, r)

	divident = 5
	divisor = -1
	h = -5
	r = divide(divident, divisor)
	t.Logf("%t divident=%d divisor=%d h=%d r=%d", r == h, divident, divisor, h, r)

	divident = 0
	divisor = -1
	h = 0
	r = divide(divident, divisor)
	t.Logf("%t divident=%d divisor=%d h=%d r=%d", r == h, divident, divisor, h, r)
}
