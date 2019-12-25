package tree

import (
	"testing"
	"fmt"
)

func Test_trie(t *testing.T) {
	var (
		trie *Trie
	)

	_t := Constructor()
	trie = &(_t)

	trie.Insert("hello")
	printStr(trie)

	t.Log("Search =", trie.Search("hello"), "hope = true")
	t.Log("Search =", trie.Search("hell"), "hope = false")

	t.Log("Search =", trie.StartsWith("hell"), "hope = true")
	t.Log("Search =", trie.StartsWith("hello"), "hope = true")
}

func printStr(t *Trie) {
	var (
		f func(node *TrieNode, bytes []byte)
	)

	f = func(node *TrieNode, bytes []byte) {
		for i := range node.Children {
			if node.Children[i].Val == '0' {
				fmt.Println(string(bytes))
			} else {
				f(node.Children[i], append(bytes, node.Children[i].Val))
			}
		}
	}

	f(t.Root, make([]byte, 0))
}