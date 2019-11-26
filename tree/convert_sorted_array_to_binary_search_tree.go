//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
//
//本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
//示例:
//
//给定有序数组: [-10,-3,0,5,9],
//
//一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5

package tree

func sortedArrayToBST(nums []int) *TreeNode {
	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	}

	node := new(TreeNode)
	mid := (numsLen-1)/2
	node.Val = nums[mid]
	node.Left = sortedArrayToBST(nums[: mid])
	if mid < numsLen-1 {
		node.Right = sortedArrayToBST(nums[mid+1: ])
	}

	return node
}

// func sortedArrayToBST(nums []int) *TreeNode {
// 	var r *TreeNode
// 	if len(nums) == 0 {
// 		return r
// 	}
// 	var f func(n *TreeNode, ns []int)
// 	f = func(n *TreeNode, ns []int) {
// 		if len(ns) == 1 {
// 			n.Val = ns[0]
// 			return
// 		}
// 		m := len(ns) / 2
// 		n.Val = ns[m]
// 		if m > 0 {
// 			n.Left = new(TreeNode)
// 			f(n.Left, ns[:m])
// 		}
// 		if m < len(ns)-1 {
// 			n.Right = new(TreeNode)
// 			f(n.Right, ns[m+1:])
// 		}
// 	}
// 	r = new(TreeNode)
// 	f(r, nums)
// 	return r
// }
