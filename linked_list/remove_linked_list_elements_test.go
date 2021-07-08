package linked_list

import (
	"testing"
)

func Test_removeElements(t *testing.T) {
	var (
		head      *ListNode
		val       int
		hope, ret []int
	)

	head = &ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}
	val = 6
	hope = []int{1, 2, 3, 4, 5}
	ret = toArray(removeElements(head, val))
	t.Log(ret)
	t.Log(hope)

	head = &ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}
	val = 1
	hope = []int{2, 6, 3, 4, 5, 6}
	ret = toArray(removeElements(head, val))
	t.Log(ret)
	t.Log(hope)

	head = &ListNode{1, nil}
	val = 1
	hope = []int{}
	ret = toArray(removeElements(head, val))
	t.Log(ret)
	t.Log(hope)
}

func toArray(head *ListNode) []int {
	ret := make([]int, 0)
	for head != nil {
		ret, head = append(ret, head.Val), head.Next
	}
	return ret
}
