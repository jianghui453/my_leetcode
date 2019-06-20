package leet_code

import (
	"testing"
)

func Test_addTwoNumbersMine(t *testing.T) {
	// l1
	l1 := new(ListNode)
	l1.Val = 3
	tmpt := l1
	l1 = new(ListNode)
	l1.Next = tmpt
	l1.Val = 4
	tmpt = l1
	l1 = new(ListNode)
	l1.Next = tmpt
	l1.Val = 2
	//t.Log(l1.Next.Val)
	// l2
	l2 := new(ListNode)
	l2.Val = 4
	tmpt = l2
	l2 = new(ListNode)
	l2.Next = tmpt
	l2.Val = 6
	tmpt = l2
	l2 = new(ListNode)
	l2.Next = tmpt
	l2.Val = 5
	rNL := addTwoNumbersMine(l1, l2)
	//rNL := l2
	for {
		t.Log(rNL.Val)
		if rNL.Next == nil {
			break
		}
		rNL = rNL.Next
	}
}
