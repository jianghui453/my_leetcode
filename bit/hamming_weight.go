package bit

func hammingWeight(num uint32) int {
	var out int
	for ; num > 0; num = num / 2 {
		if num % 2 == 1 {
			out++
		}
	}
	return out
}
