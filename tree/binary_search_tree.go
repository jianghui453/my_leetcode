package tree

type BST struct{
	Root *TreeNode
}

func NewBST(nums []int) *BST {
	b := &BST{}
	if len(nums) == 0 {
		return b
	}
	b.Root = &TreeNode{
		Val: nums[0],
	}
	for i := 1; i < len(nums); i++ {
		b.Insert(nums[i])
	}
	return b
}

func (b *BST) Insert(num int) {
	var gp, g, p *TreeNode
	node := b.Root
	for {
		if num == node.Val {
			break
		}
		gp, g, p := g, p, node
		if num < node.Val {
			if node.Left == nil {
				node.Left = &TreeNode{
					Val: num,
				}
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = &TreeNode{
					Val: num,
				}
			}
			node = node.Right
		}
	}
	if g != nil && g.Left == nil {
		if p.Left == nil {
			//       g
			//         \
			//           p
			//             \
			//               node
			if gp.Left == g {
				gp.Left, p.Left, g.Right = p, g, nil
			} else {
				gp.Right, p.Left, g.Right = p, g, nil
			}
		} else {
			//       g
			//         \
			//           p
			//          /
			//         node
			if gp.Left == g {
				gp.Left, p.Left, p.Right = p, g, node
			}
		}
	} else {
		if p.Left == nil {
			//         g
			//       /
			//     p
			//       \
			//        node
			if gp.Left == g {
				gp.Left, p.Left, p.Right, g.Left = p, node, g, nil
			} else {
				gp.Right, p.Left, p.Right, g.Left = p, node, g, nil
			}
		} else {
			//         g
			//       /
			//     p
			//   /
			// node
			if gp.Left == g {
				gp.Left, p.Right, g.Left = p, g, nil
			} else {
				gp.Right, p.Right, g.Left = p, g, nil
			}
		}
	}
}
