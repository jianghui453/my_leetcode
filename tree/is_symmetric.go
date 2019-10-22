//给定一个二叉树，检查它是否是镜像对称的。
//
//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
//说明:
//
//如果你可以运用递归和迭代两种方法解决这个问题，会很加分。

package tree

import "fmt"

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	sLeft := make([]*TreeNode, 0)
	sRight := make([]*TreeNode, 0)
	lenLeft := 0
	lenRight := 0
	var f func() bool
	f = func() bool {
		fmt.Printf("sLeft=%v sRight=%v\n", sLeft, sRight)
		sLeftT := make([]*TreeNode, 0)
		lenLeftT := 0
		sRightT := make([]*TreeNode, 0)
		lenRightT := 0

		for i := 0; i < lenLeft; i++ {
			tLeft := sLeft[i]
			tRight := sRight[i]
			if tLeft.Val != tRight.Val {
				return false
			}

			if tLeft.Left == nil {
				if tRight.Right != nil {
					return false
				}
			} else {
				if tRight.Right == nil {
					return false
				}
				sLeftT = append(sLeftT, tLeft.Left)
				sRightT = append(sRightT, tRight.Right)
				lenLeftT++
				lenRightT++
			}

			if tLeft.Right == nil {
				if tRight.Left != nil {
					return false
				}
			} else {
				if tRight.Left == nil {
					return false
				}
				sLeftT = append(sLeftT, tLeft.Right)
				sRightT = append(sRightT, tRight.Left)
				lenLeftT++
				lenRightT++
			}
		}
		sLeft = sLeftT
		sRight = sRightT
		lenLeft = lenLeftT
		lenRight = lenRightT
		if lenLeft == 0 {
			return lenRight == 0
		}
		if lenLeft != lenRight {
			return false
		}
		r := f()
		return r
	}
	if root.Left == nil {
		return root.Right == nil
	}
	if root.Right == nil {
		return false
	}
	sLeft = append(sLeft, root.Left)
	sRight = append(sRight, root.Right)
	lenLeft++
	lenRight++
	r := f()
	return r
}
