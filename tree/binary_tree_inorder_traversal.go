//给定一个二叉树，返回它的中序 遍历。
//
//示例:
//
//输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,3,2]
//进阶: 递归算法很简单，你可以通过迭代算法完成吗？

package tree

func inorderTraversal(root *TreeNode) []int {
	r := make([]int, 0)
	if root == nil {
		return r
	}

	s := make([]*TreeNode, 0)
	for root != nil {
		if root.Left != nil {
			s = append(s, root)
			root = root.Left
			continue
		}

		r = append(r, root.Val)
		if root.Right != nil {
			root = root.Right
		} else if len(s) > 0 {
			node := new(TreeNode)
			node.Right = s[len(s)-1].Right
			node.Val = s[len(s)-1].Val
			root = node
			s = s[:len(s)-1]
		} else {
			root = nil
		}
	}

	return r
}

// func inorderTraversal(root *TreeNode) []int {
// 	r := make([]int, 0)
// 	if root == nil {
// 		return r
// 	}
// 	s := make([]*TreeNode, 0)
// 	sL := 0
// 	n := root
// 	for {
// 	    if n != nil {
//            if n.Left != nil {
//                s = append(s, n)
//                sL++
//                n = n.Left
//                continue
//            }
//        } else {
//            if sL > 0 {
//                n = s[sL-1]
//                s = s[: sL-1]
//                sL --
//            } else {
//                break
//            }
//        }
//        r = append(r, n.Val)
//        if n.Right != nil {
//            n = n.Right
//            continue
//        }
//        n = nil
// 	}
// 	return r
// }
