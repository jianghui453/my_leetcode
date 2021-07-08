//给定一个非空二叉树，返回其最大路径和。
//
//本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
//
//示例 1:
//
//输入: [1,2,3]
//
//       1
//      / \
//     2   3
//
//输出: 6
//示例 2:
//
//输入: [-10,9,20,null,null,15,7]
//
//   -10
//   / \
//  9  20
//    /  \
//   15   7
//
//输出: 42

package tree

import (
	"math"
)

func maxPathSum(root *TreeNode) int {
	ret := math.MinInt64
	var f func(node *TreeNode) int
	f = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		r := node.Val
		left := f(node.Left)
		right := f(node.Right)
		ret = int(math.Max(float64(ret), float64(r+left+right)))

		r = int(math.Max(float64(r), float64(r)+math.Max(float64(left), float64(right))))
		ret = int(math.Max(float64(ret), float64(r)))
		return r
	}

	f(root)
	return ret
}
