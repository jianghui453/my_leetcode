//给定一个二叉树，找出其最小深度。
//
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//
//给定二叉树 [3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回它的最小深度  2.

package tree

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 1
    }
    l := 0
    r := 0
    if root.Left != nil {
        l = minDepth(root.Left)
    }
    if root.Right != nil {
        r = minDepth(root.Right)
    }
    if l == 0 {
        return r+1
    }
    if r == 0 {
        return l+1
    }
    if l > r {
        return r+1
    }
    return l+1
}