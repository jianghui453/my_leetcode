//根据一棵树的前序遍历与中序遍历构造二叉树。
//
//注意:
//你可以假设树中没有重复的元素。
//
//例如，给出
//
//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//返回如下的二叉树：
//
//    3
//   / \
//  9  20
//    /  \
//   15   7

package tree

//func buildTree(preorder []int, inorder []int) *TreeNode {
//    var r *TreeNode
//    if len(preorder) == 0 {
//        return r
//    }
//    var f func(*TreeNode, []int)
//    f = func(n *TreeNode, in []int) {
//        if len(preorder) == 0 {
//            return
//        }
//        n.Val = preorder[0]
//        if len(preorder) == 1 {
//            return
//        }
//        preorder = preorder[1: ]
//        for i := 0; i < len(in); i ++ {
//            if in[i] == n.Val {
//                if i > 0 {
//                    n.Left = new(TreeNode)
//                    f(n.Left, in[: i])
//                }
//                if i < len(in)-1 {
//                    n.Right = new(TreeNode)
//                    f(n.Right, in[i+1: ])
//                }
//            }
//        }
//    }
//    r = new(TreeNode)
//    f(r, inorder)
//    return r
//}

//根据一棵树的中序遍历与后序遍历构造二叉树。
//
//注意:
//你可以假设树中没有重复的元素。
//
//例如，给出
//
//中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3]
//返回如下的二叉树：
//
//    3
//   / \
//  9  20
//    /  \
//   15   7

func buildTree(inorder []int, postorder []int) *TreeNode {
	var r *TreeNode
	if len(postorder) == 0 {
		return r
	}
	var f func(*TreeNode, []int)
	f = func(n *TreeNode, in []int) {
		if len(postorder) == 0 {
			return
		}
		n.Val = postorder[len(postorder)-1]
		if len(postorder) == 1 {
			return
		}
		postorder = postorder[:len(postorder)-1]
		for i := 0; i < len(in); i++ {
			if in[i] == n.Val {
				if i < len(in)-1 {
					n.Right = new(TreeNode)
					f(n.Right, in[i+1:])
				}
				if i > 0 {
					n.Left = new(TreeNode)
					f(n.Left, in[:i])
				}
			}
		}
	}
	r = new(TreeNode)
	f(r, inorder)
	return r
}
