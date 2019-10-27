package linked_list

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	var head *ListNode
	var n int
	var h, r []int

	head = &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}
	n = 2
	h = []int{1, 2, 3, 5}
	head = removeNthFromEnd(head, n)
	r = make([]int, 0)
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	t.Logf("n=%d \nh=%v \nr=%v", n, h, r)

	head = &ListNode{
		1,
		nil,
	}
	n = 1
	h = []int{}
	head = removeNthFromEnd(head, n)
	r = make([]int, 0)
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	t.Logf("n=%d \nh=%v \nr=%v", n, h, r)
}
