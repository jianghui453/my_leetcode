package leet_code

import (
	"testing"
)

func TestMergeKLists(t *testing.T) {
	var lists = []*ListNode{
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

	_ = mergeKLists(lists)
}