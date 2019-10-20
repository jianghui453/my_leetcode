package tree

import (
    "fmt"
    "math"
)

type AVLNode struct {
    Val int
    Left *AVLNode
    Right *AVLNode
    Height int
}

type AVL struct {
}

func (a *AVL)Insert(root *AVLNode, num int) *AVLNode {
    if root == nil {
        return &AVLNode{
            num,
            nil,
            nil,
            1,
        }
    }
    if num > root.Val {
        root.Right = a.Insert(root.Right, num)
    } else if num < root.Val {
        root.Left = a.Insert(root.Left, num)
    } else {
        return root
    }
    root.Height = a.getHeight(root)
    heightDiff := a._getHeight(root.Left) - a._getHeight(root.Right)
    if heightDiff > 1 {
        if num > root.Left.Val {
            root.Left = a.leftRotate(root.Left)
        }
        return a.rightRotate(root)
    }
    if heightDiff < -1 {
        if num < root.Right.Val {
            root.Right = a.rightRotate(root.Right)
        }
        return a.leftRotate(root)
    }
    return root
}

func (a *AVL)Delete(root *AVLNode, num int) *AVLNode {
    fmt.Printf("Delete\n")
    if root == nil {
        return root
    }
    if num > root.Val {
        root.Right = a.Delete(root.Right, num)
    } else if num < root.Val {
        root.Left = a.Delete(root.Left, num)
    } else {
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }
        root.Val = a.getSmallVal(root.Right)
        root.Right = a.Delete(root.Right, root.Val)
    }
    root.Height = a.getHeight(root)
    fmt.Printf("root=%v\n", *root)
    heightDiff := a._getHeight(root.Left) - a._getHeight(root.Right)
    fmt.Printf("heightDiff=%d\n", heightDiff)
    if heightDiff > 1 {
        if num < root.Left.Val {
            root.Left = a.leftRotate(root.Left)
        }
        return a.rightRotate(root)
    }
    if heightDiff < -1 {
        if num > root.Right.Val {
            root.Right = a.rightRotate(root.Right)
        }
        return a.leftRotate(root)
    }
    return root
}

func (a *AVL)leftRotate(n *AVLNode) *AVLNode {
    fmt.Printf("leftRotate\n")
    n, n.Right, n.Right.Left = n.Right, n.Right.Left, n
    n.Left.Height = a.getHeight(n.Left)
    n.Height = a.getHeight(n)
    return n
}

func (a *AVL)rightRotate(n *AVLNode) *AVLNode {
    fmt.Printf("rightRotate\n")
    n, n.Left, n.Left.Right = n.Left, n.Left.Right, n
    n.Right.Height = a.getHeight(n.Right)
    n.Height = a.getHeight(n)
    return n
}

func (a *AVL)getSmallVal(n *AVLNode) int {
    if n.Left == nil {
        return n.Val
    }
    return a.getSmallVal(n.Left)
}

func (a *AVL)getHeight(n *AVLNode) int {
    return 1 + int(math.Max(float64(a._getHeight(n.Left)), float64(a._getHeight(n.Right))))
}

func (a *AVL)_getHeight(n *AVLNode) int {
    if n == nil {
        return 0
    }
    return n.Height
}

func (a *AVL)preOrder(root *AVLNode) []int {
    nums := make([]int, 0)
    nums = append(nums, root.Val)
    if root.Left != nil {
        nums = append(nums, a.preOrder(root.Left)...)
    }
    if root.Right != nil {
        nums = append(nums, a.preOrder(root.Right)...)
    }
    return nums
}
