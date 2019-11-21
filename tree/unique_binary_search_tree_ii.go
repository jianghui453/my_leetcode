//给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
//
//示例:
//
//输入: 3
//输出:
//[
//  [1,null,3,2],
//  [3,2,null,1],
//  [3,1,null,null,2],
//  [2,1,3],
//  [1,null,2,null,3]
//]
//解释:
//以上的输出对应以下 5 种不同结构的二叉搜索树：
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3

package tree

func generateTrees(n int) []*TreeNode {
	r := make([]*TreeNode, 0)
	if n < 1 {
		return r
	}
	var f func(int, int) []*TreeNode
	f = func(s, l int) []*TreeNode {
		if s > l {
			return []*TreeNode{nil}
		}
		trees := make([]*TreeNode, 0)
		for i := s; i <= l; i++ {
			leftTrees := f(s, i-1)
			rightTrees := f(i+1, l)
			for _, leftTree := range leftTrees {
				for _, rightTree := range rightTrees {
					tree := &TreeNode{i, leftTree, rightTree}
					trees = append(trees, tree)
				}
			}
		}
		return trees
	}
	return f(1, n)
}
