package leet_code

import (
	"fmt"
)

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		l1, l2 := lists[0], lists[1]
		if l1 == nil {
			lists = append(lists, l2)
			lists = lists[2:]
			continue
		}
		if l2 == nil {
			lists = append(lists, l1)
			lists = lists[2:]
			continue
		}
		if l1.Val > l2.Val {
			l1, l2 = l2, l1
		}
		start, end := l1, l1
		l1 = l1.Next
		for l1 != nil && l2 != nil {
			if l1.Val > l2.Val {
				l1, l2 = l2, l1
			}
			end.Next = l1
			end = end.Next
			l1 = l1.Next
		}
		if l1 != nil {
			end.Next = l1
		} else if l2 != nil {
			end.Next = l2
		}
		lists = append(lists, start)
		lists = lists[2:]
	}
	return lists[0]
}

func mergeKListsV1(lists []*ListNode) *ListNode {
	var start, end *ListNode
	var minK int
	fmt.Println(lists[0])
	for {
		minK = -1
		for k, list := range lists {
			fmt.Printf("k = %d; minK = %d; len(lists) = %d\n", k, minK, len(lists))
			if list == nil {
				continue
			}
			if minK == -1 || list.Val < lists[minK].Val {
				minK = k
			}
		}
		if minK == -1 {
			break
		}
		if end == nil {
			end = lists[minK]
			lists[minK] = lists[minK].Next
			start = end
		} else {
			end.Next = lists[minK]
			lists[minK] = lists[minK].Next
			end = end.Next
		}
	}
	return start
}