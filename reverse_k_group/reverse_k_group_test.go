package reverse_k_group

import (
	"testing"
)

func TestReversKGroup(t *testing.T) {
	var head, ret *ListNode
	var k int

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
	ret = reverseKGroup(head, k)
	t.Logf("k = %d\n", k)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}

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
	ret = reverseKGroup(head, k)
	t.Logf("k = %d\n", k)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}
}
