package leet_code

import "testing"

func Test_singleNumberMine(t *testing.T) {
	nums := []int{1, 2, 3, 3, 2}
	if num := singleNumberMine(nums); num != 1 {
		t.Error("只出现一次的数字测试没通过")
	} else {
		t.Log("只出现一次的数字测试通过了")
	}
}

func Test_singleNumberBest(t *testing.T) {
	nums := []int{1, 2, 3, 3, 2}
	if num := singleNumberBest(nums); num != 1 {
		t.Error("只出现一次的数字测试没通过")
	} else {
		t.Log("只出现一次的数字测试通过了")
	}
}
