package linked_list

import (
	"testing"
)

func TestMergeKLists(t *testing.T) {
	var lists []*ListNode
	var h, r []int

	lists = []*ListNode{
		&ListNode{
			1,
			&ListNode{
				4,
				&ListNode{
					5,
					nil,
				},
			},
		},
		&ListNode{
			1,
			&ListNode{
				3,
				&ListNode{
					4,
					nil,
				},
			},
		},
		&ListNode{
			2,
			&ListNode{
				6,
				nil,
			},
		},
	}
	h = []int{1, 1, 2, 3, 4, 4, 5, 6}
	head := mergeKLists(lists)
	r = []int{}
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	t.Logf("\nh=%v \nr=%v", h, r)
}
