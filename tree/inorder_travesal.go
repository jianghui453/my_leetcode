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
	var f func(*TreeNode)
	f = func(n *TreeNode) {
		if n.Left != nil {
			f(n.Left)
		}
		r = append(r, n.Val)
		if n.Right != nil {
			f(n.Right)
		}
	}
	f(root)
	return r
}

//func inorderTraversal(root *TreeNode) []int {
//	r := make([]int, 0)
//	if root == nil {
//		return r
//	}
//	s := make([]*TreeNode, 0)
//	sL := 0
//	n := root
//	for {
//	    if n != nil {
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
//	}
//	return r
//}
