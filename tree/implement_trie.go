// 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

// 示例:

// Trie trie = new Trie();

// trie.insert("apple");
// trie.search("apple");   // 返回 true
// trie.search("app");     // 返回 false
// trie.startsWith("app"); // 返回 true
// trie.insert("app");   
// trie.search("app");     // 返回 true
// 说明:

// 你可以假设所有的输入都是由小写字母 a-z 构成的。
// 保证所有输入均为非空字符串。

package tree

import (
	// "fmt"
)

type TrieNode struct {
	Val byte
	Children []*TrieNode
}

func NewTrieNode(v byte) *TrieNode {
	n := new(TrieNode)
	n.Val = v
	n.Children = make([]*TrieNode, 0)
	return n
}

type Trie struct {
    Root *TrieNode
}


/** Initialize your data structure here. */
func Constructor() Trie {
	var t Trie
	t.Root = NewTrieNode('0')
	t.Root.Children = append(t.Root.Children, NewTrieNode('0'))
	return t
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	l := len(word)
	node := this.Root
loop:
	for i := 0; i < l; i++ {
		for j := range node.Children {
			if node.Children[j].Val == word[i] {
				node = node.Children[j]
				continue loop
			}
		}

		newnode := NewTrieNode(word[i])
		node.Children = append(node.Children, newnode)
		node = newnode
	}

	for i := range node.Children {
		if node.Children[i].Val == '0' {
			return
		}
	}

	node.Children = append(node.Children, NewTrieNode('0'))
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	r := this.traversal(word)
	if r == nil {
		return false
	}

	for i := range r.Children {
		if r.Children[i].Val == '0' {
			return true
		}
	}

	return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	r := this.traversal(prefix)

	return r != nil
}

func (this *Trie) traversal(str string) *TrieNode {
	node := this.Root
	l := len(str)
loop:
	for i := 0; i < l; i++ {
		for j := range node.Children {
			if str[i] == node.Children[j].Val {
				node = node.Children[j]
				continue loop
			}
		}

		return nil
	}

	return node
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */