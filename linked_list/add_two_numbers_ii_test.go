package linked_list

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	var l1, l2 *ListNode
	var hope, ret []int

	l1 = &ListNode{7, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}
	l2 = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	hope = []int{7, 8, 0, 7}
	ret = toSlice(addTwoNumbers(l1, l2))
	t.Log(hope, ret)
}

func toSlice(node *ListNode) []int {
	ret := make([]int, 0)
	for ; node != nil; node = node.Next {
		ret = append(ret, node.Val)
	}
	return ret
}
