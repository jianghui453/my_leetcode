package leet_code

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				4,
				nil,
			},
		},
	}

	l2 := &ListNode{
		1,
		&ListNode{
			3,
			&ListNode{
				4,
				nil,
			},
		},
	}

	ret := mergeTwoLists(l1, l2)
	t.Logf("l1 = %v; l2 = %v; ret = %v", l1, l2, ret)
}