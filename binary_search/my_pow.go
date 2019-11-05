// Implement pow(x, n), which calculates x raised to the power n (xn).

// Example 1:

// Input: 2.00000, 10
// Output: 1024.00000
// Example 2:

// Input: 2.10000, 3
// Output: 9.26100
// Example 3:

// Input: 2.00000, -2
// Output: 0.25000
// Explanation: 2-2 = 1/22 = 1/4 = 0.25
// Note:

// -100.0 < x < 100.0
// n is a 32-bit signed integer, within the range [−231, 231 − 1]

package binary_search

func myPow(x float64, n int) float64 {
	if n == 0 {
		return float64(1)
	}

	var ret float64
	if n < 0 {
		ret = float64(1) / myPow(x, -n)
	} else {
		ret = myPow(x*x, n/2)
		if n%2 == 1 {
			ret *= x
		}
	}
	
	return ret
}

// func myPow(x float64, n int) float64 {
// 	if n < 0 {
// 		x = 1 / x
// 		n = -n
// 	}
// 	return _pow(x, n)
// }

// func _pow(x float64, n int) float64 {
// 	if n == 0 {
// 		return 1.0
// 	}
// 	if n == 1 {
// 		return x
// 	}
// 	if n%2 == 0 {
// 		return _pow(x*x, n/2)
// 	} else {
// 		return _pow(x*x, (n-1)/2) * x
// 	}
// }

// func myPow(x float64, n int) float64 {
// 	if n < 0 {
// 		x = 1 / x
// 		n = -n
// 	}
// 	ret := 1.0
// 	for i := n; i > 0; i /= 2 {
// 		if i%2 == 1 {
// 			ret *= x
// 		}
// 		x *= x
// 	}
// 	return ret
// }

// func myPow(x float64, n int) float64 {
// 	if n < 0 {
// 		x = 1 / x
// 		n = -n
// 	}
// 	return pow(x, n)
// }

// func pow(x float64, n int) float64 {
// 	if n == 0 {
// 		return 1.0
// 	}
// 	if n == 1 {
// 		return x
// 	}
// 	if n%2 == 0 {
// 		return pow(x, n/2) * pow(x, n/2)
// 	} else {
// 		return pow(x, n-1/2) * pow(x, n-1/2) * x
// 	}
// }
