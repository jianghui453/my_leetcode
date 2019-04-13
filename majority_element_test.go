package leet_code

import "testing"

func Test_majorityElement(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	if 2 == majorityElementMine(nums) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
