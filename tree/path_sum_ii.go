//给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//给定如下二叉树，以及目标和 sum = 22，
//
//              5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   1
//返回:
//
//[
//   [5,4,11,2],
//   [5,8,4,5]
//]
package tree

import "fmt"

func pathSum(root *TreeNode, sum int) [][]int {
	ret := make([][]int, 0)
	var f func(node *TreeNode, s int, r []int)
	f = func(node *TreeNode, s int, r []int) {
		fmt.Printf("s=%d r=%v\n", s, r)
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if node.Val == s {
				r = append(r, node.Val)
				ret = append(ret, r)
			}
			return
		}
		r = append(r, node.Val)
		if node.Left != nil {
			rNew := make([]int, len(r))
			copy(rNew, r)
			f(node.Left, s-node.Val, rNew)
		}
		if node.Right != nil {
			rNew := make([]int, len(r))
			copy(rNew, r)
			f(node.Right, s-node.Val, rNew)
		}
	}
	f(root, sum, []int{})
	return ret
}
