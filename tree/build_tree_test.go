package tree

import "testing"

//func TestBuildTree(t *testing.T) {
//    var inorder, preorder []int
//    var r []int
//    var _r *TreeNode
//
//    preorder = []int{3,9,20,15,7}
//    inorder = []int{9,3,15,20,7}
//    _r = buildTree(preorder, inorder)
//    r = getInorder(_r)
//    t.Logf("\nh=%v \nr=%v", preorder, r)
//}

func TestBuildTree(t *testing.T) {
   var inorder, postorder []int
   var r []int
   var _r *TreeNode

   postorder = []int{9,15,7,20,3}
   inorder = []int{9,3,15,20,7}
   _r = buildTree(inorder, postorder)
   r = getInorder(_r)
   t.Logf("\nh=%v \nr=%v", inorder, r)
}

func getInorder(n *TreeNode) []int {
    var f func(node *TreeNode)
    var r []int
    f = func (node *TreeNode) {
        if node.Left != nil {
            f(node.Left)
        }
        r = append(r, node.Val)
        if node.Right != nil {
            f(node.Right)
        }
    }
    f(n)
    return r
}
