package linked_list

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	var head *ListNode
	var hope, ret bool

	head = &ListNode{1, &ListNode{2, nil}}
	hope = false
	ret = isPalindrome(head)
	t.Log(ret == hope, hope, ret)

	head = &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	hope = true
	ret = isPalindrome(head)
	t.Log(ret == hope, hope, ret)

	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{1, nil}}}}
	hope = false
	ret = isPalindrome(head)
	t.Log(ret == hope, hope, ret)

	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}
	hope = true
	ret = isPalindrome(head)
	t.Log(ret == hope, hope, ret)
}