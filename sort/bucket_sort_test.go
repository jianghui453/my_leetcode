package sort

import "testing"

func TestBucketSort(t *testing.T) {
	var nums []float64

	nums = []float64{0.897, 0.565, 0.656, 0.1234, 0.665, 0.3434}
	//nums = []float64{2, 1}
	t.Logf("start nums=%v", nums)
	bucketSort(nums)
	t.Logf("end nums=%v", nums)
}
