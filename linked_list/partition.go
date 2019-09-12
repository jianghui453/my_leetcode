//给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
//
//你应当保留两个分区中每个节点的初始相对位置。
//
//示例:
//
//输入: head = 1->4->3->2->5->2, x = 3
//输出: 1->2->2->4->3->5

package linked_list

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	var smaller, larger, sHead, lHead *ListNode
	for head != nil {
		fmt.Println(head.Val)
		if head.Val < x {
			if smaller == nil {
				smaller = head
				sHead = smaller
			} else {
				smaller.Next = head
				smaller = smaller.Next
			}
		} else {
			if larger == nil {
				larger = head
				lHead = larger
			} else {
				larger.Next = head
				larger = larger.Next
			}
		}
		head = head.Next
	}
	if sHead == nil {
		if lHead != nil {
			return lHead
		}
		return nil
	}
	smaller.Next = lHead
	if larger != nil {
		larger.Next = nil
	}
	return sHead
}
