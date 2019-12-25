//给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
//
//例如:
//给定二叉树: [3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回其层次遍历结果：
//
//[
//  [3],
//  [9,20],
//  [15,7]
//]

//给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回锯齿形层次遍历如下：
//
//[
//  [3],
//  [20,9],
//  [15,7]
//]

package tree

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	s := []*TreeNode{root}
	for len(s) > 0 {
		l := len(s)
		retLen := len(ret)
		ret = append(ret, make([]int, 0))
		for i := 0; i < l; i++ {
			ret[retLen] = append(ret[retLen], s[i].Val)
			if s[i].Left != nil {
				s = append(s, s[i].Left)
			}
			if s[i].Right != nil {
				s = append(s, s[i].Right)
			}
		}

		if len(s) > l {
			s = s[l:]
		} else {
			break
		}
	}

	return ret
}

// func levelOrder(root *TreeNode) [][]int {
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
// 		r = append(r, rItem)
// 		s = sNew
// 		lenS = lenSNew
// 	}
// 	return r
// }

// func zigzagLevelOrder(root *TreeNode) [][]int {
// 	r := make([][]int, 0)
// 	if root == nil {
// 		return r
// 	}
// 	s := []*TreeNode{root}
// 	z := true
// 	lenS := 1
// 	for lenS > 0 {
// 		rItem := make([]int, 0)
// 		sNew := make([]*TreeNode, 0)
// 		lenSNew := 0
// 		for i := 0; i < lenS; i++ {
// 			rItem = append(rItem, s[i].Val)
// 			if z {
// 				if s[i].Left != nil {
// 					sNew = append([]*TreeNode{s[i].Left}, sNew...)
// 					lenSNew++
// 				}
// 				if s[i].Right != nil {
// 					sNew = append([]*TreeNode{s[i].Right}, sNew...)
// 					lenSNew++
// 				}
// 			} else {
// 				if s[i].Right != nil {
// 					sNew = append([]*TreeNode{s[i].Right}, sNew...)
// 					lenSNew++
// 				}
// 				if s[i].Left != nil {
// 					sNew = append([]*TreeNode{s[i].Left}, sNew...)
// 					lenSNew++
// 				}
// 			}
// 		}
// 		z = !z
// 		r = append(r, rItem)
// 		s = sNew
// 		lenS = lenSNew
// 	}
// 	return r
// }
