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

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	var first, second, prev, cur, prevm *TreeNode
	cur = root
	for cur != nil {
	    //fmt.Printf("cur.Val=%d cur.Left=%v cur.Right=%v\n", cur.Val, cur.Left, cur.Right)
		if cur.Left == nil {
		    //fmt.Printf("leftNil prev=%v first=%v second=%v\n", prev, first, second)
			if prev != nil && cur.Val < prev.Val {
				if first == nil {
					first = prev
					second = cur
				} else {
					second = cur
				}
			}
			prev = cur
			cur = cur.Right
		} else {
			prevm = cur.Left
			for prevm.Right != nil && prevm.Right != cur {
			    //fmt.Printf("a prevm.Val=%d\n", prevm.Val)
				prevm = prevm.Right
			}
			if prevm.Right == nil {
				prevm.Right = cur
				cur = cur.Left
			} else {
				prevm.Right = nil
				if prev != nil && cur.Val < prev.Val {
					if first == nil {
						first = prev
						second = cur
					} else {
						second = cur
					}
				}
				prev = cur
				cur = cur.Right
			}
		}
	}

	first.Val, second.Val = second.Val, first.Val
	return
}
