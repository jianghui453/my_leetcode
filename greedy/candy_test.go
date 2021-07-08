package greedy

import "testing"

func TestCandy(t *testing.T) {
	var ratings []int
	var h, r int

	ratings = []int{1, 0, 2}
	h = 5
	r = candy(ratings)
	t.Logf("ratings=%v h=%d r=%d", ratings, h, r)

	ratings = []int{1, 2, 2}
	h = 4
	r = candy(ratings)
	t.Logf("ratings=%v h=%d r=%d", ratings, h, r)
}
