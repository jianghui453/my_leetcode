package leet_code

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 22
	res := twoSum(nums, target)
	fmt.Printf("twoSum=%v\n", res)
	for _, v := range res {
		target -= nums[v]
	}
	if target == 0 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
