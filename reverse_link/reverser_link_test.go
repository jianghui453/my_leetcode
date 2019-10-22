package reverse_link

import "testing"

func TestReverseLink(t *testing.T) {
	var head, ret *ListNode

	head = &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				nil,
			},
		},
	}
	ret = reverseLink(head)
	for ret != nil {
		t.Log(ret.val)
		ret = ret.next
	}
}
