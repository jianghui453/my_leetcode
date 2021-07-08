"""
序列化是指将一个数据结构转化为位序列的过程，因此可以将其存储在文件中或内存缓冲区中，以便稍后在相同或不同的计算机环境中恢复结构。

设计一个序列化和反序列化 N 叉树的算法。一个 N 叉树是指每个节点都有不超过 N 个孩子节点的有根树。序列化 / 反序列化算法的算法实现没有限制。你只需要保证 N 叉树可以被序列化为一个字符串并且该字符串可以被反序列化成原树结构即可。

例如，你需要序列化下面的 3-叉 树。

 



 

为 [1 [3[5 6] 2 4]]。你不需要以这种形式完成，你可以自己创造和实现不同的方法。

 

注意：

N 的范围在 [1, 1000]
不要使用类成员 / 全局变量 / 静态变量来存储状态。你的序列化和反序列化算法应是无状态的。

# Definition for a Node.
class Node(object):
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""

class Node(object):
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children

class Codec:

    def serialize(self, root: 'Node') -> str:
        """Encodes a tree to a single string.
        
        :type root: Node
        :rtype: str
        """
        if root is None:
            return ""

        ret = str(root.val)
        cs = []
        if root.children is not None:
            for c in root.children:
                cs.append("[" + self.serialize(c) + "]")

        ret = ret + "[" + ''.join(cs) + "]"
        return ret

    def deserialize(self, data: str) -> 'Node':
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: Node
        """
        # pprint.pprint("data=" + data)
        if len(data) == 0:
            return None
        
        v = ""
        for b in data:
            if b == '[':
                break
            else:
                v += b

        # pprint.pprint("v=" + v)
        
        c_datas = []
        i = len(v)+1
        j = len(v)+1
        cnt = 0
        while i < len(data):
            if data[i] == '[':
                cnt += 1
            elif data[i] == ']':
                cnt -= 1
            
            if cnt == 0:
                c_datas.append(data[j+1:i])
                pprint.pprint("c_data=" + data[j+1:i])
                j = i+1
            i += 1
        
        c_nodes = []
        for d in c_datas:
            c_nodes.append(self.deserialize(d))
        

        # pprint.pprint(c_nodes)
        n = Node(int(v), c_nodes)

        return n

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))

import pprint

def test():
    root = Node(1, [
        Node(3, [
            Node(5, None),
            Node(6, None)
        ]),
        Node(2, None), 
        Node(4, None)
    ])
    c = Codec()
    s = c.serialize(root)
    pprint.pprint(s)
    n = c.deserialize("1[[3[[5[]][6[]]]][2[]][4[]]]")
    s = c.serialize(n)
    pprint.pprint(s)

test()