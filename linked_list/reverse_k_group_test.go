package linked_list

import (
	"testing"
)

func TestReversKGroup(t *testing.T) {
	var head *ListNode
	var k int
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
	k = 2
	head = reverseKGroup(head, k)
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
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}
	k = 3
	head = reverseKGroup(head, k)
	h = []int{3, 2, 1, 4, 5}
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
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}
	k = 6
	head = reverseKGroup(head, k)
	h = []int{1, 2, 3, 4, 5}
	r = []int{}
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	t.Logf("\nh=%v \nr=%v", h, r)
}
