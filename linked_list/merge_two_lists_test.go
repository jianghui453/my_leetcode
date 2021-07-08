package linked_list

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	var l1, l2, head *ListNode
	var h, r []int

	l1 = &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				4,
				nil,
			},
		},
	}
	l2 = &ListNode{
		1,
		&ListNode{
			3,
			&ListNode{
				4,
				nil,
			},
		},
	}
	h = []int{1, 1, 2, 3, 4, 4}
	r = []int{}
	head = mergeTwoLists(l1, l2)
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	t.Logf("\nh=%v \nr=%v", h, r)
}
