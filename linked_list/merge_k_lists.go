// 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
// 示例:
// 输入:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// 输出: 1->1->2->3->4->4->5->6

package linked_list

// import (
// 	"fmt"
// )

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		l1, l2 := lists[0], lists[1]

		var h, t *ListNode
		for l1 != nil && l2 != nil {
			if l1.Val > l2.Val {
				l1, l2 = l2, l1
			}
			if h == nil {
				h, t = l1, l1
			} else {
				t.Next = l1
				t = t.Next
			}
			l1 = l1.Next
		}

		if l1 != nil {
			l1, l2 = l2, l1
		}

		for l2 != nil {
			if h == nil {
				h, t = l2, l2
			} else {
				t.Next = l2
				t = t.Next
			}
			l2 = l2.Next
		}
		
		lists[0] = h
		if len(lists) == 2 {
			return h
		}
		lists = append(lists[: 1], lists[2: ]...)
	}
	return lists[0]
}

// func mergeKLists(lists []*ListNode) *ListNode {
// 	if len(lists) == 0 {
// 		return nil
// 	}
// 	for len(lists) > 1 {
// 		l1, l2 := lists[0], lists[1]
// 		if l1 == nil {
// 			lists = append(lists, l2)
// 			lists = lists[2:]
// 			continue
// 		}
// 		if l2 == nil {
// 			lists = append(lists, l1)
// 			lists = lists[2:]
// 			continue
// 		}
// 		if l1.Val > l2.Val {
// 			l1, l2 = l2, l1
// 		}
// 		start, end := l1, l1
// 		l1 = l1.Next
// 		for l1 != nil && l2 != nil {
// 			if l1.Val > l2.Val {
// 				l1, l2 = l2, l1
// 			}
// 			end.Next = l1
// 			end = end.Next
// 			l1 = l1.Next
// 		}
// 		if l1 != nil {
// 			end.Next = l1
// 		} else if l2 != nil {
// 			end.Next = l2
// 		}
// 		lists = append(lists, start)
// 		lists = lists[2:]
// 	}
// 	return lists[0]
// }

// func mergeKLists(lists []*ListNode) *ListNode {
// 	var start, end *ListNode
// 	var minK int
// 	fmt.Println(lists[0])
// 	for {
// 		minK = -1
// 		for k, list := range lists {
// 			fmt.Printf("k = %d; minK = %d; len(lists) = %d\n", k, minK, len(lists))
// 			if list == nil {
// 				continue
// 			}
// 			if minK == -1 || list.Val < lists[minK].Val {
// 				minK = k
// 			}
// 		}
// 		if minK == -1 {
// 			break
// 		}
// 		if end == nil {
// 			end = lists[minK]
// 			lists[minK] = lists[minK].Next
// 			start = end
// 		} else {
// 			end.Next = lists[minK]
// 			lists[minK] = lists[minK].Next
// 			end = end.Next
// 		}
// 	}
// 	return start
// }
