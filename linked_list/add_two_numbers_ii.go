// 给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。

//  

// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。

// 进阶:

// 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

// 示例:

// 输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出: 7 -> 8 -> 0 -> 7

package linked_list

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

}

func reverseList(l *ListNode) *ListNode {
	if l == nil {
		return l
	}

	h := l
	l = l.Next
	for l != nil {
		l.Next, l, h1 = h1, l.Next, l
	}

	return 
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	nums1, nums2 := make([]int, 0), make([]int, 0)
// 	for ; l1 != nil; l1 = l1.Next { 
// 		nums1 = append(nums1, l1.Val)
// 	}
// 	for ; l2 != nil; l2 = l2.Next {
// 		nums2 = append(nums2, l2.Val)
// 	}

// 	len1, len2 := len(nums1)-1, len(nums2)-1
// 	nums, carry := make([]int, 0), 0
// 	for len1 >= 0 && len2 >= 0 {
// 		v := nums1[len1]+nums2[len2]+carry
// 		nums = append(nums, v%10)
// 		carry = v / 10
// 		len1, len2 = len1-1, len2-1
// 	}

// 	for len1 >= 0 {
// 		v := nums1[len1]+carry
// 		nums = append(nums, v%10)
// 		carry = v / 10
// 		len1--
// 	}

// 	for len2 >= 0 {
// 		v := nums2[len2]+carry
// 		nums = append(nums, v%10)
// 		carry = v / 10
// 		len2--
// 	}

// 	if carry > 0 {
// 		nums = append(nums, carry)
// 	}

// 	l := len(nums)
// 	var head, tail *ListNode
// 	for i := l-1; i >= 0; i-- {
// 		if head == nil {
// 			head = new(ListNode)
// 			head.Val = nums[i]
// 			tail = head
// 		} else {
// 			tail.Next = &ListNode{nums[i], nil}
// 			tail = tail.Next
// 		}
// 	}

// 	return head
// }