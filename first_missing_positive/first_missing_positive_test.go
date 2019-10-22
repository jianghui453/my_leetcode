package first_missing_positive

import "testing"
import "fmt"

func TestFirstMissingPositive(t *testing.T) {
	var nums []int
	var hope, ret int

	nums = []int{1, 2, 0}
	hope = 3
	ret = firstMissingPositive(nums)
	fmt.Printf("nums = %v, hope = %d, ret = %d.\n", nums, hope, ret)

	nums = []int{3, 4, -1, 1}
	hope = 2
	ret = firstMissingPositive(nums)
	fmt.Printf("nums = %v, hope = %d, ret = %d.\n", nums, hope, ret)

	nums = []int{7, 8, 9, 11, 12}
	hope = 1
	ret = firstMissingPositive(nums)
	fmt.Printf("nums = %v, hope = %d, ret = %d.\n", nums, hope, ret)

	nums = []int{}
	hope = 1
	ret = firstMissingPositive(nums)
	fmt.Printf("nums = %v, hope = %d, ret = %d.\n", nums, hope, ret)
}
