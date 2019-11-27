//给定一个二叉树，找出其最小深度。
//
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//
//给定二叉树 [3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回它的最小深度  2.

package tree

import (
	"math"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return minDepth(root.Right)+1
	}
	if root.Right == nil {
		return minDepth(root.Left)+1
	}
	return int(math.Min(float64(minDepth(root.Left)), float64(minDepth(root.Right))))+1
}

// func minDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	var f func(node *TreeNode) float64
// 	f = func(node *TreeNode) float64 {
// 		if node.Left == nil && node.Right == nil {
// 			return 1
// 		}
// 		if node.Left == nil {
// 			return f(node.Right)+1
// 		}
// 		if node.Right == nil {
// 			return f(node.Left)+1
// 		}
// 		return math.Min(f(node.Left), f(node.Right))+1
// 	}

// 	if root.Left == nil && root.Right == nil {
// 		return 1
// 	}
// 	if root.Left == nil {
// 		return int(f(root.Right))+1
// 	}
// 	if root.Right == nil {
// 		return int(f(root.Left))+1
// 	}

// 	return int(f(root))
// }

// func minDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	if root.Left == nil && root.Right == nil {
// 		return 1
// 	}
// 	l := 0
// 	r := 0
// 	if root.Left != nil {
// 		l = minDepth(root.Left)
// 	}
// 	if root.Right != nil {
// 		r = minDepth(root.Right)
// 	}
// 	if l == 0 {
// 		return r + 1
// 	}
// 	if r == 0 {
// 		return l + 1
// 	}
// 	if l > r {
// 		return r + 1
// 	}
// 	return l + 1
// }
