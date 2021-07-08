package math_algorithm

import (
	"math"
	"strconv"
	"strings"
// 	"fmt"
)

func fractionToDecimal(numerator int, denominator int) string {
	sign := ""
	if numerator * denominator < 0 {
		sign = "-"
	}
	numerator = int(math.Abs(float64(numerator)))
	denominator = int(math.Abs(float64(denominator)))
	// fmt.Println("numerator:", numerator, "denominator:", denominator)

	integer := []string{sign}
	for numerator > 0 && len(integer) <= 10000 {
		res, reminder := numerator / denominator, numerator % denominator
		// fmt.Println("res:", res, "reminder:", reminder)
		numerator = reminder
		if res == 0 {
			break
		}

		integer = append(integer, strconv.Itoa(res))
	}
	if len(integer) <= 1 {
		integer = append(integer, "0")
	}
	// fmt.Println("integer:", integer)

	decimal := []string{}
	numerator *= 10
	repeated := map[int]int{numerator: 0}
	for i := 1; numerator > 0 && len(decimal) <= 10000; i++ {
		res, reminder := numerator / denominator, numerator % denominator
		decimal = append(decimal, strconv.Itoa(res))

		numerator = reminder * 10
		if idx, ok := repeated[numerator]; ok {
			decimal1 := make([]string, idx)
			copy(decimal1, decimal[:idx])

			decimal2 := make([]string, len(decimal) - idx + 2)
			copy(decimal2[1: len(decimal2)-1], decimal[idx:])
			decimal2[0] = "("
			decimal2[len(decimal2)-1] = ")"
			// fmt.Println("decimal1:", decimal1, "decimal2:", decimal2, "len(decimal2):", len(decimal2))
			decimal = append(decimal1, decimal2...)
			break
		}
		repeated[numerator] = i
	}
	if len(decimal) == 0 || (len(decimal) == 1 && decimal[0] == "0") {
		return strings.Join(integer, "")
	}
	return strings.Join(integer, "") + "." + strings.Join(decimal, "")
}
