package linked_list

import (
	"testing"
)

func TestSwapPairs(t *testing.T) {
	var head *ListNode
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
	head = swapPairs(head)
	h = []int{2, 1, 4, 3, 5}
	r = []int{}
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	t.Logf("\nh=%v \nr=%v", h, r)

	head = &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					nil,
				},
			},
		},
	}
	head = swapPairs(head)
	h = []int{2, 1, 4, 3}
	r = []int{}
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	t.Logf("\nh=%v \nr=%v", h, r)
}
