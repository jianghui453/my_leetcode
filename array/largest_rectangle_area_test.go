package array

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	var heights []int
	var hope, ret int

	heights = []int{2, 1, 5, 6, 2, 3}
	hope = 10
	ret = largestRectangleArea(heights)
	t.Logf("heights=%v hope=%d ret=%d", heights, hope, ret)

	heights = []int{2, 1, 5, 1, 6, 3}
	hope = 6
	ret = largestRectangleArea(heights)
	t.Logf("heights=%v hope=%d ret=%d", heights, hope, ret)

	heights = []int{2}
	hope = 2
	ret = largestRectangleArea(heights)
	t.Logf("heights=%v hope=%d ret=%d", heights, hope, ret)

	heights = []int{1, 1, 1}
	hope = 3
	ret = largestRectangleArea(heights)
	t.Logf("heights=%v hope=%d ret=%d", heights, hope, ret)

	heights = []int{2, 1, 2}
	hope = 3
	ret = largestRectangleArea(heights)
	t.Logf("heights=%v hope=%d ret=%d", heights, hope, ret)

	heights = []int{3, 2, 1, 2}
	hope = 4
	ret = largestRectangleArea(heights)
	t.Logf("heights=%v hope=%d ret=%d", heights, hope, ret)

	heights = []int{5, 4, 1, 2}
	hope = 8
	ret = largestRectangleArea(heights)
	t.Logf("heights=%v hope=%d ret=%d", heights, hope, ret)
}
