package leet_code

func romanToInt(s string) int {
	hash := map[string]int{
		"I": 1,
		"IV": 4,
		"V": 5,
		"IX": 9,
		"X": 10,
		"XL": 40,
		"L": 50,
		"XC": 90,
		"C": 100,
		"CD": 400,
		"D": 500,
		"CM": 900,
		"M": 1000,
	}
	var ret, i int
	var key string
	for i < len(s) {
		if _, ok := hash[string(s[i])]; !ok {
			return 0
		}
		if i + 1 < len(s) {
			if _, ok := hash[string(s[i + 1])]; !ok {
				return 0
			}
		}
		if (i + 2 < len(s)) {
			key = s[i: i + 2]
		} else {
			key = s[i:]
		}
		if _, ok := hash[key]; ok {
			ret += hash[key]
			i +=2
			continue
		}
		ret += hash[string(s[i])]
		i++
	}
	return ret
}