package linked_list

import (
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	var headA, headB, hope, ret *ListNode

	hope = &ListNode{
		8,
		&ListNode{
			4,
			&ListNode{
				5,
				nil,
			},
		},
	}
	headA = &ListNode{
		4,
		&ListNode{
			1,
			hope,
		},
	}
	headB = &ListNode{
		5,
		&ListNode{
			0,
			&ListNode{
				1,
				hope,
			},
		},
	}
	ret = getIntersectionNode(headA, headB)
	t.Logf("%t hope=%v ret=%v", ret == hope, hope, ret)
}