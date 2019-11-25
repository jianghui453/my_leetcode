//给定一个二叉树，找出其最大深度。
//
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回它的最大深度 3 。

package tree

import (
	"math"
)

func maxDepth(root *TreeNode) int {
	ret := 0

	var f func(node *TreeNode, depth int)
	f = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		depth++
		ret = int(math.Max(float64(ret), float64(depth)))
		
		f(node.Left, depth)
		f(node.Right, depth)
	}

	f(root, 0)

	return ret
}

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	var f func(*TreeNode, int) int
// 	f = func(node *TreeNode, r int) int {
// 		rLeft := r + 1
// 		rRight := r + 1
// 		if node.Left != nil {
// 			rLeft = f(node.Left, r+1)
// 		}
// 		if node.Right != nil {
// 			rRight = f(node.Right, r+1)
// 		}
// 		if rLeft > rRight {
// 			return rLeft
// 		}
// 		return rRight
// 	}
// 	r := f(root, 0)
// 	return r
// }
