package list_node

import "testing"

func TestRotateRight(t *testing.T) {
	var head, h, ret *ListNode

	head = new(ListNode)
	h = head
	for i := 1; i < 5; i ++ {
		h.Val = i
		h.Next = new(ListNode)
		h = h.Next
	}
	h.Val = 5
	t.Log("\nk=2")
	ret = rotateRight(head, 2)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}

	head = new(ListNode)
	h = head
	for i := 1; i < 5; i ++ {
		h.Val = i
		h.Next = new(ListNode)
		h = h.Next
	}
	h.Val = 5
	t.Log("\nk=0")
	ret = rotateRight(head, 0)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}

	head = new(ListNode)
	h = head
	for i := 1; i < 5; i ++ {
		h.Val = i
		h.Next = new(ListNode)
		h = h.Next
	}
	h.Val = 5
	t.Log("\nk=5")
	ret = rotateRight(head, 5)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}

	head = new(ListNode)
	h = head
	for i := 1; i < 5; i ++ {
		h.Val = i
		h.Next = new(ListNode)
		h = h.Next
	}
	h.Val = 5
	t.Log("\nk=1")
	ret = rotateRight(head, 1)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}
}