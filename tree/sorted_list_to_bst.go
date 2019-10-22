//给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
//
//本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
//示例:
//
//给定的有序链表： [-10, -3, 0, 5, 9],
//
//一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5

package tree

//func sortedListToBST(head *ListNode) *TreeNode {
//    var r *TreeNode
//    if head == nil {
//        return r
//    }
//    nums := make([]int, 0)
//    for head != nil {
//        nums = append(nums, head.Val)
//        head = head.Next
//    }
//    var f func(n *TreeNode, ns []int)
//    f = func (n *TreeNode, ns []int) {
//        if len(ns) == 1 {
//            n.Val = ns[0]
//            return
//        }
//        m := len(ns)/2
//        n.Val = ns[m]
//        if m > 0 {
//            n.Left = new(TreeNode)
//            f(n.Left, ns[: m])
//        }
//        if m < len(ns) - 1 {
//            n.Right = new(TreeNode)
//            f(n.Right, ns[m+1: ])
//        }
//    }
//    r = new(TreeNode)
//    f(r, nums)
//    return r
//}

func sortedListToBST(head *ListNode) *TreeNode {
	var list []int
	for p := head; p != nil; p = p.Next {
		list = append(list, p.Val)
	}

	var build func([]int) *TreeNode
	build = func(list []int) *TreeNode {
		if len(list) == 0 {
			return nil
		}
		mid := len(list) / 2
		return &TreeNode{
			Val:   list[mid],
			Left:  build(list[:mid]),
			Right: build(list[mid+1:]),
		}
	}

	return build(list)
}
