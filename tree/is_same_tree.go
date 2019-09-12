//给定两个二叉树，编写一个函数来检验它们是否相同。
//
//如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
//
//示例 1:
//
//输入:       1         1
//          / \       / \
//         2   3     2   3
//
//        [1,2,3],   [1,2,3]
//
//输出: true
//示例 2:
//
//输入:      1          1
//          /           \
//         2             2
//
//        [1,2],     [1,null,2]
//
//输出: false
//示例 3:
//
//输入:       1         1
//          / \       / \
//         2   1     1   2
//
//        [1,2,1],   [1,1,2]
//
//输出: false

package tree

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil {
        return q == nil
    }
    if q == nil {
        return p == nil
    }
    var prevp, prevq *TreeNode
    for p != nil && q != nil {
        if p.Left == nil {
            if q.Left != nil || p.Val != q.Val {
                return false
            }
            p = p.Right
            q = q.Right
        } else {
            if q.Left == nil {
                return false
            }
            prevp = p.Left
            prevq = q.Left
            for prevp.Right != nil && prevp.Right != p {
                if prevq.Right == nil || prevq.Right == q {
                    return false
                }
                prevp = prevp.Right
                prevq = prevq.Right
            }
            if prevp.Right == nil {
                if prevq.Right != nil {
                    return false
                }
                prevp.Right = p
                prevq.Right = q
                p = p.Left
                q = q.Left
            } else {
                if prevq.Right != q || p.Val != q.Val  {
                    return false
                }
                prevp.Right = nil
                prevq.Right = nil
                p = p.Right
                q = q.Right
            }
        }
    }
    return p == nil && q == nil
}