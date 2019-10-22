package is_valid_sudoku

// import "fmt"

func isValidSudoku(board [][]byte) bool {
	var raw [9][10]int
	// fmt.Printf("raw = %v\n", raw)
	// return true
	var column [9][10]int
	var square [9][10]int
	for k, list := range board {
		for key, val := range list {
			if val == '.' {
				continue
			}
			val = val - '0'
			var squareKey int
			if 0 <= key && key <= 2 {
				if 0 <= k && k <= 2 {
					squareKey = 0
				} else if 3 <= k && k <= 5 {
					squareKey = 1
				} else {
					squareKey = 2
				}
			} else if 3 <= key && key <= 5 {
				if 0 <= k && k <= 2 {
					squareKey = 3
				} else if 3 <= k && k <= 5 {
					squareKey = 4
				} else {
					squareKey = 5
				}
			} else {
				if 0 <= k && k <= 2 {
					squareKey = 6
				} else if 3 <= k && k <= 5 {
					squareKey = 7
				} else {
					squareKey = 8
				}
			}
			// fmt.Printf("k = %d; key = %d; val = %d;\n", k, key, val)
			if raw[k][val] > 0 ||
				column[key][val] > 0 ||
				square[squareKey][val] > 0 {
				return false
			}
			raw[k][val]++
			column[key][val]++
			square[squareKey][val]++
		}
	}
	return true
}
