package tree

type RBTNode struct {
    Val int
    Color bool
    Parent *RBTNode
    Left *RBTNode
    Right *RBTNode
}

const (
    RED = false
    BLACK = true
)

type RBT struct {
}

func (r *RBT) Insert(root *RBTNode, num int) *RBTNode {
    // 没有节点
    if root == nil {
        return &RBTNode{
            num,
            BLACK,
            nil,
            nil,
            nil,
        }
    }
    var newNode *RBTNode
    _root := root
    for {
        if num > _root.Val {
            if _root.Right == nil {
                newNode = &RBTNode{
                    num,
                    RED,
                    _root,
                    nil,
                    nil,
                }
                _root.Right = newNode
                break
            }
            _root = _root.Right
        } else if num < _root.Val {
            if _root.Left == nil {
                newNode = &RBTNode{
                    num,
                    RED,
                    _root,
                    nil,
                    nil,
                }
                _root.Left = newNode
                break
            }
            _root = _root.Left
        } else {
            // 已经存在
            return root
        }
    }
    root = r.fixInsert(root, newNode)
    root.Color = BLACK
    return root
}

func (r *RBT) fixInsert(root, n *RBTNode) *RBTNode {
    var parentNode, grandParentNode *RBTNode
    for n.Color == RED && n.Parent != nil && n.Parent.Color == RED {
        parentNode = n.Parent
        grandParentNode = parentNode.Parent
        if grandParentNode == nil {
            break
        }
        if parentNode == grandParentNode.Left {
            uncleNode := grandParentNode.Right
            if uncleNode != nil && uncleNode.Color == RED {
                grandParentNode.Color = RED
                parentNode.Color = BLACK
                uncleNode.Color = BLACK
                n = grandParentNode
            } else {
                // left right
                if n == parentNode.Right {
                    grandParentNode.Left = r.leftRotate(grandParentNode.Left)
                    grandParentNode.Left.Parent = grandParentNode
                }
                // left left
                if grandParentNode == root {
                    root = r.rightRotate(grandParentNode)
                    root.Parent = nil
                    root.Right.Color = RED
                } else {
                    grandParentParentNode := grandParentNode.Parent
                    if grandParentNode == grandParentParentNode.Left {
                        grandParentParentNode.Left = r.rightRotate(grandParentParentNode.Left)
                        grandParentParentNode.Left.Parent = grandParentParentNode
                        grandParentParentNode.Left.Color = BLACK
                        grandParentParentNode.Left.Right.Color = RED
                    } else {
                        grandParentParentNode.Right = r.rightRotate(grandParentParentNode.Right)
                        grandParentParentNode.Right.Parent = grandParentParentNode
                        grandParentParentNode.Right.Color = BLACK
                        grandParentParentNode.Right.Right.Color = RED
                    }
                }
            }
        } else {
            uncleNode := grandParentNode.Left
            if uncleNode != nil && uncleNode.Color == RED {
                grandParentNode.Color = RED
                parentNode.Color = BLACK
                uncleNode.Color = BLACK
                n = grandParentNode
            } else {
                // right left
                if n == parentNode.Left {
                    grandParentNode.Right = r.rightRotate(grandParentNode.Right)
                    grandParentNode.Right.Parent = grandParentNode
                }
                // right right
                if grandParentNode == root {
                    root = r.leftRotate(root.Right)
                    root.Parent = nil
                    root.Left.Color = RED
                } else {
                    grandParentParentNode := grandParentNode.Parent
                    if grandParentNode == grandParentParentNode.Left {
                        grandParentParentNode.Left = r.leftRotate(grandParentParentNode.Left)
                        grandParentParentNode.Left.Parent = grandParentParentNode
                        grandParentParentNode.Left.Color = BLACK
                        grandParentParentNode.Left.Left.Color = RED
                    } else {
                        grandParentParentNode.Right = r.leftRotate(grandParentParentNode.Right)
                        grandParentParentNode.Right.Parent = grandParentParentNode
                        grandParentParentNode.Right.Color = BLACK
                        grandParentParentNode.Right.Left.Color = RED
                    }
                }
            }
        }
    }
    return root
}

func (r *RBT) Delete(root *RBTNode, num int) *RBTNode {
    v := r.findDeleteNode(root, num)
    if v == nil {
        return root
    }
    u := r.findSuccessor(v)
    if u == nil {
        if v == root {
            root = nil
        } else {
            // double black
            if v.Color == BLACK && (u == nil || u.Color == BLACK) {

            }
        }
    }
}

func (r *RBT) fixDoubleBlack(root, x *RBTNode) *RBTNode {
    if x == root {
        return root
    }
    //var s *RBTNode
    //if x ==
}

func (r *RBT) leftRotate(n *RBTNode) *RBTNode {
    n, n.Right, n.Right.Left = n.Right, n.Right.Left, n
    n.Left.Parent = n
    if n.Left.Right != nil {
        n.Left.Right.Parent = n.Left
    }
    return n
}

func (r *RBT) rightRotate(n *RBTNode) *RBTNode{
    n, n.Left, n.Left.Right = n.Left, n.Left.Right, n
    n.Right.Parent = n
    if n.Right.Left != nil {
        n.Right.Left.Parent = n.Right
    }
    return n
}

func (r *RBT) inOrder(root *RBTNode) []int {
    ret := make([]int, 0)
    if root.Left != nil {
        ret = append(ret, r.inOrder(root.Left)...)
    }
    ret = append(ret, root.Val)
    if root.Right != nil {
        ret = append(ret, r.inOrder(root.Right)...)
    }
    return ret
}

func (r *RBT) levelOrder(root *RBTNode) []int {
    ret := make([]int, 0)
    q := make([]*RBTNode, 0)
    q = append(q, root)
    for len(q) > 0 {
        node := q[0]
        ret = append(ret, node.Val)
        if node.Left != nil {
            q = append(q, node.Left)
        }
        if node.Right != nil {
            q = append(q, node.Right)
        }
        if len(q) == 1 {
            break
        }
        q = q[1: ]
    }
    return ret
}

func (r *RBT) findDeleteNode(root *RBTNode, num int) *RBTNode {
    if num > root.Val {
        return r.findDeleteNode(root.Right, num)
    } else if num < root.Val {
        return r.findDeleteNode(root.Left, num)
    } else if num == root.Val {
        return root
    }
    return nil
}

func (r *RBT) findSuccessor(v *RBTNode) *RBTNode {
    if v.Right != nil {
        return r.findSuccessor(v.Right)
    }
    if v.Left != nil {
        return v.Left
    }
    return nil
}
