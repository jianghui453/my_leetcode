<?php
//给定无向连通图中一个节点的引用，返回该图的深拷贝（克隆）。图中的每个节点都包含它的值 val（Int） 和其邻居的列表（list[Node]）。
//
//示例：
//
//
//
//输入：
//{"$id":"1","neighbors":[{"$id":"2","neighbors":[{"$ref":"1"},{"$id":"3","neighbors":[{"$ref":"2"},{"$id":"4","neighbors":[{"$ref":"3"},{"$ref":"1"}],"val":4}],"val":3}],"val":2},{"$ref":"4"}],"val":1}
//
//解释：
//节点 1 的值是 1，它有两个邻居：节点 2 和 4 。
//节点 2 的值是 2，它有两个邻居：节点 1 和 3 。
//节点 3 的值是 3，它有两个邻居：节点 2 和 4 。
//节点 4 的值是 4，它有两个邻居：节点 1 和 3 。
// 
//
//提示：
//
//节点数介于 1 到 100 之间。
//无向图是一个简单图，这意味着图中没有重复的边，也没有自环。
//由于图是无向的，如果节点 p 是节点 q 的邻居，那么节点 q 也必须是节点 p 的邻居。
//必须将给定节点的拷贝作为对克隆图的引用返回。

class Solution {

    /**
     * @param Node $node
     * @return Node
     */
    function cloneGraph($node) {
        if ($node === null) {
            return $node;
        }
        $nodes_old = [];
        $nodes_new = [];
        $this->copyNodes($node, $nodes_old, $nodes_new);
        for ($i = 0; $i < count($nodes_new); $i++) {
            foreach ($nodes_new[$i]->neighbors as $j => $neighbor) {
                $idx = array_search($neighbor, $nodes_old);
                $nodes_new[$i]->neighbors[$j] = $nodes_new[$idx];
            }
        }
        return $nodes_new[0];
    }

    function copyNodes(Node $node, &$nodes_old, &$nodes_new) {
        if ($node === null) {
            return;
        }
        if (in_array($node, $nodes_old)) {
            return;
        }
        $node_new = new Node($node->val, $node->neighbors);
        array_push($nodes_old, $node);
        array_push($nodes_new, $node_new);
        foreach ($node->neighbors as $neighbor) {
            $this->copyNodes($neighbor, $nodes_old, $nodes_new);
        }
    }
}