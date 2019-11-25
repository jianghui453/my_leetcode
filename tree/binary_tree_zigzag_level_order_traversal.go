// Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its zigzag level order traversal as:
// [
//   [3],
//   [20,9],
//   [15,7]
// ]

package tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	s := []*TreeNode{root}
	for len(s) > 0 {
		sLen := len(s)
		retLen := len(ret)

		newS := make([]*TreeNode, 0)
		retItem := make([]int, 0)

		for i := 0; i < sLen; i++ {
			retItem = append(retItem, s[i].Val)
			
			if s[i].Left != nil {
				newS = append(newS, s[i].Left)
			}
			if s[i].Right != nil {
				newS = append(newS, s[i].Right)
			}
		}

		if retLen%2 == 1 {
			retItemLen := len(retItem)
			for i := 0; i < retItemLen/2; i++ {
				retItem[i], retItem[retItemLen-1-i] = retItem[retItemLen-1-i], retItem[i]
			}
		}
		
		s = newS
		ret = append(ret, retItem)
	}

	return ret
}
