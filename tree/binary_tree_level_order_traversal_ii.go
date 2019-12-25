// Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its bottom-up level order traversal as:
// [
//   [15,7],
//   [9,20],
//   [3]
// ]

package tree

func levelOrderBottom(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		newNodes := make([]*TreeNode, 0)
		retItem := make([]int, 0)

		l := len(nodes)
		for i := 0; i < l; i++ {
			retItem = append(retItem, nodes[i].Val)
			if nodes[i].Left != nil {
				newNodes = append(newNodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				newNodes = append(newNodes, nodes[i].Right)
			}
		}

		nodes = newNodes
		ret = append([][]int{retItem}, ret...)
	}

	return ret
}

// func levelOrderBottom(root *TreeNode) [][]int {
// 	r := make([][]int, 0)
// 	if root == nil {
// 		return r
// 	}
// 	s := []*TreeNode{root}
// 	lenS := 1
// 	for lenS > 0 {
// 		rItem := make([]int, 0)
// 		sNew := make([]*TreeNode, 0)
// 		lenSNew := 0
// 		for i := 0; i < lenS; i++ {
// 			rItem = append(rItem, s[i].Val)
// 			if s[i].Left != nil {
// 				sNew = append(sNew, s[i].Left)
// 				lenSNew++
// 			}
// 			if s[i].Right != nil {
// 				sNew = append(sNew, s[i].Right)
// 				lenSNew++
// 			}
// 		}
// 		r = append([][]int{rItem}, r...)
// 		s = sNew
// 		lenS = lenSNew
// 	}
// 	return r
// }
