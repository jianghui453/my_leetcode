//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
//示例 1:
//
//输入: 1->1->2
//输出: 1->2
//示例 2:
//
//输入: 1->1->2->3->3
//输出: 1->2->3

package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	record := make([]*ListNode, 0)
	record = append(record, head)
	lenRec := 1
	curVal := head.Val
	h := head.Next
	for h != nil {
		if h.Val != curVal {
			curVal = h.Val
			record = append(record, h)
			lenRec++
		}
		h = h.Next
	}
	for i := 0; i < lenRec-1; i++ {
		record[i].Next = record[i+1]
	}
	record[lenRec-1].Next = nil
	return record[0]
}
