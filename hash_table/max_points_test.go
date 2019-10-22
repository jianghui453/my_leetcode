package hash_table

import "testing"

func TestMaxPoints(t *testing.T) {
	var points [][]int
	var h, r int

	points = [][]int{{1, 1}, {2, 2}, {3, 3}}
	h = 3
	r = maxPoints(points)
	t.Logf("%t points=%v h=%d r=%d", h == r, points, h, r)

	points = [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	h = 4
	r = maxPoints(points)
	t.Logf("%t points=%v h=%d r=%d", h == r, points, h, r)

	points = [][]int{{1, 1}, {1, 1}, {1, 1}}
	h = 3
	r = maxPoints(points)
	t.Logf("%t points=%v h=%d r=%d", h == r, points, h, r)

	points = [][]int{{1, 1}, {0, 1}, {1, 0}}
	h = 2
	r = maxPoints(points)
	t.Logf("%t points=%v h=%d r=%d", h == r, points, h, r)

	points = [][]int{{1, 1}, {2, 2}, {1, 1}}
	h = 3
	r = maxPoints(points)
	t.Logf("%t points=%v h=%d r=%d", h == r, points, h, r)

	points = [][]int{{1, 1}, {2, 2}, {1, 1}, {2, 2}, {1, 1}}
	h = 5
	r = maxPoints(points)
	t.Logf("%t points=%v h=%d r=%d", h == r, points, h, r)

	points = [][]int{{1, 1}, {1, 1}, {2, 3}}
	h = 3
	r = maxPoints(points)
	t.Logf("%t points=%v h=%d r=%d", h == r, points, h, r)

	points = [][]int{{4, 0}, {4, -1}, {4, 5}}
	h = 3
	r = maxPoints(points)
	t.Logf("%t points=%v h=%d r=%d", h == r, points, h, r)

	points = [][]int{{0, 0}, {94911151, 94911150}, {94911152, 94911151}}
	h = 2
	r = maxPoints(points)
	t.Logf("%t points=%v h=%d r=%d", h == r, points, h, r)
}
