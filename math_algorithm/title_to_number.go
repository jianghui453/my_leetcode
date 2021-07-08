package math_algorithm

func titleToNumber(columnTitle string) int {
	ret := 0
	for _, b := range columnTitle {
		bval := int(b - 'A') + 1
		ret = ret * 26 + bval
	}
	return ret
}
