//二叉搜索树中的两个节点被错误地交换。
//
//请在不改变其结构的情况下，恢复这棵树。
//
//示例 1:
//
//输入: [1,3,null,null,2]
//
//   1
//  /
// 3
//  \
//   2
//
//输出: [3,1,null,null,2]
//
//   3
//  /
// 1
//  \
//   2
//示例 2:
//
//输入: [3,1,4,null,null,2]
//
//  3
// / \
//1   4
//   /
//  2
//
//输出: [2,1,4,null,null,3]
//
//  2
// / \
//1   4
//   /
//  3
//进阶:
//
//使用 O(n) 空间复杂度的解法很容易实现。
//你能想出一个只使用常数空间的解决方案吗？

package tree

import (
	// "math"
	// "fmt"
)

func recoverTree(root *TreeNode) {
	var lastNode, nextNode, node1, node2 *TreeNode
	for root != nil {
		if root.Left == nil {
			if lastNode != nil && root.Val < lastNode.Val {
				if node1 == nil {
					node1, node2 = lastNode, root
				} else {
					node2 = root
				}
			}
			lastNode, root = root, root.Right
		} else {
			nextNode = root.Left
			for nextNode.Right != nil && nextNode.Right != root {
				nextNode = nextNode.Right
			}

			if nextNode.Right == nil {
				nextNode.Right = root
				root = root.Left
			} else {
				if lastNode != nil && root.Val < lastNode.Val {
					if node1 == nil {
						node1, node2 = lastNode, root
					} else {
						node2 = root
					}
				}
				nextNode.Right = nil
				lastNode, root = root, root.Right
			}
		}
	}

	if node1 != nil {
		node1.Val, node2.Val = node2.Val, node1.Val
	}
}

// func recoverTree(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	var first, second, prev, cur, prevm *TreeNode
// 	cur = root
// 	for cur != nil {
// 		if cur.Left == nil {
// 			if prev != nil && cur.Val < prev.Val {
// 				if first == nil {
// 					first = prev
// 					second = cur
// 				} else {
// 					second = cur
// 				}
// 			}
// 			prev = cur
// 			cur = cur.Right
// 		} else {
// 			prevm = cur.Left
// 			for prevm.Right != nil && prevm.Right != cur {
// 				prevm = prevm.Right
// 			}
// 			if prevm.Right == nil {
// 				prevm.Right = cur
// 				cur = cur.Left
// 			} else {
// 				prevm.Right = nil
// 				if prev != nil && cur.Val < prev.Val {
// 					if first == nil {
// 						first = prev
// 						second = cur
// 					} else {
// 						second = cur
// 					}
// 				}
// 				prev = cur
// 				cur = cur.Right
// 			}
// 		}
// 	}

// 	first.Val, second.Val = second.Val, first.Val
// 	return
// }
