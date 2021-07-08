//给定一个二叉树，判断它是否是高度平衡的二叉树。
//
//本题中，一棵高度平衡二叉树定义为：
//
//一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
//
//示例 1:
//
//给定二叉树 [3,9,20,null,null,15,7]
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回 true 。
//
//示例 2:
//
//给定二叉树 [1,2,2,3,3,null,null,4,4]
//
//       1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
//返回 false 。

package tree

import (
	"math"
)

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}

	var f func(node *TreeNode) int
	f = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return int(math.Max(float64(f(node.Left)), float64(f(node.Right)))) + 1
	}

	if math.Abs(float64(f(root.Left)-f(root.Right))) > 1 {
		return false
	}

	return true
}

// func isBalanced(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	if !isBalanced(root.Left) || !isBalanced(root.Right) {
// 		return false
// 	}
// 	var f func(node *TreeNode) int
// 	f = func(node *TreeNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		left := 1
// 		right := 1
// 		if node.Left != nil {
// 			left += f(node.Left)
// 		}
// 		if node.Right != nil {
// 			right += f(node.Right)
// 		}
// 		if left > right {
// 			return left
// 		}
// 		return right
// 	}
// 	dLeft := 0
// 	dRight := 0
// 	if root.Left != nil {
// 		dLeft = f(root.Left)
// 	}
// 	if root.Right != nil {
// 		dRight = f(root.Right)
// 	}
// 	if dLeft >= dRight {
// 		return dLeft-dRight < 2
// 	}
// 	return dRight-dLeft < 2
// }

//func isBalanced(root *TreeNode) bool {
//    return balanceNode(root)!=-1
//}
//
//func balanceNode(node *TreeNode) int {
//    if node==nil{
//        return 0
//    }
//    lefth:=balanceNode(node.Left)
//    if lefth==-1{
//        return -1
//    }
//    righth:=balanceNode(node.Right)
//    if righth==-1{
//        return -1
//    }
//    maxh:=lefth
//    minush:=lefth-righth
//    if minush<0{
//        maxh=righth
//        minush=0-minush
//    }
//    if minush<=1{
//        return maxh+1
//    }
//    return -1
//}
