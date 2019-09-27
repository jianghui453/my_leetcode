//在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
//
//示例 1:
//
//输入: 4->2->1->3
//输出: 1->2->3->4
//示例 2:
//
//输入: -1->5->3->4->0
//输出: -1->0->3->4->5

package linked_list

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    fast, slow := head, head
    for {
        if fast.Next != nil && fast.Next.Next != nil {
            fast = fast.Next.Next
            slow = slow.Next
        } else {
            break
        }
    }
    h := slow.Next
    slow.Next = nil
    head = sortList(head)
    h = sortList(h)
    return merge(head, h)
}

func merge(h1, h2 *ListNode) *ListNode {
    if h1 == nil {
        return h2
    }
    if h2 == nil {
        return h1
    }
    var h *ListNode
    if h1.Val < h2.Val {
        h = h1
        h1 = h1.Next
    } else {
        h = h2
        h2 = h2.Next
    }
    t := h
    for h1 != nil && h2 != nil {
        if h1.Val < h2.Val {
            t.Next = h1
            h1 = h1.Next
        } else {
            t.Next = h2
            h2 = h2.Next
        }
        t = t.Next
    }
    if h1 != nil {
        t.Next = h1
    }
    if h2 != nil {
        t.Next = h2
    }
    return h
}
