//给定一个非空二叉树，返回其最大路径和。
//
//本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
//
//示例 1:
//
//输入: [1,2,3]
//
//       1
//      / \
//     2   3
//
//输出: 6
//示例 2:
//
//输入: [-10,9,20,null,null,15,7]
//
//   -10
//   / \
//  9  20
//    /  \
//   15   7
//
//输出: 42

package tree

// 本节点最大路径等于本节点大小、本节点加左子树最大路径和、本节点加右子树最大路径和、本节点加左右子树路径和之间取最大值
// 递归求取子树的路径和，求子树和的时候只能选择其中一边走，因为这个路径不支持跳跃或者重复路径
// 使用map保存已求得的子树最大路径和
// 返回所有节点最大路径和中的最大值
func maxPathSum(root *TreeNode) int {
	var f func(node *TreeNode) int
	r := root.Val
	f = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftPathSum := f(node.Left)
		rightPathSum := f(node.Right)
		maxPathSum := max(max(node.Val, node.Val+leftPathSum), node.Val+rightPathSum)
		maxR := max(maxPathSum, node.Val+leftPathSum+rightPathSum)
		if maxR > r {
			r = maxR
		}
		return maxPathSum
	}
	f(root)
	return r
}

//const (
//  minInt32 int = -1 << 31
//)
//
//func max(args ...int) int {
//  r := args[0]
//  for _, v := range args {
//
//      if v > r {
//          r = v
//      }
//  }
//  return r
//}
//
//func helper(r *TreeNode, globalMaxP *int) (int, *int) {
//  leftMax := 0
//  rightMax := 0
//  if r.Left == nil {
//      leftMax = minInt32
//  } else {
//      leftMax, _ = helper(r.Left, globalMaxP)
//  }
//  if r.Right == nil {
//      rightMax = minInt32
//  } else {
//      rightMax, _ = helper(r.Right, globalMaxP)
//  }
//
//  fmt.Printf("rightMax=%d leftMax=%d\n", rightMax, leftMax)
//  res := r.Val
//  if leftMax > 0 {
//      res += leftMax
//  }
//  if rightMax > 0 {
//      res += rightMax
//  }
//
//    fmt.Printf("r.Val=%d *globalMaxP=%d res=%d\n", r.Val, *globalMaxP, res)
//  *globalMaxP = max(*globalMaxP, res)
//
//  return r.Val + max(0, leftMax, rightMax), globalMaxP
//}
//
//func maxPathSum(root *TreeNode) int {
//  var globalMax = minInt32
//  globalMaxP := &globalMax
//  _, maxP := helper(root, globalMaxP)
//
//  return *maxP
//}
