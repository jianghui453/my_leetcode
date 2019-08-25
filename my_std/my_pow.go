package my_std

// import "math"

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1/x
		n = -n
	}
	return _pow(x, n)
}

func _pow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	if n % 2 == 0 {
		return _pow(x * x, n/2)
	} else {
		return _pow(x * x, (n-1)/2)* x
	}
}

func myPowV2(x float64, n int) float64 {
	if n < 0 {
		x = 1/x
		n = -n
	}
	ret := 1.0
	for i := n; i > 0; i /= 2 {
		if i % 2 == 1 {
			ret *= x
		}
		x *= x
	}
	return ret
}

func myPowV1(x float64, n int) float64 {
	if n < 0 {
		x = 1/x
		n = -n
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	if n % 2 == 0 {
		return pow(x, n/2) * pow(x, n/2)
	} else {
		return pow(x, n-1/2) * pow(x, n-1/2) * x
	}
}