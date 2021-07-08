# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        nodes = []
        def preorder(node):
            if node is None:
                nodes.append("None")
                return
            
            nodes.append(str(node.val))
            preorder(node.left)
            preorder(node.right)
        
        preorder(root)

        return "_".join(nodes)

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        nodes = data.split("_")
        l = len(nodes)
        idx = 0
        def build():
            nonlocal idx
            nonlocal nodes
            nonlocal l

            if idx >= l:
                return
            
            if nodes[idx] == 'None':
                idx += 1
                return

            node = TreeNode(int(nodes[idx]))
            idx += 1

            node.left = build()
            node.right = build()

            return node

        root = build()
        return root

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))

import pprint

def test():
    pprint.pprint("test")
    # root = TreeNode(10)
    # root.left = TreeNode(5)
    # root.left.left = TreeNode(2)
    # root.left.left.left = TreeNode(1)
    # root.left.left.right = TreeNode(3)
    # root.left.right = TreeNode(7)
    # root.right = TreeNode(15)

    root = TreeNode(0)
    root.left = TreeNode(-1)
    root.right = TreeNode(1)

    c = Codec()
    _str = c.serialize(root)
    pprint.pprint(_str)

    node = c.deserialize(_str)
    
    def preorder(n):
        if n is not None:
            pprint.pprint(n.val)
            preorder(n.left)
            preorder(n.right)
        else:
            pprint.pprint("None")
    preorder(node)

test()