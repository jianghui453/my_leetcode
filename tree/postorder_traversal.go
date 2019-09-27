//给定一个二叉树，返回它的 后序 遍历。
//
//示例:
//
//输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [3,2,1]
//进阶: 递归算法很简单，你可以通过迭代算法完成吗？

package tree

func postorderTraversal(root *TreeNode) []int {
    ret := make([]int, 0)
    s := make([]*TreeNode, 0)
    his := make(map[*TreeNode]bool)
    for root != nil {
        if root.Left != nil {
            if _, ok := his[root.Left]; !ok {
                s = append(s, root)
                root = root.Left
                continue
            }
        }
        if root.Right != nil {
            if _, ok := his[root.Right]; !ok {
                s = append(s, root)
                root = root.Right
                continue
            }
        }
        ret = append(ret, root.Val)
        his[root] = true
        if len(s) > 0 {
            root = s[len(s)-1]
            s = s[: len(s)-1]
        } else {
            root = nil
        }
    }
    return ret
}
