package sort_colors

import "testing"

func TestSortColors(t *testing.T) {
    var nums, hope []int

    nums = []int{2,0,2,1,1,0}
    t.Logf("before=%v\n", nums)
    sortColors(nums)
    hope = []int{0,0,1,1,2,2}
    t.Logf("\nafter=%v \n hope=%v\n", nums, hope)

    nums = []int{}
    t.Logf("before=%v\n", nums)
    sortColors(nums)
    hope = []int{}
    t.Logf("\nafter=%v \n hope=%v\n", nums, hope)

    nums = []int{1, 0}
    t.Logf("before=%v\n", nums)
    sortColors(nums)
    hope = []int{0, 1}
    t.Logf("\nafter=%v \n hope=%v\n", nums, hope)

    nums = []int{2, 1, 0}
    t.Logf("before=%v\n", nums)
    sortColors(nums)
    hope = []int{0, 1, 2}
    t.Logf("\nafter=%v \n hope=%v\n", nums, hope)
}