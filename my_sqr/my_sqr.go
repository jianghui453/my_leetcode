package my_sqr

import "math"
import "fmt"

func mySqr(n float64) float64 {
	ret := n/2
	for math.Abs((ret + n/ret)/2 - ret) >= 0.0000000001 {
		ret = math.Abs(ret + n/ret)/2
		fmt.Printf("ret=%g next=%g codition=%g\n", ret, math.Abs(ret + n/ret)/2, math.Abs((ret + n/ret)/2 - ret))
	}
	return ret
}