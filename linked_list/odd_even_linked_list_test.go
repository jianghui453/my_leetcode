package linked_list

import (
	"testing"
)

func TestOddEvenList(t *testing.T) {
	var (
		head      *ListNode
		hope, ret []int
	)

	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	hope = []int{1, 3, 5, 2, 4}
	ret = ToArray(oddEvenList(head))
	t.Log(hope)
	t.Log(ret)

	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	hope = []int{1, 3, 5, 2, 4, 6}
	ret = ToArray(oddEvenList(head))
	t.Log(hope)
	t.Log(ret)

	head = &ListNode{1, &ListNode{2, nil}}
	hope = []int{1, 2}
	ret = ToArray(oddEvenList(head))
	t.Log(hope)
	t.Log(ret)
}
