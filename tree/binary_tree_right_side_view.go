// 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

// 示例:

// 输入: [1,2,3,null,5,null,4]
// 输出: [1, 3, 4]
// 解释:

//    1            <---
//  /   \
// 2     3         <---
//  \     \
//   5     4       <---

package tree

func rightSideView(root *TreeNode) []int {
	ret := make([]int, 0)
	var f func(node *TreeNode, depth int)
	f = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth > len(ret) {
			ret = append(ret, node.Val)
		}
		if node.Right != nil {
			f(node.Right, depth+1)
		}
		if node.Left != nil {
			f(node.Left, depth+1)
		}
	}

	f(root, 1)
	return ret
}