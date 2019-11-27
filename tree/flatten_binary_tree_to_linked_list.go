//给定一个二叉树，原地将它展开为链表。
//
//例如，给定二叉树
//
//    1
//   / \
//  2   5
// / \   \
//3   4   6
//将其展开为：
//
//1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6

package tree

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}
	if root.Left != nil {
		left, right := root.Left, root.Right
		if left.Left != nil {
			flatten(left)
		}
		root.Left, root.Right = nil, left
		for left.Right != nil {
			left = left.Right
		}
		left.Right = right
	}
	flatten(root.Right)
}

//func flatten(root *TreeNode)  {
//    if root == nil {
//        return
//    }
//    var f func(node *TreeNode)
//    f = func(node *TreeNode) {
//        nodeR := node.Right
//        nodeL := node.Left
//        node.Right = nil
//        node.Left = nil
//        if nodeL != nil {
//            f(nodeL)
//            node.Right = nodeL
//        }
//        if nodeR != nil {
//            for node.Right != nil {
//                node = node.Right
//            }
//            node.Right = nodeR
//            f(nodeR)
//        }
//    }
//    f(root)
//}

// func flatten(root *TreeNode) {
// 	if root == nil || (root.Left == nil && root.Right == nil) {
// 		return
// 	}
// 	getFlatten(root)
// }

// func getFlatten(root *TreeNode) (*TreeNode, *TreeNode) {
// 	if root.Left == nil && root.Right == nil {
// 		return root, root
// 	}
// 	var left, bottom, temp *TreeNode
// 	if root.Left != nil {
// 		left, bottom = getFlatten(root.Left)
// 	}
// 	temp = root.Right
// 	if left != nil {
// 		root.Left = nil
// 		root.Right = left
// 		bottom.Right = temp
// 	}
// 	if root.Right != nil {
// 		_, bottom = getFlatten(root.Right)
// 	}
// 	return root, bottom
// }
