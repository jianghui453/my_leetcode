// Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

// Example:

// Input: 3
// Output:
// [
//  [ 1, 2, 3 ],
//  [ 8, 9, 4 ],
//  [ 7, 6, 5 ]
// ]

package array

func generateMatrix(n int) [][]int {
	var ret [][]int
	if n == 0 {
		return ret
	}

	for i := 0; i < n; i++ {
		ret = append(ret, make([]int, n))
	}

	nsqu, num, t, r, b, l := n*n, 1, 0, n-1, n-1, 0

	for num <= nsqu {
		for i := l; i <= r; i++ {
			ret[t][i] = num
			num++
		}
		for i := t + 1; i <= b; i++ {
			ret[i][r] = num
			num++
		}
		for i := r - 1; i >= l; i-- {
			ret[b][i] = num
			num++
		}
		for i := b - 1; i > t; i-- {
			ret[i][l] = num
			num++
		}
		t++
		r--
		b--
		l++
	}
	return ret
}
