//给定一个二叉树，返回它的 前序 遍历。
//
// 示例:
//
//输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,2,3]

package tree

func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		ret = append(ret, node.Val)
		f(node.Left)
		f(node.Right)
	}
	f(root)
	return ret
}
