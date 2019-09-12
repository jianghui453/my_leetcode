//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
//假设一个二叉搜索树具有如下特征：
//
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
//示例 1:
//
//输入:
//    2
//   / \
//  1   3
//输出: true
//示例 2:
//
//输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。

package tree

import "math"

func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    var f func(*TreeNode, int, int) bool
    f = func(n *TreeNode, min, max int) bool {
        if n == nil {
            return true
        }
        if n.Val >= max || n.Val <= min {
            return false
        }
        if n.Left != nil {
            _max := max
            if n.Val < max {
                _max = n.Val
            }
            if !f(n.Left, min, _max) {
                return false
            }
        }
        if n.Right != nil {
            _min := min
            if n.Val > min {
                _min = n.Val
            }
            if !f(n.Right, _min, max) {
                return false
            }
        }
        return true
    }
    return f(root, math.MinInt64, math.MaxInt64)
}