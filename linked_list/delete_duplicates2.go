//给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
//
//示例 1:
//
//输入: 1->2->3->3->4->4->5
//输出: 1->2->5
//示例 2:
//
//输入: 1->1->1->2->3
//输出: 2->3

package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates2(head *ListNode) *ListNode {
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
		} else if lenRec > 0 && record[lenRec-1].Val == curVal {
			record = record[:lenRec-1]
			lenRec--
		}
		h = h.Next
	}
	if lenRec == 0 {
		return nil
	}
	for i := 0; i < lenRec-1; i++ {
		record[i].Next = record[i+1]
	}
	record[lenRec-1].Next = nil
	return record[0]
}
