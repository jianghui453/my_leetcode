package leet_code

import (
	"strings"
	"fmt"
)

func intToRoman(num int) string {
	// var hash = make(map[int]string)
	hash := map[int]string{
		1000: "M",
		500: "D",
		100: "C",
		50: "L",
		10: "X",
		5: "V",
		1: "I",
		4: "IV",
		9: "IX",
		40: "XL",
		90: "XC",
		400: "CD",
		900: "CM",
	}
	nums := []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}
	sb := strings.Builder{}
	for _, v := range nums {
		if num < v {
			continue
		}
		mutiple := num / v
		for ; mutiple > 0; mutiple -- {
			sb.WriteString(hash[v])
		}
		fmt.Printf("num = %d; mutiple = %d; k = %d\n", num, mutiple, v)
		num = num % v
		if num == 0 {
			break
		}
	}
	return sb.String()
}