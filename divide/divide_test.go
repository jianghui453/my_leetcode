package divide

import (
	"testing"
)

func TestDevice(t *testing.T) {
	var divident, divisor, hope, ret int32

	divident = -2147483648
	divisor = -1
	hope = 2147483648
	ret = divide(divident, divisor)
	if hope == ret {

	}
}
